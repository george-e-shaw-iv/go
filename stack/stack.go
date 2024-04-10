package stack

import "sync"

// Stack implements a thread-safe stack data structure. It is important
// to note that this implementation will not stop you from retrieving
// or popping elements from an empty stack. It is on the user to ensure
// that there are items on the stack before calling those methods.
type Stack[T any] struct {
	data []T
	mu   sync.Mutex
}

func (s *Stack[T]) Push(val T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, val)
}

func (s *Stack[T]) Top() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.data)
}

func (s *Stack[T]) Pop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = s.data[:len(s.data)-1]
}
