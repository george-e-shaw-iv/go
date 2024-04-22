package graph

import "golang.org/x/exp/constraints"

type Graph[T constraints.Ordered] interface {
	AddNode(node T)
	AddEdge(from, to T)
}
