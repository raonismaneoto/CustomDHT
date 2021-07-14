package node

import (
	"errors"
	"github.com/raonismaneoto/CustomDHT/commons/grpc_api"
	"log"
	"math"
	"time"
)

type NodeRepresentation struct {
	Id      int64
	Address string
}

type Node struct {
	fingerTable       []NodeRepresentation
	id                int64
	address           string
	storage           map[int64][]byte
	predecessor       NodeRepresentation
	nSucc             NodeRepresentation
	m                 int
	replicationBuffer chan struct {
		key  int64
		data []byte
	}
}

func New(id int64, address string, m int) *Node {
	return &Node{id: id, address: address, m: m}
}

func (n *Node) Start(partner *NodeRepresentation) {
	log.Println("Starting node: %v", n.id)
	log.Println("creating replicationBuffer")
	n.replicationBuffer = make(chan struct {
		key  int64
		data []byte
	}, 50)
	log.Println("going to call Join")
	n.fingerTable = make([]NodeRepresentation, n.m, n.m)
	n.storage = make(map[int64][]byte)
	n.predecessor = NodeRepresentation{
		Id:      0,
		Address: "",
	}
	n.nSucc = NodeRepresentation{
		Id:      0,
		Address: "",
	}
	n.Join(partner)

	periodicInvocation(n.checkSucc, 5)
	periodicInvocation(n.stabilize, 5)

	go n.syncReplicatedKeys()

	periodicInvocation(n.logStoredKeys, 60)
}

func (n *Node) syncReplicatedKeys() {
	for msg := range n.replicationBuffer {
		n.storage[msg.key] = msg.data
	}
}

func (n *Node) checkSucc() {
	if n.fingerTable == nil || len(n.fingerTable) == 0 {
		return
	}

	client := &Client{}
	_, err := client.Ping(n.fingerTable[0].Address)

	if err != nil {
		if n.nSucc.Address == "" {
			return
		}

		response, err := client.HandleNewPredecessor(n.nSucc.Address, NodeRepresentation{Id: n.id, Address: n.address})

		if err != nil || !response.Ok {
			log.Println("handle new predecessor failed for nSucc.")
			return
		}

		n.fingerTable[0] = NodeRepresentation{Id: n.nSucc.Id, Address: n.nSucc.Address}

		nSucc, err := client.Successor(n.nSucc.Address)
		if err != nil {
			log.Println("unable to reach nSucc in checkSucc. Err: " + err.Error())
		}

		n.nSucc = NodeRepresentation{Id: nSucc.Id, Address: n.address}

		n.syncKeys()
	}
}

func (n *Node) Join(partner *NodeRepresentation) {
	if partner == nil {
		return
	}

	client := Client{}
	log.Println("Starting finger table")
	n.startFingerTable(partner, client)
	// notify other nodes to update their predecessors and finger table
	log.Println("Starting join notifications")
	nodeRepresentation := NodeRepresentation{Id: n.id, Address: n.address}
	newPredResponse, err := client.HandleNewPredecessor(n.fingerTable[0].Address, nodeRepresentation)
	if err != nil || !newPredResponse.Ok {
		log.Println("Successor did not accept new predecessor" + err.Error())
		panic("Successor did not accept new predecessor")
	}
	newSuccResponse, err := client.HandleNewSuccessor(n.predecessor.Address, nodeRepresentation, NodeRepresentation{Id: 0, Address: ""})
	if err != nil || !newSuccResponse.Ok {
		log.Println("Predecessor did not accept new sucessor" + err.Error())
		panic("Predecessor did not accept new sucessor")
	}
	// get the data
	n.syncKeys()
}

func (n *Node) RepSave(key int64, data []byte) {
	go func() {
		n.replicationBuffer <- struct {
			key  int64
			data []byte
		}{key: key, data: data}
	}()
}

func (n *Node) Save(key int64, value []byte) error {
	client := Client{}
	if n.mustKeyBeInNode(key) {
		n.storage[key] = value
		inflectionPoint := (n.id-n.predecessor.Id)/2 + n.predecessor.Id
		if key >= inflectionPoint {
			client.RepSave(n.fingerTable[0].Address, key, value)
		} else {
			client.RepSave(n.predecessor.Address, key, value)
		}
		return nil
	}

	// else the request must be passed to the responsible node
	response := n.Query(key)
	_, err := client.Save(response.ResponsibleNodeEndpoint, key, value)
	return err
}

func (n *Node) Delete(key int64) {
	delete(n.storage, key)

	client := Client{}
	inflectionPoint := (n.id-n.predecessor.Id)/2 + n.predecessor.Id
	if key >= inflectionPoint {
		client.Delete(n.fingerTable[0].Address, key)
	} else {
		client.Delete(n.predecessor.Address, key)
	}
}

func (n *Node) Query(key int64) grpc_api.QueryResponse {
	if n.mustKeyBeInNode(key) {
		data, ok := n.storage[key]

		if !ok {
			log.Println("Key" + string(key) + " not found.")
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

	var aimingNode = NodeRepresentation{
		Id:      0,
		Address: "",
	}

	if n.id > key {
		//take the smallest distance node
		smallestDistance := int64(math.Pow(2, float64(n.m))) + 1000
		for _, finger := range n.fingerTable {
			currentDistance := distance(finger.Id, key, n.m)
			if currentDistance < smallestDistance {
				smallestDistance = currentDistance
				aimingNode = finger
			}

		}
	} else {
		for _, finger := range n.fingerTable {
			aimingNode = finger
			if finger.Id > key {
				break
			}
		}
	}

	if aimingNode.Address != "" {
		log.Println("key not found in node, going to forward the query to:")
		log.Println("nodeAddress: " + aimingNode.Address)
		client := Client{}
		return *client.Query(aimingNode.Address, key)
	}

	log.Println("Key" + string(key) + " not found.")

	return grpc_api.QueryResponse{
		Data:                    []byte{},
		ResponsibleNodeEndpoint: "",
		ResponsibleNodeId:       -1,
	}
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

func (n *Node) HandleNewSuccessor(newSucc NodeRepresentation, nNSucc NodeRepresentation) error {
	client := &Client{}
	if n.fingerTable[0].Address != "" && newSucc.Id > n.fingerTable[0].Id {
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

func (n *Node) HandleNewPredecessor(nPred NodeRepresentation) error {
	client := &Client{}
	if nPred.Id < n.predecessor.Id {
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

func (n *Node) Successor() (*NodeRepresentation, error) {
	if n.fingerTable == nil || len(n.fingerTable) == 0 || n.fingerTable[0].Address == "" {
		return nil, errors.New("There is no successor")
	}

	return &n.fingerTable[0], nil
}

func (n *Node) Predecessor() (NodeRepresentation, error) {
	if n.predecessor.Address == "" {
		return n.predecessor, errors.New("There is no predecessor")
	}

	return n.predecessor, nil
}

func (n *Node) stabilize() {
	for i := 1; i < n.m; i++ {
		currNodeInfo := n.Query((n.id + int64(math.Pow(2, float64(i-1)))) % int64(math.Pow(2, float64(n.m))))
		if currNodeInfo.ResponsibleNodeEndpoint == "" {
			continue
		}
		n.fingerTable[i] = NodeRepresentation{Id: currNodeInfo.ResponsibleNodeId, Address: currNodeInfo.ResponsibleNodeEndpoint}
	}
}

func (n *Node) startFingerTable(partner *NodeRepresentation, client Client) {
	log.Println("querying succ info in startFingerTable")
	succInfo := client.Query(partner.Address, n.id)
	succ := NodeRepresentation{Id: succInfo.ResponsibleNodeId, Address: succInfo.ResponsibleNodeEndpoint}
	n.fingerTable[0] = succ

	nSuccInfo, err := client.Successor(n.fingerTable[0].Address)
	if err != nil {
		//if the partner does not have successor it does not have predecessor either
		log.Println("Unable to find nsucc on finger table startup. Err: %v", err)
		n.predecessor = succ
	} else {
		n.nSucc = NodeRepresentation{Id: nSuccInfo.Id, Address: nSuccInfo.Endpoint}
		predecessor, err := client.Predecessor(n.fingerTable[0].Address)
		if err != nil {
			log.Println(err.Error())
		}
		n.predecessor = NodeRepresentation{Id: predecessor.Id, Address: predecessor.Endpoint}
	}

	for i := 1; i < n.m; i++ {
		currNodeInfo := n.Query((n.id + int64(math.Pow(2, float64(i-1)))) % int64(math.Pow(2, float64(n.m))))
		n.fingerTable[i] = NodeRepresentation{Id: currNodeInfo.ResponsibleNodeId, Address: currNodeInfo.ResponsibleNodeEndpoint}
	}
}

func (n *Node) mustKeyBeInNode(key int64) bool {
	return n.predecessor.Id == 0 || (key >= n.predecessor.Id && key < n.id)
}

func (n *Node) keysRange() (int64, int64) {
	client := &Client{}
	predOfPred, _ := client.Predecessor(n.predecessor.Address)
	start := (n.predecessor.Id-predOfPred.Id)/2 + predOfPred.Id
	end := (n.nSucc.Id-n.fingerTable[0].Id)/2 + n.fingerTable[0].Id
	return start, end
}

func (n *Node) syncKey(address string, key int64) {
	client := Client{}
	response := client.Query(address, key)
	if response.Data != nil && len(response.Data) > 0 {
		n.storage[key] = response.Data
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
	client := &Client{}
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

func periodicInvocation(f func(), secs int) {
	go func() {
		ticker := time.NewTicker(time.Duration(secs) * time.Second)
		for {
			select {
			case <-ticker.C:
				f()
			}
		}
	}()
}

func (n *Node) logStoredKeys() {
	msg := "Stored keys: "
	for k, _ := range n.storage {
		msg += "" + string(k) + " "
	}
	log.Println(msg)
}
