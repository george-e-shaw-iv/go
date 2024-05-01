package graph_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/george-e-shaw-iv/go/graph"
)

func TestAdjancency(t *testing.T) {
	// vector<vector<int> > edges
	//     = { { 0, 1 }, { 1, 2 }, { 3, 1 }, { 3, 2 } };

	// // Graph represented as an adjacency list
	// vector<vector<int> > adj(V);

	// for (auto i : edges) {
	//     adj[i[0]].push_back(i[1]);
	// }
	g2 := graph.NewAdjacency()
	g2.AddEdge(0, 1)
	g2.AddEdge(1, 2)
	g2.AddEdge(3, 1)
	g2.AddEdge(3, 2)
	g2.Debug(os.Stdout)
	fmt.Println(g2.TopologicalSort())

	t.Fail()
}
