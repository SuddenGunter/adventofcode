package task2

import (
	"aoc-2021-day18/input"
	"aoc-2021-day18/tree"
	"errors"
	"fmt"
	"math"
	"reflect"
)

func Solve(file string) (int, tree.Node, error) {
	maxMagnitude := math.MinInt

	data, err := input.ParseInput(file)
	if err != nil {
		// todo: fix this mess
		panic(err)
	}

	for i := 0; i < len(data.Numbers); i++ {
		data, err = input.ParseInput(file)
		if err != nil {
			// todo: fix this mess
			panic(err)
		}

		for j := 0; j < len(data.Numbers); j++ {
			if i == j {
				continue
			}

			magn, _, err := calculate(data.Numbers[i], data.Numbers[j])
			if err != nil {
				// todo: fix
				panic(err)
			}

			if magn > maxMagnitude {
				maxMagnitude = magn
			}

			data, err = input.ParseInput(file)
			if err != nil {
				// todo: fix this mess
				panic(err)
			}

			magn, _, err = calculate(data.Numbers[j], data.Numbers[i])
			if err != nil {
				// todo: fix
				panic(err)
			}

			if magn > maxMagnitude {
				maxMagnitude = magn
			}

			data, err = input.ParseInput(file)
			if err != nil {
				// todo: fix this mess
				panic(err)
			}

		}
	}

	return maxMagnitude, nil, nil
}

func calculate(first, second tree.Node) (int, tree.Node, error) {
	first = sum(first, second)

	res, _, err := reduce(first)
	if err != nil {
		return 0, nil, err
	}

	first = res

	return magnitude(first), first, nil
}

func reduce(root tree.Node) (tree.Node, bool, error) {
	var exploded, splitten, reduced bool
	var err error

	for {
		root, exploded, err = explode(root)
		if err != nil {
			return nil, false, err
		}

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
		return nil, false, fmt.Errorf("expected *tree.PairNode, received: %v", reflect.TypeOf(leftmost.Parent))
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

func explode(root tree.Node) (tree.Node, bool, error) {
	leftmost, err := findLeftmostPair(root, func(val *tree.PairNode, depth int) bool {
		return depth >= 4
	})

	switch {
	case errors.Is(err, errNotFound):
		return root, false, nil
	case !errors.Is(err, nil):
		return nil, false, err
	default:
	}

	asVal, ok := leftmost.LeftChild.(*tree.ValueNode)
	if !ok {
		return nil, false, fmt.Errorf("expected *tree.ValueNode, received: %v", reflect.TypeOf(leftmost.Parent))
	}

	propagateLeft(asVal)

	asVal, ok = leftmost.RightChild.(*tree.ValueNode)
	if !ok {
		return nil, false, fmt.Errorf("expected *tree.ValueNode, received: %v", reflect.TypeOf(leftmost.Parent))
	}

	propagateRight(asVal)
	// todo: move left num left and right num right

	replacement := &tree.ValueNode{
		Parent: leftmost.Parent,
		Value:  0,
	}

	asPair, ok := leftmost.Parent.(*tree.PairNode)
	if !ok {
		return nil, false, fmt.Errorf("expected *tree.PairNode, received: %v", reflect.TypeOf(leftmost.Parent))
	}

	if asPair.LeftChild == leftmost {
		asPair.LeftChild = replacement
	} else {
		asPair.RightChild = replacement
	}

	return root, true, nil
}

func propagateRight(value *tree.ValueNode) {
	// todo can it work correctrly for interfaces??
	// todo: what if value parent == nil
	ancestors := make(map[tree.Node]struct{})
	ancestors[value.Parent] = struct{}{}
	ancestors[value] = struct{}{}

	next := value.Parent
	var subtreeToFindLeftest tree.Node
	for next != nil && subtreeToFindLeftest == nil {
		asPair, ok := next.(*tree.PairNode)
		if !ok {
			// todo err
			panic("panic")
		}

		if _, found := ancestors[asPair.RightChild]; found {
			ancestors[next] = struct{}{}
			next = next.GetParent()
			continue
		}

		subtreeToFindLeftest = asPair.RightChild
	}

	if subtreeToFindLeftest == nil {
		// we're at root and cannot find other ways to the right. That means we started with the rightest subtree.
		return
	}

	leftmost, err := findLeftmostVal(subtreeToFindLeftest, func(val *tree.ValueNode) bool {
		return true
	})

	if err != nil {
		// todo: handle
		panic("err")
	}

	leftmost.Value += value.Value
}

func propagateLeft(value *tree.ValueNode) {
	// todo can it work correctrly for interfaces??
	// todo: what if value parent == nil
	ancestors := make(map[tree.Node]struct{})
	ancestors[value.Parent] = struct{}{}
	ancestors[value] = struct{}{}

	next := value.Parent
	var subtreeToFindRightest tree.Node
	for next != nil && subtreeToFindRightest == nil {
		asPair, ok := next.(*tree.PairNode)
		if !ok {
			// todo err
			panic("panic")
		}

		if _, found := ancestors[asPair.LeftChild]; found {
			ancestors[next] = struct{}{}
			next = next.GetParent()
			continue
		}

		subtreeToFindRightest = asPair.LeftChild
	}

	if subtreeToFindRightest == nil {
		// we're at root and cannot find other ways to the left. That means we started with the leftest subtree.
		return
	}

	rightmost, err := findRightmostVal(subtreeToFindRightest, func(val *tree.ValueNode) bool {
		return true
	})

	if err != nil {
		// todo: handle
		panic("err")
	}

	rightmost.Value += value.Value
}

func findRightmostVal(root tree.Node, fn filterVal) (*tree.ValueNode, error) {
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

		// todo: only diff in here - switched these if statements places
		if asPair.RightChild != nil {
			s.push(stackEntry{
				node:  asPair.RightChild,
				depth: entry.depth + 1,
			})
		}
		if asPair.LeftChild != nil {
			s.push(stackEntry{
				node:  asPair.LeftChild,
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

func sum(root, n tree.Node) tree.Node {
	newRoot := &tree.PairNode{
		Parent:     nil,
		LeftChild:  root,
		RightChild: n,
	}

	root.SetParent(newRoot)
	n.SetParent(newRoot)

	return newRoot
}

func magnitude(root tree.Node) int {
	asVal, ok := root.(*tree.ValueNode)
	if ok {
		return asVal.Value
	}

	asPair, ok := root.(*tree.PairNode)
	if !ok {
		// todo: err
		panic("UNKNOWN TYPE")
	}

	res := 3*magnitude(asPair.LeftChild) + 2*magnitude(asPair.RightChild)

	return res
}

var errNotFound = errors.New("not found")
