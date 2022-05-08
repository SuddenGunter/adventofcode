package task1

import (
	"math"

	"aoc-2015-day9/graph"
)

type path struct {
	visited []int
	len     int
}

func Solve(g *graph.Graph) int {
	nodePaths := make(map[int][]path)
	for node := 0; node < g.Len(); node++ {
		p := getPath(node, path{visited: []int{node}}, g)
		nodePaths[node] = p
	}

	minPath := math.MaxInt
	for _, node := range nodePaths {
		for _, p := range node {
			if p.len < minPath {
				minPath = p.len
			}
		}
	}

	return minPath
}

func getPath(node int, p path, g *graph.Graph) []path {
	if len(p.visited) == g.Len() {
		return []path{p}
	}

	paths := make([]path, 0)

	for destination := 0; destination < g.Len(); destination++ {
		if contains(p.visited, destination) {
			continue
		}

		if node == destination {
			continue
		}

		dist, err := g.Distance(node, destination)
		if err != nil {
			continue
		}

		visited := make([]int, len(p.visited)+1)
		copy(visited, p.visited)

		visited[len(visited)-1] = destination
		plen := p.len + dist

		paths = append(paths, getPath(destination, path{
			visited: visited,
			len:     plen,
		}, g)...)
	}

	return paths
}

func contains(visited []int, destination int) bool {
	for _, v := range visited {
		if v == destination {
			return true
		}
	}

	return false
}
