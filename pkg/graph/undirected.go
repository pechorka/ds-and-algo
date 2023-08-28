package graph

import "ds-and-algo/pkg/set"

type Undirected struct {
	nodes map[string]string                   // map of node name to node
	edges map[string]*set.Set[undirectedEdge] // map of node name to set of node names
}

func NewUndirected() *Undirected {
	return &Undirected{
		nodes: make(map[string]string),
		edges: make(map[string]*set.Set[undirectedEdge]),
	}
}

func (g *Undirected) AddNode(name string) {
	if _, ok := g.nodes[name]; ok {
		return
	}
	g.nodes[name] = name
}

func (g *Undirected) AddEdge(name1, name2 string, distance int) {
	if _, ok := g.nodes[name1]; !ok {
		return
	}
	if _, ok := g.nodes[name2]; !ok {
		return
	}
	g.addEdge(name1, name2, distance)
	g.addEdge(name2, name1, distance)
}

func (g *Undirected) addEdge(name1, name2 string, distance int) {
	edgeSet, ok := g.edges[name1]
	if !ok {
		edgeSet = set.New[undirectedEdge]()
		g.edges[name1] = edgeSet
	}
	edgeSet.Add(newUndirectedEdge(name2, distance))
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
	for start, edgeSet := range g.edges {
		edgeSet.ForEach(func(e undirectedEdge) {
			edges = append(edges, [2]string{start, e.destination})
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
	neighbors.ForEach(func(e undirectedEdge) {
		g.dfs(e.destination, visited, visitedSet)
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
		neighbors.ForEach(func(e undirectedEdge) {
			queue = append(queue, e.destination)
		})
	}
	return visited
}

func (g *Undirected) Dijkstra(start string) map[string]int { // map of node name to distance
	visitedSet := set.New[string]()
	distances := make(map[string]int)
	const infinity = int(^uint(0) >> 1) // max int
	for name := range g.nodes {
		distances[name] = infinity
	}
	current := start
	distances[current] = 0
	for {
		if current == "" {
			return distances
		}
		if visitedSet.Contains(current) {
			current = selectSmallest(distances, visitedSet)
			continue
		}
		visitedSet.Add(current)
		neighbors, ok := g.edges[current]
		if !ok {
			current = selectSmallest(distances, visitedSet)
			continue
		}
		currentDistance := distances[current]
		neighbors.ForEach(func(e undirectedEdge) {
			totalDistance := e.distance + currentDistance
			if totalDistance < distances[e.destination] {
				distances[e.destination] = totalDistance
			}
		})
		current = selectSmallest(distances, visitedSet)
	}
}

func selectSmallest(distances map[string]int, visitedSet *set.Set[string]) string {
	var smallest string
	smallestDistance := int(^uint(0) >> 1) // max int
	for name, distance := range distances {
		if visitedSet.Contains(name) {
			continue
		}
		if distance < smallestDistance {
			smallestDistance = distance
			smallest = name
		}
	}
	return smallest
}

type undirectedEdge struct {
	destination string
	distance    int
}

func newUndirectedEdge(distanation string, distances int) undirectedEdge {
	return undirectedEdge{destination: distanation, distance: distances}
}
