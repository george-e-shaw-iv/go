package stack

import (
	"sync"

	"github.com/george-e-shaw-iv/go/queue"
)

type Stack[T any] interface {
	Push(val T)
	Top() T
	Pop()
	Len() int
}

type Classic[T any] struct {
	data []T
	mu   *sync.Mutex
}

func NewClassic[T any]() *Classic[T] {
	return &Classic[T]{
		mu: &sync.Mutex{},
	}
}

func (s *Classic[T]) Push(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, val)
}

func (s *Classic[T]) Top() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data[len(s.data)-1]
}

func (s *Classic[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.data)
}

func (s *Classic[T]) Pop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = s.data[:len(s.data)-1]
}

type QueueBased[T any] struct {
	main, staging *queue.Queue[T]
	mu            sync.Mutex
}

func NewQueueBased[T any]() *QueueBased[T] {
	return &QueueBased[T]{
		main:    queue.NewQueue[T](),
		staging: queue.NewQueue[T](),
	}
}

func (s *QueueBased[T]) Push(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.staging.Enqueue(val)

	for s.main.Len() > 0 {
		s.staging.Enqueue(s.main.Dequeue())
	}

	tmp := s.main
	s.main = s.staging
	s.staging = tmp
}

func (s *QueueBased[T]) Top() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.main.Peek()
}

func (s *QueueBased[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.main.Len()
}

func (s *QueueBased[T]) Pop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.main.Dequeue()
}
