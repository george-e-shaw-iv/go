package graph

import (
	"fmt"
	"io"
)

var _ Graph = &Adjacency{}

// Adjacency represents an undirected graph stored using an adjacency
// list.
type Adjacency struct {
	data map[int][]int
	opts *opts
}

func NewAdjacency(options ...GraphOption) *Adjacency {
	a := Adjacency{
		data: make(map[int][]int),
		opts: &opts{},
	}
	a.opts.apply(options...)

	return &a
}

func (a *Adjacency) AddNode(node int) {
	if _, exists := a.data[node]; !exists {
		a.data[node] = nil
	}
}

func (a *Adjacency) AddEdge(from, to int) {
	a.data[from] = append(a.data[from], to)

	if a.opts.undirected {
		a.data[to] = append(a.data[to], from)
	} else {
		a.AddNode(to) // only adds if it doesn't exist already
	}
}

func (a *Adjacency) RemoveEdge(from, to int) {
	if _, exists := a.data[from]; !exists {
		return
	}

	if _, exists := a.data[to]; !exists {
		return
	}

	for i, adjacent := range a.data[from] {
		if adjacent == to {
			a.data[from][i] = a.data[from][len(a.data[from])-1]
			a.data[from] = a.data[from][:len(a.data[from])-1]
			break
		}
	}

	for i, adjacent := range a.data[to] {
		if adjacent == from {
			a.data[to][i] = a.data[to][len(a.data[to])-1]
			a.data[to] = a.data[to][:len(a.data[to])-1]
			break
		}
	}
}

func (a *Adjacency) RemoveNode(node int) {
	if _, exists := a.data[node]; !exists {
		return
	}

	for _, adjacent := range a.data[node] {
		a.RemoveEdge(node, adjacent)
	}
	delete(a.data, node)
}

func (a *Adjacency) Debug(writer io.Writer) {
	for node, edges := range a.data {
		fmt.Fprintf(writer, "%v: %+v\n", node, edges)
	}
}

func (a *Adjacency) BFS(start int) []int {
	if _, exists := a.data[start]; !exists {
		return nil
	}

	visited := make(map[int]struct{}, len(a.data))
	queue, order := []int{start}, []int{start}
	visited[start] = struct{}{}

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		for _, edge := range a.data[current] {
			if _, exists := visited[edge]; !exists {
				visited[edge] = struct{}{}
				order = append(order, edge)
				queue = append(queue, edge)
			}
		}
	}

	return order
}

func (a *Adjacency) DFSIterative(start int) []int {
	if _, exists := a.data[start]; !exists {
		return nil
	}

	stack := []int{start}
	order := make([]int, 0, len(a.data))
	visited := make(map[int]struct{}, len(a.data))

	for len(stack) != 0 {
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if _, exists := visited[currentNode]; !exists {
			visited[currentNode] = struct{}{}
			order = append(order, currentNode)
			stack = append(stack, a.data[currentNode]...)
		}
	}

	return order
}

func (a *Adjacency) DFSRecursive(start int) []int {
	if _, exists := a.data[start]; !exists {
		return nil
	}

	order := make([]int, 0, len(a.data))
	visited := make(map[int]struct{}, len(a.data))

	var dfs func(node int)
	dfs = func(node int) {
		if _, exists := visited[node]; exists {
			return
		}

		order = append(order, node)
		visited[node] = struct{}{}

		for i := range a.data[node] {
			dfs(a.data[node][i])
		}
	}

	dfs(start)
	return order
}

func (a *Adjacency) TopologicalSort() ([]int, error) {
	order, orderIdx := make([]int, len(a.data)), len(a.data)-1
	visited := make(map[int]struct{}, len(a.data))
	visiting := make(map[int]struct{}, len(a.data))

	var err error
	var dfs func(node int)
	dfs = func(node int) {
		if _, exists := visited[node]; exists {
			return
		}

		if _, exists := visiting[node]; exists {
			err = ErrGraphHasCycle
			return
		}

		visiting[node] = struct{}{}
		for i := range a.data[node] {
			dfs(a.data[node][i])
		}
		delete(visiting, node)

		visited[node] = struct{}{}
		order[orderIdx] = node
		orderIdx--
	}

	for node := range a.data {
		dfs(node)
	}
	return order, err
}
