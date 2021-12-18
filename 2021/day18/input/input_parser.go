package input

import (
	"aoc-2021-day18/tree"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Numbers []tree.Node
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	nodes := make([]tree.Node, 0, len(lines))
	for _, l := range lines {
		node, err := parseNode(l)
		if err != nil {
			return Data{}, err
		}

		nodes = append(nodes, node)
	}

	return Data{Numbers: nodes}, nil
}

func parseNode(l string) (tree.Node, error) {
	if !strings.ContainsAny(l, "[]") {
		val, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		return &tree.ValueNode{
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

	currentNode := &tree.PairNode{
		Parent:     nil,
		LeftChild:  nil,
		RightChild: nil,
	}

	leftNode, err := parseNode(left)
	if err != nil {
		return nil, err
	}

	rightNode, err := parseNode(right)
	if err != nil {
		return nil, err
	}

	currentNode.LeftChild = leftNode
	currentNode.RightChild = rightNode

	return currentNode, nil
}
