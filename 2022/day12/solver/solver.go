package solver

import (
	"aoc-2022-day12/graph"
	"math"
	"sync"
)

func SolveTask(g *graph.Graph, num int) int {
	switch num {
	case 1:
		s := &dijkstraSolver{}
		start := g.Start()[0]
		return s.solve(g, start)
	case 2:
		shortestPaths := make([]int, 0)
		results := make(chan int)
		go func() {
			for v := range results {
				shortestPaths = append(shortestPaths, v)
			}
		}()
		wg := sync.WaitGroup{}
		for _, start := range g.Start() {
			wg.Add(1)
			go func(startV int) {
				s := &dijkstraSolver{}
				results <- s.solve(g, startV)
				wg.Done()
			}(start)
		}
		wg.Wait()
		close(results)

		return min(shortestPaths)
	default:
		panic("unexpected task")
	}
}

func min(paths []int) int {
	minP := math.MaxInt
	for _, v := range paths {
		if v < minP {
			minP = v
		}
	}

	return minP
}

type dijkstraSolver struct {
	edgeTo  []*graph.Edge
	distTo  []int
	visited map[int]struct{}
}

func (s *dijkstraSolver) solve(g *graph.Graph, start int) int {
	// edge that is the latest part of a shortest path to vertex.
	s.edgeTo = make([]*graph.Edge, g.VerticesCount())
	// shortest distance to a vertex.
	s.distTo = make([]int, g.VerticesCount())
	s.visited = make(map[int]struct{})

	for i := 0; i < g.VerticesCount(); i++ {
		s.distTo[i] = math.MaxInt
	}
	v := start
	s.distTo[v] = 0

	for {
		if _, ok := s.visited[v]; ok {
			break
		}

		v = s.relax(g, v)
	}

	return s.distTo[g.End()]
}

func (s *dijkstraSolver) relax(g *graph.Graph, v int) int {
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
