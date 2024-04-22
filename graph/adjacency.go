package graph

import (
	"fmt"
	"io"
)

// Adjacency represents an undirected graph stored using an adjacency
// list.
type Adjacency struct {
	data map[int][]int
}

func NewAdjacency(nodes ...int) *Adjacency {
	a := Adjacency{
		data: make(map[int][]int),
	}

	for i := range nodes {
		a.AddNode(nodes[i])
	}

	return &a
}

func (a *Adjacency) AddNode(node int) {
	if _, exists := a.data[node]; !exists {
		a.data[node] = nil
	}
}

func (a *Adjacency) AddEdge(from, to int) {
	a.data[from] = append(a.data[from], to)
	a.data[to] = append(a.data[to], from)
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

func (a *Adjacency) DeleteNode(node int) {
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

	visited := make([]bool, len(a.data))
	queue, order := []int{start}, []int{start}
	visited[start] = true

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		for _, edge := range a.data[current] {
			if !visited[edge] {
				visited[edge] = true
				order = append(order, edge)
				queue = append(queue, edge)
			}
		}
	}

	return order
}

func (a *Adjacency) DFS(start int) []int {
	if _, exists := a.data[start]; !exists {
		return nil
	}

	stack := []int{start}
	var order []int
	visited := make([]bool, len(a.data))

	for len(stack) != 0 {
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[currentNode] {
			visited[currentNode] = true
			order = append(order, currentNode)
			stack = append(stack, a.data[currentNode]...)
		}
	}

	return order
}
