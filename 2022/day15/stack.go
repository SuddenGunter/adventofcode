package main

type stack[T any] struct {
	data []T
}

func newStack[T any]() *stack[T] {
	return &stack[T]{
		data: make([]T, 0),
	}
}

func (s *stack[T]) pop() T {
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return el
}

func (s *stack[T]) push(p T) {
	s.data = append(s.data, p)
}

func (s *stack[T]) isEmpty() bool {
	return len(s.data) == 0
}

func (s *stack[T]) peek() T {
	return s.data[len(s.data)-1]
}
