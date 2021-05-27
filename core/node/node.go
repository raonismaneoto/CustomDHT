package node

import (
	"errors"
	"github.com/raonismaneoto/CustomDHT/node/api"
	"github.com/raonismaneoto/CustomDHT/commons/grpc_api"
	"math"
	"time"
)

type NodeRepresentation struct {
	Id int64
	Address string
}

type Node struct {
	fingerTable []NodeRepresentation
	id int64
	address string
	storage map[int64][]byte
	predecessor NodeRepresentation
	nSucc NodeRepresentation
	m int
	replicationBuffer chan struct{key int64; data []byte}
}

func New(id int64) *Node{
	return &Node{id: id}
}

func (n *Node) Start(partner *NodeRepresentation) {
	n.replicationBuffer = make(chan struct{key int64; data []byte}, 50)

	n.Join(partner)

	periodicInvocation(n.checkSucc)
	periodicInvocation(n.stabilize)

	go n.syncReplicatedKeys()
}

func (n *Node) syncReplicatedKeys() {
	for msg := range n.replicationBuffer {
		n.storage[msg.key] = msg.data
	}
}

func (n *Node) checkSucc() {
	if len(n.fingerTable) == 0 {
		return
	}

	client := &api.Client{}
	_, err := client.Ping(n.fingerTable[0].Address)

	if err != nil {
		if n.nSucc.Address == "" {
			return
		}

		response := client.HandleNewPredecessor(n.nSucc.Address, NodeRepresentation{Id: n.id, Address: n.address})

		if ! response.Ok {
			panic("")
		}

		n.fingerTable[0] = NodeRepresentation{Id: n.nSucc.Id, Address: n.nSucc.Address }

		nSucc := client.Successor(n.nSucc.Address)

		n.nSucc = NodeRepresentation {Id: nSucc.Id, Address: n.address}

		n.syncKeys()
	}
}

func (n *Node) Join(partner *NodeRepresentation) {
	if partner == nil {
		return
	}

	client := api.Client{}
	n.startFingerTable(partner, client)
	// notify other nodes to update their predecessors and finger table
	nodeRepresentation := NodeRepresentation {Id: n.id, Address: n.address}
	newPredResponse := client.HandleNewPredecessor(n.fingerTable[0].Address, nodeRepresentation)
	if !newPredResponse.Ok {
		panic("Successor did not accept new predecessor")
	}
	newSuccResponse := client.HandleNewSuccessor(n.predecessor.Address, nodeRepresentation)
	if !newSuccResponse.Ok {
		panic("Predecessor did not accept new sucessor")
	}
	// get the data
	n.syncKeys()
}

func (n *Node) RepSave(message struct{key int64; data []byte}) {
	go func() {
		n.replicationBuffer <- message
	}()
}

func (n *Node) Save(key int64, value []byte) {
	client := api.Client{}
	if n.mustKeyBeInNode(key) {
		n.storage[key] = value
		inflectionPoint := (n.id - n.predecessor.Id)/2 + n.predecessor.Id
		if key >= inflectionPoint {
			client.RepSave(n.fingerTable[0].Address, key, value)
		} else {
			client.RepSave(n.predecessor.Address, key, value)
		}
		return
	}

	// else the request must be passed to the responsible node
	response := n.Query(key)
	client.Save(response.ResponsibleNodeEndpoint, key, value)
}

func (n *Node) Delete(key int64) {
	delete(n.storage, key)

	client := api.Client{}
	inflectionPoint := (n.id - n.predecessor.Id)/2 + n.predecessor.Id
	if key >= inflectionPoint {
		client.Delete(n.fingerTable[0].Address, key)
	} else {
		client.Delete(n.predecessor.Address, key)
	}
}

func (n *Node) Query(key int64) grpc_api.QueryResponse{
	if n.mustKeyBeInNode(key) {
		data, ok := n.storage[key]

		if !ok {
			return grpc_api.QueryResponse{
				Data: []byte{},
				ResponsibleNodeEndpoint: n.address,
				ResponsibleNodeId: n.id,
			}
		}

		return grpc_api.QueryResponse{
			Data: data,
			ResponsibleNodeEndpoint: n.address,
			ResponsibleNodeId: n.id,
		}
	}

	var aimingNode *NodeRepresentation
	if n.id > key {
		//take the smallest distance node
		smallestDistance := int64(math.Pow(2, float64(n.m))) + 1000
		for _, finger := range n.fingerTable {
			currentDistance := distance(finger.Id, key, n.m)
			if currentDistance < smallestDistance {
				smallestDistance = currentDistance
				aimingNode = &finger
			}

		}
	} else {
		for _, finger := range n.fingerTable {
			aimingNode = &finger
			if finger.Id > key {
				break
			}
		}
	}

	if aimingNode != nil {
		client := api.Client{}
		return *client.Query(aimingNode.Address, key)
	}

	return grpc_api.QueryResponse{
		Data: []byte{},
		ResponsibleNodeEndpoint: "",
		ResponsibleNodeId: -1,
	}
}

func distance(i int64, j int64, m int) int64 {
	if j == i { return 0 }
	if j > i {
		return j - i
	}
	return int64(math.Pow(2, float64(m))) - i + j
}

func (n *Node) HandleNewSuccessor(newSucc NodeRepresentation) error {
	client := &api.Client{}
	if newSucc.Id > n.fingerTable[0].Id {
		_, err := client.Ping(n.fingerTable[0].Address)
		if err == nil { return errors.New("invalid successor") }
	}

	n.fingerTable[0] = newSucc
	nNSuccData := client.Successor(newSucc.Address)
	n.nSucc = NodeRepresentation {Id: nNSuccData.Id, Address: nNSuccData.Endpoint}

	go n.syncSuccKeys()

	return nil
}

func (n *Node) HandleNewPredecessor(nPred NodeRepresentation) error {
	client := &api.Client{}
	if nPred.Id < n.predecessor.Id {
		_, err := client.Ping(n.predecessor.Address)
		if err == nil { return errors.New("invalid predecessor") }
	}

	n.predecessor = nPred
	// check if this goroutine will die after this function return statement
	go n.syncPredKeys()

	return nil
}

func (n *Node) Successor() NodeRepresentation {
	return n.fingerTable[0]
}

func (n *Node) Predecessor() NodeRepresentation {
	return n.predecessor
}

func (n *Node) stabilize() {
	for i := 1; i < n.m; i++ {
		currNodeInfo := n.Query((n.id + int64(math.Pow(2, float64(i-1))))%int64(math.Pow(2, float64(n.m))))
		if currNodeInfo.ResponsibleNodeEndpoint == "" {
			continue
		}
		n.fingerTable[i] = NodeRepresentation {Id: currNodeInfo.ResponsibleNodeId, Address:currNodeInfo.ResponsibleNodeEndpoint }
	}
}

func (n *Node) startFingerTable(partner *NodeRepresentation, client api.Client) {
	n.fingerTable = []NodeRepresentation{}
	succInfo := client.Query(partner.Address, n.id)
	n.fingerTable[0] = NodeRepresentation {Id: succInfo.ResponsibleNodeId, Address:succInfo.ResponsibleNodeEndpoint }

	nSuccInfo := client.Successor(n.fingerTable[0].Address)
	n.nSucc = NodeRepresentation {Id: nSuccInfo.Id, Address: nSuccInfo.Endpoint}

	//set predecessor before other entries of the finger table because it is needed in the query func
	predecessor := client.Predecessor(n.fingerTable[0].Address)
	n.predecessor = NodeRepresentation {Id: predecessor.Id, Address: predecessor.Endpoint}

	for i := 1; i < n.m; i++ {
		currNodeInfo := n.Query((n.id + int64(math.Pow(2, float64(i-1))))%int64(math.Pow(2, float64(n.m))))
		n.fingerTable[i] = NodeRepresentation {Id: currNodeInfo.ResponsibleNodeId, Address:currNodeInfo.ResponsibleNodeEndpoint }
	}
}

func (n *Node) mustKeyBeInNode(key int64) bool{
	return n.predecessor.Id == 0 || (key >= n.predecessor.Id && key < n.id)
}

func (n *Node) keysRange() (int64, int64) {
	client := &api.Client{}
	predOfPred := client.Predecessor(n.predecessor.Address)
	start := (n.predecessor.Id - predOfPred.Id)/2 + predOfPred.Id
	end := (n.nSucc.Id - n.fingerTable[0].Id)/2 + n.fingerTable[0].Id
	return start, end
}

func (n *Node) syncKey(address string, key int64) {
	client := api.Client{}
	response := client.Query(address, key)
	if response.Data != nil && len(response.Data) > 0 {
		n.storage[key] = response.Data
	}
}

func (n *Node) syncKeys() {
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
	start := n.id
	end := (n.nSucc.Id - n.fingerTable[0].Id)/2 + n.fingerTable[0].Id

	for i := start; i <= end; i++ {
		n.syncKey(n.fingerTable[0].Address, i)
	}
}

func (n *Node) syncPredKeys() {
	client := &api.Client{}
	predOfPred := client.Predecessor(n.predecessor.Address)
	start := (n.predecessor.Id - predOfPred.Id)/2 + predOfPred.Id
	end := n.predecessor.Id - 1

	for i := start; i <= end; i++ {
		n.syncKey(n.predecessor.Address, i)
	}
}

func periodicInvocation(f func()) {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				f()
			}
		}
	}()
}
