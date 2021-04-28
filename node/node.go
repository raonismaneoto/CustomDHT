package node

import (
	"github.com/raonismaneoto/CustomDHT/helpers"
	"github.com/raonismaneoto/CustomDHT/node/api"
	"github.com/raonismaneoto/CustomDHT/node/api/grpc_api"
	"math"
	"time"
)

type Node struct {
	fingerTable []struct {id int64; address string}
	id int64
	address string
	storage map[int64][]byte
	predecessor struct {id int64; address string}
}

func (n *Node) Start(partner *Node, m int) {
	n.Join(partner, m)

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-ticker.C:
				n.stabilize()
			}
		}
	}()
}

func (n *Node) Join(partner *Node, m int) {
	n.id = int64(helpers.GetHash(n.address, m))
	client := api.Client{}
	n.startFingerTable(partner, m, client)
	// notify other nodes to update their predecessors and finger table
	nodeRepresentation := struct {id int64; address string} {id: n.id, address: n.address}
	client.Notify(nodeRepresentation, n.fingerTable[0])
	client.Notify(nodeRepresentation, n.predecessor)
	// get the data
	client
}

func (n *Node) Leave() {

}

func (n *Node) Save(key int64, value []byte) {
	if n.isKeyInRange(key) {
		n.storage[key] = value
		return
	}

	// else the request must be passed to the responsible node
}

func (n *Node) Delete(key int64) {
	delete(n.storage, key)
}

func (n *Node) Query(key int64) grpc_api.QueryResponse{
	if n.isKeyInRange(key) {
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

	var aimingNode struct {id int64; address string}
	for _, finger := range n.fingerTable {
		if finger.id > key {
			break
		}
		aimingNode = finger
	}

	client := api.Client{}
	return *client.Query(aimingNode.address, key)
}

func (n *Node) HandleNotification() {

}

func (n *Node) HandleChurn() {

}

func (n *Node) HandleNewSuccessor() {

}

func (n *Node) HandleNewPredecessor() {

}

func (n *Node) Successor() {

}

func (n *Node) Predecessor() {

}

func (n *Node) stabilize() {
	// here we want to check for each entry in the finger table
}

func (n *Node) startFingerTable(partner *Node, m int, client api.Client) {
	n.fingerTable = []struct {id int64; address string}{}
	succInfo := client.Query(partner.address, n.id)
	n.fingerTable[0] = struct {
		id      int64
		address string
	}{id: succInfo.ResponsibleNodeId, address:succInfo.ResponsibleNodeEndpoint }

	//set predecessor before other entries of the finger table because it is needed in the query func
	predecessor := client.Predecessor(n.fingerTable[0].address, n.fingerTable[0].id)
	n.predecessor = struct {
		id      int64
		address string
	}{id: predecessor.Id, address: predecessor.Endpoint}

	for i := 1; i < m; i++ {
		currNodeInfo := n.Query((n.id + int64(math.Pow(2, float64(i-1))))%int64(math.Pow(2, float64(m))))
		n.fingerTable[i] = struct {
			id      int64
			address string
		}{id: currNodeInfo.ResponsibleNodeId, address:currNodeInfo.ResponsibleNodeEndpoint }
	}
}

func (n *Node) isKeyInRange(key int64) bool{
	return n.predecessor.id == 0 || (key >= n.predecessor.id && key < n.id)
}

