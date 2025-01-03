// Package heap implements min heap. Taken from https://github.com/SuddenGunter/dsa-playground/blob/main/skiena/4/heap/heap.go.
package heap

import (
	"errors"
)

func minDominatesComparator[T any]() func(a, b Entity[T]) int8 {
	return func(a, b Entity[T]) int8 {
		switch {
		case a.Priority > b.Priority:
			return 1
		case a.Priority < b.Priority:
			return -1
		default:
			return 0
		}
	}
}

type Entity[T any] struct {
	Move     T
	Priority int
}

type Heap[T any] struct {
	body       []Entity[T]
	comparator func(a, b Entity[T]) int8
}

func NewHeap[T any]() *Heap[T] {
	return &Heap[T]{
		body:       []Entity[T]{},
		comparator: minDominatesComparator[T](),
	}
}

func FromSlice[T any](data []Entity[T]) *Heap[T] {
	h := &Heap[T]{
		body:       make([]Entity[T], 0, cap(data)),
		comparator: minDominatesComparator[T](),
	}

	for _, v := range data {
		h.Insert(v)
	}

	return h
}

func FromSliceFast[T any](data []Entity[T]) *Heap[T] {
	h := &Heap[T]{
		body:       make([]Entity[T], cap(data)),
		comparator: minDominatesComparator[T](),
	}

	for i, v := range data {
		h.body[i] = v
	}

	for i := len(h.body)/2 - 1; i >= 0; i-- {
		h.bubbleDown(i)
	}

	return h
}

func Heapsort[T any](data []Entity[T]) []Entity[T] {
	buffer := make([]Entity[T], 0, cap(data))
	h := FromSlice(data)
	for i := 0; i < cap(buffer); i++ {
		top, _ := h.TakeTop()
		buffer = append(buffer, top)
	}
	return buffer
}

func (h *Heap[T]) parentIndexOf(child int) int {
	child += 1

	if child == 1 {
		return -1
	}

	return (child / 2) - 1
}

func (h *Heap[T]) leftChildIndexOf(parent int) int {
	return 2*(parent+1) - 1
}

func (h *Heap[T]) Insert(x Entity[T]) {
	h.body = append(h.body, x)
	h.bubbleUp(len(h.body) - 1)
}

func (h *Heap[T]) bubbleUp(index int) {
	parentIndex := h.parentIndexOf(index)

	if parentIndex == -1 {
		return
	}

	if h.comparator(h.body[parentIndex], h.body[index]) == 1 {
		h.body[index], h.body[parentIndex] = h.body[parentIndex], h.body[index]
		h.bubbleUp(parentIndex)
	}
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.body) == 0
}

func (h *Heap[T]) TakeTop() (Entity[T], error) {
	if len(h.body) == 0 {
		return Entity[T]{}, errors.New("empty heap")
	}

	top := h.body[0]
	h.body[0] = h.body[len(h.body)-1]
	h.body = h.body[:len(h.body)-1]
	h.bubbleDown(0)

	return top, nil
}

func (h *Heap[T]) bubbleDown(index int) {
	leftChildIndex := h.leftChildIndexOf(index)

	topIndex := index

	// check that h.body[index] dominates both left and right children
	for i := 0; i <= 1; i++ {
		if leftChildIndex+i <= len(h.body)-1 {
			if h.comparator(h.body[topIndex], h.body[leftChildIndex+i]) == 1 {
				topIndex = leftChildIndex + i
			}
		}
	}

	// if h.body[index] doesn't dominate on child - swap and check the same for the lower level
	if topIndex != index {
		h.body[index], h.body[topIndex] = h.body[topIndex], h.body[index]
		h.bubbleDown(topIndex)
	}
}
