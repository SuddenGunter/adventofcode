package task2

type queue struct {
	elements []position
}

func (q *queue) enqueue(p position) {
	q.elements = append(q.elements, p)
}

func (q *queue) dequeue() position {
	returned := q.elements[0]

	q.elements = q.elements[1:]

	return returned
}

func (q *queue) isEmpty() bool {
	return len(q.elements) == 0
}
