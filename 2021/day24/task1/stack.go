package task1

type Entry struct {
	Value int
	Digit int
}

type stack struct {
	data []Entry
}

func (s *stack) pop() Entry {
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el
}

func (s *stack) push(p Entry) {
	s.data = append(s.data, p)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) peek() Entry {
	return s.data[len(s.data)-1]
}
