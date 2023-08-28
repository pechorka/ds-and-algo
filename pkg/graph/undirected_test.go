package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGraphOperations(t *testing.T) {
	g := NewUndirected()

	g.AddNode("A")
	g.AddNode("B")
	g.AddEdge("A", "B", 1)

	require.ElementsMatch(t, []string{"A", "B"}, g.Nodes())
	expectedEdges := [][2]string{{"A", "B"}, {"B", "A"}}
	for _, e := range g.Edges() {
		require.Contains(t, expectedEdges, e)
	}
}

func TestDFS(t *testing.T) {
	g := NewUndirected()

	nodes := []string{"A", "B", "C", "D"}
	for _, n := range nodes {
		g.AddNode(n)
	}

	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 1)
	g.AddEdge("B", "D", 1)
	// now graph looks like
	// A
	// ├── B
	// │   └── D
	// └── C

	visited := g.DFS("A")
	expectedOrder := []string{"A", "B", "D", "C"}
	require.Equal(t, expectedOrder, visited)
}

func TestBFS(t *testing.T) {
	g := NewUndirected()

	nodes := []string{"A", "B", "C", "D"}
	for _, n := range nodes {
		g.AddNode(n)
	}

	g.AddEdge("A", "B", 1)
	g.AddEdge("A", "C", 1)
	g.AddEdge("B", "D", 1)
	// now graph looks like
	// A
	// ├── B
	// │   └── D
	// └── C

	visited := g.BFS("A")
	// either order is possible
	expectedOrder1 := []string{"A", "B", "C", "D"}
	expectedOrder2 := []string{"A", "C", "B", "D"}
	require.Subset(t, [][]string{expectedOrder1, expectedOrder2}, [][]string{visited})
}

func TestDijkstrasAlgorithm(t *testing.T) {
	g := NewUndirected()

	g.AddNode("A")
	g.AddNode("B")
	g.AddNode("C")
	g.AddEdge("A", "B", 1)
	g.AddEdge("B", "C", 2)
	g.AddEdge("A", "C", 4)
	// now graph looks like
	// A - 1 - B
	//   \     |
	//     4   2
	//       \ |
	//         C

	distances := g.Dijkstra("A")
	require.Equal(t, map[string]int{
		"A": 0,
		"B": 1,
		"C": 3, // A -> B -> C is the shortest path
	}, distances)
}
