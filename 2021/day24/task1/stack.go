package task1

type stack struct {
	data []int64
}

func (s *stack) pop() int64 {
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el
}

func (s *stack) push(p int64) {
	s.data = append(s.data, p)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) peek() int64 {
	return s.data[len(s.data)-1]
}
