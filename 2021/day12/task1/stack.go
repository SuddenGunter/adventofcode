package task1

type stack struct {
	data []pathEntry
}

func (s *stack) pop() pathEntry {
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el
}

func (s *stack) push(p pathEntry) {
	s.data = append(s.data, p)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) peek() pathEntry {
	return s.data[len(s.data)-1]
}
