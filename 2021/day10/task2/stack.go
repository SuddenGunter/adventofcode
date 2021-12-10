package task2

type stack struct {
	data []rune
}

func (s *stack) pop() rune {
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el
}

func (s *stack) push(p rune) {
	s.data = append(s.data, p)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) peek() rune {
	return s.data[len(s.data)-1]
}
