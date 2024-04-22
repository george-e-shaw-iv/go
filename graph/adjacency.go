package graph

import "golang.org/x/exp/constraints"

type Adjacency[T constraints.Ordered] struct {
	data map[T][]T
}

func NewAdjacency[T constraints.Ordered]() *Adjacency[T] {
	return &Adjacency[T]{
		data: make(map[T][]T),
	}
}

func (a *Adjacency[T]) AddNode(node T) {
	if _, exists := a.data[node]; !exists {
		a.data[node] = nil
	}
}

func (a *Adjacency[T]) AddEdge(from, to T) {
	a.AddNode(from)
	a.AddNode(to)

	a.data[from] = append(a.data[from], to)
	a.data[from] = append(a.data[to], from)
}

func (a *Adjacency[T]) DeleteNode(node T) {
	if _, exists := a.data[node]; !exists {
		return
	}
}
