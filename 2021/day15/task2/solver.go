package task2

import (
	"aoc-2021-day15/graph"
	"aoc-2021-day15/heap"
	"math"
)

func Solve(data graph.Data) (int, error) {
	known := make(map[graph.Position]struct{})
	h := heap.NewHeap()

	distance := make(map[graph.Position]int)
	for k := range data.Graph.Nodes {
		distance[k] = math.MaxInt
	}

	distance[data.Start.ID] = 0
	known[data.Start.ID] = struct{}{}
	h.Insert(heap.Entity{Node: data.Start, Priority: 0})

	// todo: could optimize and left when finish is found
	for !h.IsEmpty() {
		top, err := h.TakeTop()
		if err != nil {
			return 0, err
		}

		last := top.Node

		known[last.ID] = struct{}{}

		for _, w := range last.Adjacent {
			_, found := known[w.ID]
			if found {
				continue
			}

			if distance[w.ID] > distance[last.ID]+w.Weight {
				distance[w.ID] = distance[last.ID] + w.Weight
				h.Insert(heap.Entity{
					Node:     w,
					Priority: distance[last.ID] + w.Weight,
				})
			}
		}
	}

	return distance[data.Finish.ID], nil
}
