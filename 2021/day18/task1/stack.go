package task1

import "aoc-2021-day18/tree"

type stack struct {
	data []tree.Node
}

func (s *stack) pop() tree.Node {
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el
}

func (s *stack) push(p tree.Node) {
	s.data = append(s.data, p)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) peek() tree.Node {
	return s.data[len(s.data)-1]
}
