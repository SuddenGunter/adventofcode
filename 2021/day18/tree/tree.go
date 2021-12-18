package tree

import (
	"fmt"
	"strconv"
)

type PairNode struct {
	Parent Node

	LeftChild  Node
	RightChild Node
}

func (n *PairNode) String() string {
	return fmt.Sprintf("[%v,%v]", n.LeftChild.String(), n.RightChild.String())
}

func (n *PairNode) GetParent() Node {
	return n.Parent
}

type ValueNode struct {
	Parent Node

	Value int
}

func (n *ValueNode) String() string {
	return strconv.Itoa(n.Value)
}

func (n *ValueNode) GetParent() Node {
	return n.Parent
}

type Node interface {
	GetParent() Node
	fmt.Stringer
}
