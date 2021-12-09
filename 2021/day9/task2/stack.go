package task2

type stack struct {
	data []position
}

func (s *stack) pop() position {
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el
}

func (s *stack) push(p position) {
	s.data = append(s.data, p)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}
