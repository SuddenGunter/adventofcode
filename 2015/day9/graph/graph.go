package graph

import "errors"

const maxSize = 25

type Graph struct {
	nodes     map[string]int
	names     map[int]string
	adjacency [maxSize][maxSize]int
}

func New() *Graph {
	return &Graph{nodes: make(map[string]int), names: make(map[int]string)}
}

func (g *Graph) AddEdge(from, to string, weight int) {
	fromIdx, ok := g.nodes[from]
	if !ok {
		fromIdx = len(g.nodes)
		g.nodes[from] = fromIdx
		g.names[fromIdx] = from
	}

	toIdx, ok := g.nodes[to]
	if !ok {
		toIdx = len(g.nodes)
		g.nodes[to] = toIdx
		g.names[toIdx] = to
	}

	g.adjacency[fromIdx][toIdx] = weight
}

func (g *Graph) Len() int {
	return len(g.nodes)
}

func (g *Graph) Distance(from, to int) (int, error) {
	if from == to {
		return 0, nil
	}

	val := g.adjacency[from][to]
	if val == 0 {
		return 0, ErrNoPath
	}

	return val, nil
}

var ErrNoPath = errors.New("no path between nodes")
