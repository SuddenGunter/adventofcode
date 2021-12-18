package task1

import (
	"aoc-2021-day18/input"
	"aoc-2021-day18/tree"
	"errors"
	"fmt"
	"reflect"
)

func Solve(data input.Data) (int, tree.Node, error) {
	root := data.Numbers[0]
	for _, n := range data.Numbers[1:] {
		fmt.Printf("%v + %v = ", root, n)

		root = sum(root, n)
		fmt.Printf("%v\n", root)

		res, reduced, err := reduce(root)
		if err != nil {
			return 0, nil, err
		}

		root = res

		if !reduced {
			fmt.Println("reduce not required")
		} else {
			fmt.Println("reduced to: ", root.String())
		}
	}

	return 0, root, nil
}

func reduce(root tree.Node) (tree.Node, bool, error) {
	var exploded, splitten, reduced bool
	var err error

	for {
		root, exploded = explode(root)
		if exploded {
			reduced = true
			continue
		}

		root, splitten, err = split(root)
		if err != nil {
			return nil, false, err
		}

		if !splitten {
			return root, reduced, nil
		}

		reduced = true
	}
}

func split(root tree.Node) (tree.Node, bool, error) {
	leftmost, err := findLeftmost(root, func(val *tree.ValueNode) bool {
		return val.Value >= 10
	})

	switch {
	case errors.Is(err, errNotFound):
		return root, false, nil
	case !errors.Is(err, nil):
		return nil, false, err
	default:
	}

	left := &tree.ValueNode{
		Value: leftmost.Value / 2,
	}

	right := &tree.ValueNode{
		Value: leftmost.Value - left.Value,
	}

	replacement := &tree.PairNode{
		Parent:     leftmost.Parent,
		LeftChild:  left,
		RightChild: right,
	}

	left.SetParent(replacement)
	right.SetParent(replacement)

	asPair, ok := leftmost.Parent.(*tree.PairNode)
	if !ok {
		return nil, false, fmt.Errorf("expected *tree.ValueNode, received: %v", reflect.TypeOf(leftmost.Parent))
	}

	if asPair.LeftChild == leftmost {
		asPair.LeftChild = replacement
	} else {
		asPair.RightChild = replacement
	}

	return root, true, nil
}

type filter func(val *tree.ValueNode) bool

func findLeftmost(root tree.Node, fn filter) (*tree.ValueNode, error) {
	if asVal, ok := root.(*tree.ValueNode); ok {
		return asVal, nil
	}

	s := stack{data: make([]tree.Node, 0)}
	s.push(root)

	results := stack{data: make([]tree.Node, 0)}

	for !s.isEmpty() {
		node := s.pop()

		if asVal, ok := node.(*tree.ValueNode); ok {
			found := fn(asVal)
			if found {
				results.push(asVal)
			}

			continue
		}

		asPair, ok := node.(*tree.PairNode)
		if !ok {
			return nil, errors.New("unknown node type")
		}

		if asPair.LeftChild != nil {
			s.push(asPair.LeftChild)
		}
		if asPair.RightChild != nil {
			s.push(asPair.RightChild)
		}
	}

	if !results.isEmpty() {
		res := results.pop()
		asVal, ok := res.(*tree.ValueNode)
		if !ok {
			return nil, fmt.Errorf("expected *tree.ValueNode, received: %v", reflect.TypeOf(res))
		}

		return asVal, nil
	}

	return nil, errNotFound
}

func explode(root tree.Node) (tree.Node, bool) {
	return root, false
}

func sum(root, n tree.Node) tree.Node {
	return &tree.PairNode{
		Parent:     nil,
		LeftChild:  root,
		RightChild: n,
	}
}

var errNotFound = errors.New("not found")
