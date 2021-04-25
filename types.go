package merkle

type Tree struct {
	root *Node
}

type Node struct {
	height int
	hash   []byte
	left   *Node
	right  *Node
}
