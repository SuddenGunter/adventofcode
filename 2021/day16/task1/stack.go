package task1

import "aoc-2021-day16/packet"

type stack struct {
	data []packet.Packet
}

func (s *stack) pop() packet.Packet {
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el
}

func (s *stack) push(p packet.Packet) {
	s.data = append(s.data, p)
}

func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) peek() packet.Packet {
	return s.data[len(s.data)-1]
}
