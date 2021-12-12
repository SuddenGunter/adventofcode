package task1

import "aoc-2021-day12/input"

type pathEntry struct {
	node string
	// copy on each step
	followedPath []string
	// copy on each step

	visitedSmall map[string]struct{}
}

func Solve(data input.Data) (int, error) {
	return 0, nil
}
