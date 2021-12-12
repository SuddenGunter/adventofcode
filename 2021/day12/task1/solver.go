package task1

import (
	"aoc-2021-day12/cave"
	"aoc-2021-day12/input"
	"errors"
)

const maxSaneLimit = 1000

type pathEntry struct {
	node cave.Cave
	// copy on each step
	followedPath []string
	// copy on each step

	visitedSmall map[string]struct{}
}

func Solve(data input.Data) (int, error) {
	toTraverse := &stack{
		data: make([]pathEntry, 0, 1),
	}

	entry := pathEntry{
		node:         data.Caves[cave.Start],
		followedPath: []string{cave.Start},
		visitedSmall: make(map[string]struct{}),
	}

	toTraverse.push(entry)

	pathCount := 0

	for !toTraverse.isEmpty() {
		if pathCount >= maxSaneLimit {
			return -1, errors.New("pathCount is too big, probably a bug")
		}

		entry := toTraverse.pop()

		if entry.followedPath[len(entry.followedPath)-1] == cave.End {
			pathCount++
			continue
		}

		for _, c := range entry.node.ConnectedTo {
			newCave := data.Caves[c]

			if newCave.Size != cave.Small {
				newCavePath := pathEntry{
					node:         newCave,
					followedPath: append(copyS(entry.followedPath), c),
					visitedSmall: copyM(entry.visitedSmall),
				}

				toTraverse.push(newCavePath)
				continue
			}

			if _, found := entry.visitedSmall[c]; !found {
				visitedSmall := copyM(entry.visitedSmall)
				visitedSmall[c] = struct{}{}

				newCavePath := pathEntry{
					node:         newCave,
					followedPath: append(copyS(entry.followedPath), c),
					visitedSmall: visitedSmall,
				}

				toTraverse.push(newCavePath)
			}
		}
	}

	return pathCount, nil
}

func copyM(m map[string]struct{}) map[string]struct{} {
	newMap := make(map[string]struct{}, len(m))
	for k, v := range m {
		newMap[k] = v
	}

	return newMap
}

func copyS(s []string) []string {
	newSlice := make([]string, 0, len(s))
	for _, v := range s {
		newSlice = append(newSlice, v)
	}

	return newSlice
}
