package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGraphOperations(t *testing.T) {
	g := NewUndirected()

	g.AddNode("A")
	g.AddNode("B")
	g.AddEdge("A", "B")

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

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
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

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	// now graph looks like
	// A
	// ├── B
	// │   └── D
	// └── C

	visited := g.BFS("A")
	expectedOrder := []string{"A", "B", "C", "D"}
	require.Equal(t, expectedOrder, visited)
}
