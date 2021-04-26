package node

type Node struct {
	fingerTable map[int]string
	id string
	address string
	storage map[int][]byte
	successor struct {id int; address string}
	predecessor struct {id int; address string}
}

func (n *Node) Start() {
	n.Join()
}

func (n *Node) Join() {

}

func (n *Node) Leave() {

}

func (n *Node) Save(key int, value []byte) {

}

func (n *Node) Delete(key int) {

}

func (n *Node) Query(key int) {

}

func (n *Node) HandleChurn() {

}

func (n *Node) HandleNewSuccessor() {

}

func (n *Node) HandleNewPredecessor() {

}


