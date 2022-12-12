package task1

import (
	"aoc-2022-day12/graph"
	"math"
)

type Solver struct {
	edgeTo  []*graph.Edge
	distTo  []int
	visited map[int]struct{}
}

func (s *Solver) Solve(g *graph.Graph) int {
	// edge that is the latest part of a shortest path to vertex.
	s.edgeTo = make([]*graph.Edge, g.VerticesCount())
	// shortest distance to a vertex.
	s.distTo = make([]int, g.VerticesCount())
	s.visited = make(map[int]struct{})

	for i := 0; i < g.VerticesCount(); i++ {
		s.distTo[i] = math.MaxInt
	}
	v := g.Start()
	s.distTo[v] = 0

	for {
		if _, ok := s.visited[v]; ok {
			break
		}

		v = s.relax(g, v)
	}

	return s.distTo[g.End()]
}

func (s *Solver) relax(g *graph.Graph, v int) int {
	s.visited[v] = struct{}{}

	for _, edge := range g.Adjacency[v] {
		dst := edge.To
		if s.distTo[dst] > s.distTo[v]+edge.Weight {
			s.distTo[dst] = s.distTo[v] + edge.Weight
			s.edgeTo[dst] = &edge
		}
	}

	// find next closest not processed vertex
	closest := math.MaxInt
	for i := 0; i < g.VerticesCount(); i++ {
		_, processed := s.visited[i]
		if !processed && (closest > s.distTo[i]) {
			closest = s.distTo[i]
			v = i
		}
	}

	return v
}
