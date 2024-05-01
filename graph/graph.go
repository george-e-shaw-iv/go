package graph

import (
	"errors"
	"io"
)

// ErrGraphHasCycle is an error that can be returned from methods that cannot
// return a result for graphs with cycles in them, like topological sort.
var ErrGraphHasCycle = errors.New("graph has cycle")

type opts struct {
	undirected bool
}

func (o *opts) apply(graphOptions ...GraphOption) {
	for i := range graphOptions {
		graphOptions[i](o)
	}
}

type GraphOption func(*opts)

// AsUndirected makes the graph an undirected graph.
func AsUndirected() GraphOption {
	return func(o *opts) {
		o.undirected = true
	}
}

type Graph interface {
	AddNode(node int)
	AddEdge(from, to int)
	RemoveNode(node int)
	RemoveEdge(from int, to int)
	Debug(writer io.Writer)
	BFS(start int) []int
	DFSIterative(start int) []int
	DFSRecursive(start int) []int
	TopologicalSort() ([]int, error)
}
