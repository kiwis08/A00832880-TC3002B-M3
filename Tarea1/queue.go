package main

type ArrayQueue struct {
	elements []interface{}
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{
		elements: make([]interface{}, 0),
	}
}

func (q *ArrayQueue) Enqueue(value interface{}) {
	q.elements = append(q.elements, value)
}

func (q *ArrayQueue) Dequeue() interface{} {
	if len(q.elements) == 0 {
		return nil
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

func (q *ArrayQueue) Front() interface{} {
	if len(q.elements) == 0 {
		return nil
	}
	return q.elements[0]
}

func (q *ArrayQueue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *ArrayQueue) Size() int {
	return len(q.elements)
}

func (q *ArrayQueue) Clear() {
	q.elements = make([]interface{}, 0)
}
