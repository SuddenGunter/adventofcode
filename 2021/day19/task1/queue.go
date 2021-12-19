package task1

import "aoc-2021-day19/point"

type queue struct {
	elements []map[point.Point3d]struct{}
}

func (q *queue) enqueue(p map[point.Point3d]struct{}) {
	q.elements = append(q.elements, p)
}

func (q *queue) dequeue() map[point.Point3d]struct{} {
	returned := q.elements[0]

	q.elements = q.elements[1:]

	return returned
}

func (q *queue) isEmpty() bool {
	return len(q.elements) == 0
}
