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
	nSucc struct {id int64; address string}
	m int
}

func (n *Node) Start(partner *Node, m int) {
	n.Join(partner, m)

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				n.checkSucc()
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-ticker.C:
				n.stabilize(m)
			}
		}
	}()
}

func (n *Node) checkSucc() {
	client := &api.Client{}
	_, err := client.Ping(n.fingerTable[0].address)

	if err != nil {
		_, err := client.HandleNewPredecessor(n.nSucc.address, struct {id int64; address string} {id: n.id, address: n.address})

		if err != nil {}

		n.fingerTable[0] = struct {
			id      int64
			address string
		}{id: n.nSucc.id, address: n.nSucc.address }

		nSucc := client.Successor(n.nSucc.address)

		n.nSucc = struct {
			id      int64
			address string
		}{id: nSucc.Id, address: n.address}

		n.syncKeys()
	}
}

func (n *Node) Join(partner *Node, m int) {
	n.id = int64(helpers.GetHash(n.address, m))
	client := api.Client{}
	n.startFingerTable(partner, m, client)
	// notify other nodes to update their predecessors and finger table
	nodeRepresentation := struct {id int64; address string} {id: n.id, address: n.address}
	client.HandleNewSuccessor(nodeRepresentation, n.fingerTable[0].address)
	client.HandleNewPredecessor(nodeRepresentation, n.predecessor.address)
	// get the data
	n.syncKeys()
}

func (n *Node) Leave() {

}

func (n *Node) Save(key int64, value []byte) {
	if n.mustKeyBeInNode(key) {
		// if the key already exists it must be updated in the replication nodes
		n.storage[key] = value
		return
	}

	// else the request must be passed to the responsible node
}

func (n *Node) Delete(key int64) {
	// delete in the replication nodes
	delete(n.storage, key)
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

	var aimingNode struct {id int64; address string}
	if n.id > key {
		//take the smallest distance node
		smallestDistance := int64(math.Pow(2, float64(n.m))) + 1000
		for _, finger := range n.fingerTable {
			currentDistance := distance(finger.id, key, n.m)
			if currentDistance < smallestDistance {
				smallestDistance = currentDistance
				aimingNode = finger
			}

		}
	} else {
		for _, finger := range n.fingerTable {
			aimingNode = finger
			if finger.id > key {
				break
			}
		}
	}

	client := api.Client{}
	return *client.Query(aimingNode.address, key)
}

func distance(i int64, j int64, m int) int64 {
	if j == i { return 0 }
	if j > i {
		return j - i
	}
	return int64(math.Pow(2, float64(m))) - i + j
}

func (n *Node) HandleNewSuccessor() {

}

func (n *Node) HandleNewPredecessor() {

}

func (n *Node) Successor() {

}

func (n *Node) Predecessor() {

}

func (n *Node) stabilize(m int) {
	for i := 1; i < m; i++ {
		currNodeInfo := n.Query((n.id + int64(math.Pow(2, float64(i-1))))%int64(math.Pow(2, float64(m))))
		n.fingerTable[i] = struct {
			id      int64
			address string
		}{id: currNodeInfo.ResponsibleNodeId, address:currNodeInfo.ResponsibleNodeEndpoint }
	}
}

func (n *Node) startFingerTable(partner *Node, m int, client api.Client) {
	n.fingerTable = []struct {id int64; address string}{}
	succInfo := client.Query(partner.address, n.id)
	n.fingerTable[0] = struct {
		id      int64
		address string
	}{id: succInfo.ResponsibleNodeId, address:succInfo.ResponsibleNodeEndpoint }

	nSuccInfo := client.Successor(n.fingerTable[0].address)
	n.nSucc = struct {
		id int64
		address string
	}{id: nSuccInfo.Id, address: nSuccInfo.Endpoint}

	//set predecessor before other entries of the finger table because it is needed in the query func
	predecessor := client.Predecessor(n.fingerTable[0].address)
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

func (n *Node) mustKeyBeInNode(key int64) bool{
	return n.predecessor.id == 0 || (key >= n.predecessor.id && key < n.id)
}

func (n *Node) keysRange() (int64, int64) {
	client := &api.Client{}
	predOfPred := client.Predecessor(n.predecessor.address)
	start := (n.predecessor.id - predOfPred.Id)/2 + predOfPred.Id
	end := (n.nSucc.id - n.fingerTable[0].id)/2 + n.fingerTable[0].id
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
			n.syncKey(n.predecessor.address, i)
		} else {
			n.syncKey(n.fingerTable[0].address, i)
		}
	}
}
