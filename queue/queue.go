package queue

import "sync"

type Queue[T any] struct {
	data []T
	mu   *sync.Mutex
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		mu: &sync.Mutex{},
	}
}

func (q *Queue[T]) Enqueue(val T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = append(q.data, val)
}

func (q *Queue[T]) Dequeue() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queue[T]) Peek() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.data[0]
}

func (q *Queue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.data)
}
