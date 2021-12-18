package tree

import (
	"strconv"
	"strings"
)

func ParseNodes(lines []string) ([]Node, error) {
	nodes := make([]Node, 0, len(lines))
	for _, l := range lines {
		node, err := ParseNode(l)
		if err != nil {
			return nil, err
		}

		nodes = append(nodes, node)
	}

	return nodes, nil
}

func ParseNode(l string) (Node, error) {
	if !strings.ContainsAny(l, "[]") {
		val, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		return &ValueNode{
			Parent: nil,
			Value:  val,
		}, nil
	}

	comaPos := -1
	depth := 0

	chars := []rune(l)
	for i := range chars {
		if chars[i] == '[' {
			depth++
		}

		if chars[i] == ']' {
			depth--
		}

		if chars[i] == ',' && depth == 1 {
			comaPos = i
			break
		}
	}

	left, right := string(chars[1:comaPos]), string(chars[comaPos+1:len(chars)-1])

	currentNode := &PairNode{
		Parent:     nil,
		LeftChild:  nil,
		RightChild: nil,
	}

	leftNode, err := ParseNode(left)
	if err != nil {
		return nil, err
	}

	rightNode, err := ParseNode(right)
	if err != nil {
		return nil, err
	}

	leftNode.SetParent(currentNode)
	rightNode.SetParent(currentNode)

	currentNode.LeftChild = leftNode
	currentNode.RightChild = rightNode

	return currentNode, nil
}
