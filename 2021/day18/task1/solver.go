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

		root, reduced := reduce(root)
		if !reduced {
			fmt.Println("reduce not required")
		} else {
			fmt.Println("reduced to: ", root)
		}
	}

	return 0, root, nil
}

func reduce(root tree.Node) (tree.Node, bool) {
	var exploded, splitten, reduced bool
	for {
		root, exploded = explode(root)
		if exploded {
			reduced = true
			continue
		}

		root, splitten = split(root)
		if !splitten {
			return root, reduced
		}

		reduced = true
	}
}

func split(root tree.Node) (tree.Node, bool) {
	// todo:
	// 1. find the leftest number above 9 using reverse tree traversal
	// 2. replace that value node with pair node
	return nil, false
}

func explode(root tree.Node) (tree.Node, bool) {
	return nil, false
}

func sum(root, n tree.Node) tree.Node {
	return &tree.PairNode{
		Parent:     nil,
		LeftChild:  root,
		RightChild: n,
	}
}
