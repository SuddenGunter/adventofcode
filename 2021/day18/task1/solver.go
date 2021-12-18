package task1

import (
	"aoc-2021-day18/input"
	"aoc-2021-day18/tree"
	"fmt"
)

func Solve(data input.Data) (int, tree.Node, error) {
	root := data.Numbers[0]
	for _, n := range data.Numbers[1:] {
		fmt.Printf("%v + %v = ", root, n)
		root = sum(root, n)
		fmt.Printf("%v\n", root)
	}

	return 0, root, nil
}

func sum(root tree.Node, n tree.Node) tree.Node {
	return &tree.PairNode{
		Parent:     nil,
		LeftChild:  root,
		RightChild: n,
	}
}
