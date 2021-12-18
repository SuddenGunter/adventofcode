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
	leftmost, err := findLeftmostVal(root, func(val *tree.ValueNode) bool {
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

type filterVal func(val *tree.ValueNode) bool

func findLeftmostVal(root tree.Node, fn filterVal) (*tree.ValueNode, error) {
	if asVal, ok := root.(*tree.ValueNode); ok {
		found := fn(asVal)
		if found {
			return asVal, nil
		}

		return nil, errNotFound
	}

	s := stack{data: make([]stackEntry, 0)}
	s.push(stackEntry{
		node: root,
	})

	results := stack{data: make([]stackEntry, 0)}

	for !s.isEmpty() {
		entry := s.pop()

		if asVal, ok := entry.node.(*tree.ValueNode); ok {
			found := fn(asVal)
			if found {
				results.push(stackEntry{
					node:  asVal,
					depth: entry.depth,
				})
			}

			continue
		}

		asPair, ok := entry.node.(*tree.PairNode)
		if !ok {
			return nil, errors.New("unknown node type")
		}

		if asPair.LeftChild != nil {
			s.push(stackEntry{
				node:  asPair.LeftChild,
				depth: entry.depth + 1,
			})
		}
		if asPair.RightChild != nil {
			s.push(stackEntry{
				node:  asPair.RightChild,
				depth: entry.depth + 1,
			})
		}
	}

	if !results.isEmpty() {
		res := results.pop()
		asVal, ok := res.node.(*tree.ValueNode)
		if !ok {
			return nil, fmt.Errorf("expected *tree.ValueNode, received: %v", reflect.TypeOf(res))
		}

		return asVal, nil
	}

	return nil, errNotFound
}

type filterPair func(val *tree.PairNode, depth int) bool

func findLeftmostPair(root tree.Node, fn filterPair) (*tree.PairNode, error) {
	if _, ok := root.(*tree.ValueNode); ok {
		return nil, errNotFound
	}

	s := stack{data: make([]stackEntry, 0)}
	s.push(stackEntry{
		node:  root,
		depth: 0,
	})

	results := stack{data: make([]stackEntry, 0)}

	for !s.isEmpty() {
		entry := s.pop()

		if _, ok := entry.node.(*tree.ValueNode); ok {
			continue
		}

		asPair, ok := entry.node.(*tree.PairNode)

		if !ok {
			return nil, errors.New("unknown node type")
		}

		found := fn(asPair, entry.depth)
		if found {
			results.push(stackEntry{
				node:  asPair,
				depth: entry.depth,
			})
		}

		if asPair.LeftChild != nil {
			s.push(stackEntry{
				node:  asPair.LeftChild,
				depth: entry.depth + 1,
			})
		}

		if asPair.RightChild != nil {
			s.push(stackEntry{
				node:  asPair.RightChild,
				depth: entry.depth + 1,
			})
		}
	}

	if !results.isEmpty() {
		res := results.pop()

		asPair, ok := res.node.(*tree.PairNode)
		if !ok {
			return nil, fmt.Errorf("expected *tree.PairNode, received: %v", reflect.TypeOf(res))
		}

		return asPair, nil
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
