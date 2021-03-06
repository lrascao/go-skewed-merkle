package skewed_merkle

type TreeSide int

const (
	Left  TreeSide = iota
	Right          = iota
)

type NotFoundError struct {
}

func (a NotFoundError) Error() string {
	return "not found"
}

type Tree struct {
	root *Node
}

type Node struct {
	height int
	hash   []byte
	left   *Node
	right  *Node
}

type Proof struct {
	side TreeSide
	hash []byte
}
