package task2

import (
	"aoc-2021-day12/cave"
	"aoc-2021-day12/input"
	"errors"
)

const maxSaneLimit = 500000

type pathEntry struct {
	node cave.Cave
	// copy on each step
	followedPath []string
	// copy on each step

	visitedSmall      map[string]byte
	smallVisitedTwice bool
}

func Solve(data input.Data) (int, error) {
	toTraverse := &stack{
		data: make([]pathEntry, 0, 1),
	}

	entry := pathEntry{
		node:         data.Caves[cave.Start],
		followedPath: []string{cave.Start},
		visitedSmall: map[string]byte{
			cave.Start: 1,
		},
		smallVisitedTwice: false,
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
			// fmt.Println("finished path", entry.followedPath, entry.visitedSmall)
			continue
		}

		for _, c := range entry.node.ConnectedTo {
			if c == cave.Start {
				continue
			}

			newCave := data.Caves[c]

			if newCave.Size != cave.Small {
				newCavePath := pathEntry{
					node:              newCave,
					followedPath:      append(copyS(entry.followedPath), c),
					visitedSmall:      copyM(entry.visitedSmall),
					smallVisitedTwice: entry.smallVisitedTwice,
				}

				toTraverse.push(newCavePath)
				continue
			}

			if entry.smallVisitedTwice == false {
				count := entry.visitedSmall[c]
				count += 1

				visitedSmall := copyM(entry.visitedSmall)
				visitedSmall[c] = count

				if count > 1 {
					newCavePath := pathEntry{
						node:              newCave,
						followedPath:      append(copyS(entry.followedPath), c),
						visitedSmall:      visitedSmall,
						smallVisitedTwice: true,
					}
					toTraverse.push(newCavePath)
				} else {
					newCavePath := pathEntry{
						node:              newCave,
						followedPath:      append(copyS(entry.followedPath), c),
						visitedSmall:      visitedSmall,
						smallVisitedTwice: false,
					}

					toTraverse.push(newCavePath)
				}

			} else {
				count := entry.visitedSmall[c]
				if count > 0 {
					continue
				}

				visitedSmall := copyM(entry.visitedSmall)
				visitedSmall[c] = count + 1

				newCavePath := pathEntry{
					node:              newCave,
					followedPath:      append(copyS(entry.followedPath), c),
					visitedSmall:      visitedSmall,
					smallVisitedTwice: entry.smallVisitedTwice,
				}

				toTraverse.push(newCavePath)
			}
		}
	}

	return pathCount, nil
}

func copyM(m map[string]byte) map[string]byte {
	newMap := make(map[string]byte, len(m))
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
