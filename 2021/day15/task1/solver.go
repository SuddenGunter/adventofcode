package task1

import (
	"aoc-2021-day15/graph"
	"math"
)

func Solve(data graph.Data) (int, error) {
	known := make(map[graph.Position]struct{})
	known[data.Start.ID] = struct{}{}

	distance := make(map[graph.Position]int)
	for k := range data.Graph.Nodes {
		distance[k] = math.MaxInt
	}

	distance[data.Start.ID] = 0

	last := data.Start

	for !last.Equals(data.Finish) {
		known[last.ID] = struct{}{}

		for _, w := range last.Adjacent {
			if distance[w.ID] > distance[last.ID]+w.Weight {
				distance[w.ID] = distance[last.ID] + w.Weight
			}
		}

		dist := math.MaxInt
		for k, v := range distance {
			_, found := known[k]
			if !found && dist > v {
				dist = v
				last = data.Graph.Nodes[k]
			}
		}

	}

	return data.Start.Weight + distance[data.Finish.ID], nil
}
