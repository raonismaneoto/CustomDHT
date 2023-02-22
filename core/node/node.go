package node

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/raonismaneoto/CustomDHT/commons/grpc_api"
	"github.com/raonismaneoto/CustomDHT/commons/helpers"
	client2 "github.com/raonismaneoto/CustomDHT/core/client"
	"github.com/raonismaneoto/CustomDHT/core/models"
	"github.com/raonismaneoto/CustomDHT/core/storage"
)

type Node struct {
	fingerTable       []models.NodeRepresentation
	id                int64
	address           string
	storage           storage.Storage
	predecessor       models.NodeRepresentation
	nSucc             models.NodeRepresentation
	M                 int
	replicationBuffer chan storage.Entry
}

func New(id int64, address string, m int) *Node {
	return &Node{id: id, address: address, M: m}
}

func (n *Node) Start(partnerId int64, partnerAddr string) {
	partner := &models.NodeRepresentation{
		Id:      partnerId,
		Address: partnerAddr,
	}

	log.Println("Starting node: %v", n.id)
	log.Println("partnerAddr: %v", partnerAddr)
	log.Println("nodeAddr: %v", n.address)
	log.Println("creating replicationBuffer")
	n.replicationBuffer = make(chan storage.Entry, 50)
	n.fingerTable = make([]models.NodeRepresentation, n.M, n.M)
	n.storage = storage.New(models.GetMemTypeFromString(os.Getenv("STORAGE_TYPE")))
	n.predecessor = models.NodeRepresentation{
		Id:      0,
		Address: "",
	}
	n.nSucc = models.NodeRepresentation{
		Id:      0,
		Address: "",
	}

	if partner.Address != n.address {
		n.Join(partner)
	}

	helpers.PeriodicInvocation(n.checkSucc, 60)
	helpers.PeriodicInvocation(n.stabilize, 120)

	go n.syncReplicatedKeys()
	n.printState()
	helpers.PeriodicInvocation(n.printState, 600)
}

func (n *Node) syncReplicatedKeys() {
	for msg := range n.replicationBuffer {
		err := n.storage.Save(msg)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (n *Node) checkSucc() {
	if !n.isFingerSet(0) {
		return
	}

	client := &client2.Client{}
	_, err := client.Ping(n.fingerTable[0].Address)

	if err != nil {
		if n.nSucc.Address == "" {
			return
		}

		response, err := client.HandleNewPredecessor(n.nSucc.Address, models.NodeRepresentation{Id: n.id, Address: n.address})

		if err != nil || !response.Ok {
			log.Println("handle new predecessor failed for nSucc.")
			return
		}

		n.fingerTable[0] = models.NodeRepresentation{Id: n.nSucc.Id, Address: n.nSucc.Address}

		nSucc, err := client.Successor(n.nSucc.Address)
		if err != nil {
			log.Println("unable to reach nSucc in checkSucc. Err: " + err.Error())
		}

		n.nSucc = models.NodeRepresentation{Id: nSucc.Id, Address: n.address}

		n.syncKeys()
	}
}

func (n *Node) Join(partner *models.NodeRepresentation) {
	client := client2.Client{}
	log.Println("Starting finger table")
	n.startFingerTable(partner, client)
	// notify other nodes to update their predecessors and finger table
	log.Println("Starting join notifications")
	nodeRepresentation := models.NodeRepresentation{Id: n.id, Address: n.address}
	newPredResponse, err := client.HandleNewPredecessor(n.fingerTable[0].Address, nodeRepresentation)
	if err != nil || !newPredResponse.Ok {
		log.Println("Successor did not accept new predecessor" + err.Error())
		panic("Successor did not accept new predecessor")
	}
	newSuccResponse, err := client.HandleNewSuccessor(n.predecessor.Address, nodeRepresentation, models.NodeRepresentation{Id: 0, Address: ""})
	if err != nil || !newSuccResponse.Ok {
		log.Println("Predecessor did not accept new sucessor" + err.Error())
		panic("Predecessor did not accept new sucessor")
	}
	// get the data
	n.syncKeys()
}

func (n *Node) RepSave(key int64, data []byte) {
	n.replicationBuffer <- storage.Entry{Key: key, Data: data}
}

func (n *Node) Save(key int64, value []byte) error {
	client := client2.Client{}
	if n.mustKeyBeInNode(key) {
		log.Println("saving the data in this node")
		err := n.storage.Save(storage.Entry{Key: key, Data: value})
		if err != nil {
			log.Println(err.Error())
			return err
		}
		inflectionPoint := (n.id-n.predecessor.Id)/2 + n.predecessor.Id
		if key >= inflectionPoint {
			if n.isFingerSet(0) {
				client.RepSave(n.fingerTable[0].Address, key, value)
			}
		} else {
			if n.predecessor.Address != "" {
				client.RepSave(n.predecessor.Address, key, value)
			}
		}
		return nil
	}

	// else the request must be passed to the responsible node
	log.Println("going to try to forward the save")
	response, err := n.Owner(key)
	if err != nil {
		return err
	}
	_, err = client.Save(response.OwnerNodeEndpoint, key, value)
	return err
}

func (n *Node) Delete(key int64) error {
	err := n.storage.Delete(key)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	client := client2.Client{}
	inflectionPoint := (n.id-n.predecessor.Id)/2 + n.predecessor.Id
	if key >= inflectionPoint {
		if n.isFingerSet(0) {
			client.Delete(n.fingerTable[0].Address, key)
		}
	} else {
		if n.predecessor.Address != "" {
			client.Delete(n.predecessor.Address, key)
		}
	}

	return nil
}

func (n *Node) QueryAsync(key int64, cbuffer chan grpc_api.QueryResponse) {
	if n.mustKeyBeInNode(key) {
		log.Println("going to return the query from this node")
		ebuffer := make(chan error)
		bcbuffer := make(chan []byte)

		go n.storage.ReadAsync(key, bcbuffer, ebuffer)

		for {
			select {
			case content, ok := <-bcbuffer:
				if !ok {
					close(cbuffer)
					return
				} else {
					resp := grpc_api.QueryResponse{
						Data:                    content,
						ResponsibleNodeEndpoint: n.address,
						ResponsibleNodeId:       n.id,
					}
					cbuffer <- resp
				}
			case err, ok := <-ebuffer:
				if !ok {
					close(cbuffer)
					return
				}
				if err != nil && err.Error() == "Key not found" {
					log.Println("Key" + strconv.FormatInt(key, 10) + " not found.")
					resp := grpc_api.QueryResponse{
						Data:                    []byte{},
						ResponsibleNodeEndpoint: "",
					}
					cbuffer <- resp
					close(cbuffer)
					return
				}
			}
		}

	}

	aimingNode := n.findAimingNode(key)
	if aimingNode.Address != "" {
		log.Println("key not found in node, going to forward the query to:")
		log.Println("nodeAddress: " + aimingNode.Address)
		client := client2.Client{}
		owner, err := client.Owner(aimingNode.Address, key)
		if err != nil {
			resp := grpc_api.QueryResponse{
				Data:                    []byte{},
				ResponsibleNodeEndpoint: "",
			}
			cbuffer <- resp
		}
		resp := grpc_api.QueryResponse{
			Data:                    []byte{},
			ResponsibleNodeEndpoint: owner.OwnerNodeEndpoint,
			ResponsibleNodeId:       owner.OwnerNodeId,
		}
		cbuffer <- resp
	}

	log.Println("unable to query for key" + strconv.FormatInt(key, 10))
	log.Println("going to panic")
	panic("unable to query for key" + strconv.FormatInt(key, 10))
}

func (n *Node) Query(key int64) grpc_api.QueryResponse {
	if n.mustKeyBeInNode(key) {
		log.Println("going to return the query from this node")
		data, err := n.storage.Read(key)

		if err != nil && err.Error() == "Key not found" {
			log.Println("Key" + strconv.FormatInt(key, 10) + " not found.")
			return grpc_api.QueryResponse{
				Data:                    []byte{},
				ResponsibleNodeEndpoint: n.address,
				ResponsibleNodeId:       n.id,
			}
		}

		return grpc_api.QueryResponse{
			Data:                    data,
			ResponsibleNodeEndpoint: n.address,
			ResponsibleNodeId:       n.id,
		}
	}

	aimingNode := n.findAimingNode(key)
	if aimingNode.Address != "" {
		log.Println("key not found in node, going to forward the query to:")
		log.Println("nodeAddress: " + aimingNode.Address)
		client := client2.Client{}
		return *client.Query(aimingNode.Address, key)
	}

	log.Println("unable to query for key" + strconv.FormatInt(key, 10))
	log.Println("going to panic")
	panic("unable to query for key" + strconv.FormatInt(key, 10))
}

func distance(i int64, j int64, m int) int64 {
	if j == i {
		return 0
	}
	if j > i {
		return j - i
	}
	return int64(math.Pow(2, float64(m))) - i + j
}

func (n *Node) HandleNewSuccessor(newSucc models.NodeRepresentation, nNSucc models.NodeRepresentation) error {
	client := &client2.Client{}
	if n.fingerTable[0].Address != "" && distance(n.id, newSucc.Id, n.M) > distance(n.id, n.fingerTable[0].Id, n.M) {
		log.Println("ping current succ")
		_, err := client.Ping(n.fingerTable[0].Address)
		if err == nil {
			return errors.New("invalid successor")
		}
	}

	n.fingerTable[0] = newSucc
	if nNSucc.Address != "" && nNSucc.Id != n.id {
		n.nSucc = nNSucc
	}

	go n.syncSuccKeys()

	return nil
}

func (n *Node) HandleNewPredecessor(nPred models.NodeRepresentation) error {
	client := &client2.Client{}
	if n.predecessor.Address != "" && distance(n.id, nPred.Id, n.M) < distance(n.id, n.predecessor.Id, n.M) {
		_, err := client.Ping(n.predecessor.Address)
		if err == nil {
			return errors.New("invalid predecessor")
		}
	}

	n.predecessor = nPred
	// check if this goroutine will die after this function return statement
	go n.syncPredKeys()

	return nil
}

func (n *Node) Successor() (*models.NodeRepresentation, error) {
	if n.fingerTable == nil || len(n.fingerTable) == 0 || n.fingerTable[0].Address == "" {
		return nil, errors.New("There is no successor")
	}

	return &n.fingerTable[0], nil
}

func (n *Node) Predecessor() (models.NodeRepresentation, error) {
	if n.predecessor.Address == "" {
		return n.predecessor, errors.New("There is no predecessor")
	}

	return n.predecessor, nil
}

func (n *Node) Owner(key int64) (*grpc_api.OwnerResponse, error) {
	if n.mustKeyBeInNode(key) {
		return &grpc_api.OwnerResponse{
			OwnerNodeId:       n.id,
			OwnerNodeEndpoint: n.address,
		}, nil
	}

	aimingNode := n.findAimingNode(key)
	log.Println("aimingNOdeAddr:")
	log.Println(aimingNode.Address)
	log.Println(aimingNode.Id)
	log.Println(n.fingerTable[0].Address)
	if aimingNode.Address != "" {
		log.Println("kgoing to forward the query to:")
		log.Println("nodeAddress: " + aimingNode.Address)
		client := client2.Client{}
		return client.Owner(aimingNode.Address, key)
	}

	return nil, errors.New("unable to get the owner of the key: " + fmt.Sprint(key))

}

func (n *Node) stabilize() {
	for i := 1; i < n.M; i++ {
		currNodeInfo, err := n.Owner((n.id + int64(math.Pow(2, float64(i-1)))) % int64(math.Pow(2, float64(n.M))))
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if currNodeInfo.OwnerNodeEndpoint == "" {
			continue
		}
		if currNodeInfo.OwnerNodeId != n.id && currNodeInfo.OwnerNodeEndpoint != n.address {
			n.fingerTable[i] = models.NodeRepresentation{Id: currNodeInfo.OwnerNodeId, Address: currNodeInfo.OwnerNodeEndpoint}
		}
	}
}

func (n *Node) startFingerTable(partner *models.NodeRepresentation, client client2.Client) {
	log.Println("querying succ info in startFingerTable")
	succInfo, err := client.Owner(partner.Address, n.id)
	if err != nil {
		log.Println("going to panic: " + err.Error())
		panic(err.Error())
	}
	succ := models.NodeRepresentation{Id: succInfo.OwnerNodeId, Address: succInfo.OwnerNodeEndpoint}
	n.fingerTable[0] = succ

	nSuccInfo, err := client.Successor(n.fingerTable[0].Address)
	if err != nil {
		//if the partner does not have successor it does not have predecessor either
		log.Println("Unable to find nsucc on finger table startup. Err: %v", err)
		n.predecessor = succ
	} else {
		n.nSucc = models.NodeRepresentation{Id: nSuccInfo.Id, Address: nSuccInfo.Endpoint}
		predecessor, err := client.Predecessor(n.fingerTable[0].Address)
		if err != nil {
			log.Println(err.Error())
		}
		n.predecessor = models.NodeRepresentation{Id: predecessor.Id, Address: predecessor.Endpoint}
	}

	for i := 1; i < n.M; i++ {
		currNodeInfo, err := n.Owner((n.id + int64(math.Pow(2, float64(i-1)))) % int64(math.Pow(2, float64(n.M))))
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if currNodeInfo.OwnerNodeId != n.id && currNodeInfo.OwnerNodeEndpoint != n.address {
			n.fingerTable[i] = models.NodeRepresentation{Id: currNodeInfo.OwnerNodeId, Address: currNodeInfo.OwnerNodeEndpoint}
		}
	}
}

func (n *Node) mustKeyBeInNode(key int64) bool {
	if n.predecessor.Address == "" {
		return true
	}

	if n.predecessor.Address != "" && n.predecessor.Id > n.id {
		return key < n.id || key >= n.predecessor.Id
	}

	return (key >= n.predecessor.Id && key < n.id)
}

func (n *Node) keysRange() (int64, int64) {
	client := &client2.Client{}
	predOfPred, _ := client.Predecessor(n.predecessor.Address)
	start := (n.predecessor.Id-predOfPred.Id)/2 + predOfPred.Id
	end := (n.nSucc.Id-n.fingerTable[0].Id)/2 + n.fingerTable[0].Id
	return start, end
}

func (n *Node) syncKey(address string, key int64) {
	client := client2.Client{}
	response := client.Query(address, key)
	if response.Data != nil && len(response.Data) > 0 {
		err := n.storage.Save(storage.Entry{Key: key, Data: response.Data})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (n *Node) syncKeys() {
	log.Println("starting sync keys")
	start, end := n.keysRange()
	for i := start; i <= end; i++ {
		if i < n.id {
			n.syncKey(n.predecessor.Address, i)
		} else {
			n.syncKey(n.fingerTable[0].Address, i)
		}
	}
}

func (n *Node) syncSuccKeys() {
	if n.fingerTable[0].Address == "" || n.nSucc.Address == "" {
		return
	}
	start := n.id
	end := (n.nSucc.Id-n.fingerTable[0].Id)/2 + n.fingerTable[0].Id

	for i := start; i <= end; i++ {
		n.syncKey(n.fingerTable[0].Address, i)
	}
}

func (n *Node) syncPredKeys() {
	client := &client2.Client{}
	predOfPred, _ := client.Predecessor(n.predecessor.Address)
	if predOfPred.Endpoint == "" {
		return
	}
	start := (n.predecessor.Id-predOfPred.Id)/2 + predOfPred.Id
	end := n.predecessor.Id - 1

	for i := start; i <= end; i++ {
		n.syncKey(n.predecessor.Address, i)
	}
}

func (n *Node) isFingerSet(index int) bool {
	return n.fingerTable != nil && len(n.fingerTable) >= index && n.fingerTable[index].Address != ""
}

func (n *Node) findAimingNode(key int64) *models.NodeRepresentation {
	aimingNode := n.fingerTable[0]
	possibleNodes := append(n.fingerTable, n.predecessor)

	//take the smallest distance node
	smallestDistance := int64(math.Pow(2, float64(n.M))) + 1000
	for _, node := range possibleNodes {
		if node.Id == 0 || node.Address == "" {
			continue
		}
		currentDistance := distance(node.Id, key, n.M)
		if currentDistance < smallestDistance {
			smallestDistance = currentDistance
			aimingNode = node
		}
	}

	return &aimingNode
}

func (n *Node) printState() {
	now := time.Now()
	time := now.Unix()
	f, err := os.OpenFile("./state-"+fmt.Sprint(time), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("error to save state file")
	}

	defer f.Close()

	stateStr := "Successor: " + n.fingerTable[0].Address + ", " + fmt.Sprint(n.fingerTable[0].Id)
	stateStr += "\nPredecessor: " + n.predecessor.Address + ", " + fmt.Sprint(n.predecessor.Id)
	stateStr += "\nNSuccessor: " + n.nSucc.Address + ", " + fmt.Sprint(n.nSucc.Id)
	for i, finger := range n.fingerTable {
		stateStr += "\n Finger " + fmt.Sprint(i) + ": " + finger.Address + ", " + fmt.Sprint(finger.Id)
	}

	if _, err := f.WriteString(stateStr); err != nil {
		log.Println("unable to write state")
	}
}
