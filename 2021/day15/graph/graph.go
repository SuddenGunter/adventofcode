package graph

type Node struct {
	ID       ID
	Adjacent []*Node
	Weight   int8
}

func (n *Node) Equals(node *Node) bool {
	return n.ID == node.ID
}

type Graph struct {
	Root *Node
}

// ID is a position of the node in the original input array.
type ID struct {
	X, Y int
}
