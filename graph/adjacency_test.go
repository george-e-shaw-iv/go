package graph_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/george-e-shaw-iv/go/graph"
)

func TestAdjancency(t *testing.T) {
	g := graph.NewAdjacency(0, 1, 2, 3, 4, 5)

	g.Debug(os.Stdout)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(0, 4)

	g.AddEdge(3, 5)
	g.AddEdge(1, 5)

	fmt.Println("-----")

	g.Debug(os.Stdout)

	fmt.Println("-----")

	fmt.Println(g.BFS(0))
	fmt.Println(g.DFS(0))

	t.Fail()
}
