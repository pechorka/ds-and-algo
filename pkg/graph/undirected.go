package graph

import "ds-and-algo/pkg/set"

type Undirected struct {
	nodes map[string]*undirectedNode  // map of node name to node
	edges map[string]*set.Set[string] // map of node name to set of node names
}

func NewUndirected() *Undirected {
	return &Undirected{
		nodes: make(map[string]*undirectedNode),
		edges: make(map[string]*set.Set[string]),
	}
}

func (g *Undirected) AddNode(name string) {
	if _, ok := g.nodes[name]; ok {
		return
	}
	g.nodes[name] = newNode(name)
}

func (g *Undirected) AddEdge(name1, name2 string) {
	if _, ok := g.nodes[name1]; !ok {
		return
	}
	if _, ok := g.nodes[name2]; !ok {
		return
	}
	g.addEdge(name1, name2)
	g.addEdge(name2, name1)
}

func (g *Undirected) addEdge(name1, name2 string) {
	edgeSet, ok := g.edges[name1]
	if !ok {
		edgeSet = set.New[string]()
		g.edges[name1] = edgeSet
	}
	edgeSet.Add(name2)
}

func (g *Undirected) Nodes() []string {
	nodes := make([]string, 0, len(g.nodes))
	for name := range g.nodes {
		nodes = append(nodes, name)
	}
	return nodes
}

func (g *Undirected) Edges() [][2]string {
	edges := make([][2]string, 0, len(g.edges))
	for name1, edgeSet := range g.edges {
		edgeSet.ForEach(func(name2 string) {
			edges = append(edges, [2]string{name1, name2})
		})
	}
	return edges
}

func (g *Undirected) DFS(start string) []string {
	visited := make([]string, 0, len(g.nodes))
	visitedSet := set.New[string]()
	g.dfs(start, &visited, visitedSet)
	return visited
}

func (g *Undirected) dfs(name string, visited *[]string, visitedSet *set.Set[string]) {
	if visitedSet.Contains(name) {
		return
	}
	visitedSet.Add(name)
	*visited = append(*visited, name)
	neighbors, ok := g.edges[name]
	if !ok {
		return
	}
	neighbors.ForEach(func(neighbor string) {
		g.dfs(neighbor, visited, visitedSet)
	})
}

func (g *Undirected) BFS(start string) []string {
	visited := make([]string, 0, len(g.nodes))
	visitedSet := set.New[string]()
	queue := []string{start}
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		if visitedSet.Contains(name) {
			continue
		}
		visitedSet.Add(name)
		visited = append(visited, name)
		neighbors, ok := g.edges[name]
		if !ok {
			continue
		}
		neighbors.ForEach(func(neighbor string) {
			queue = append(queue, neighbor)
		})
	}
	return visited
}

type undirectedNode struct {
	name string
}

func newNode(name string) *undirectedNode {
	return &undirectedNode{
		name: name,
	}
}
