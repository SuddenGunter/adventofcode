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

func (n *PairNode) SetParent(parent Node) {
	n.Parent = parent
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

func (n *ValueNode) SetParent(parent Node) {
	n.Parent = parent
}

type Node interface {
	GetParent() Node
	SetParent(n Node)
	fmt.Stringer
}
