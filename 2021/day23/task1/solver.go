package task1

import (
	"aoc-2021-day23/task1/amphipod"
	"math"
)

type Move struct {
	TotalCost      uint64
	StateAfterMove amphipod.Burrow
}

func Solve(data amphipod.Burrow) uint64 {
	if done(data.Rooms) {
		return 0
	}

	best := uint64(math.MaxUint64)

	for _, move := range getValidMoves(data) {
		cost := move.TotalCost
		cost += Solve(data)

		if cost < best {
			best = cost
		}
	}

	return best
}

func done(rooms [amphipod.Rooms][amphipod.RoomSize]rune) bool {
	char := 'A'
	for i := range rooms {
		for _, r := range rooms[i] {
			if r != char {
				return false
			}
		}

		char++
	}

	return true
}

func getValidMoves(data amphipod.Burrow) []Move {
	movesFromHall := getValidMovesFromHall(data)
	movesFromRoom := getValidMovesFromRoom(data)

	return append(movesFromHall, movesFromRoom...)
}

func getValidMovesFromRoom(data amphipod.Burrow) []Move {
	movesFromRoom := make([]Move, 0)

	for i, room := range data.Rooms {
		j := 0
		for j < amphipod.RoomSize {
			if room[j] == '.' {
				j++
				continue
			}

			break
		}

		// room is empty
		// todo: check if correct
		if j == amphipod.RoomSize {
			continue
		}

		for h := range data.Hall {
			if !amphipod.EnterableHallCell(h) {
				continue
			}

			cost := getMoveCost(h, i, j, data, room[j])

			// no path or no place in cell
			if cost == math.MaxUint64 {
				continue
			}

			clone := data.Clone()
			clone.Hall[h] = clone.Rooms[i][j]
			clone.Rooms[i][j] = '.'

			movesFromRoom = append(movesFromRoom, Move{
				TotalCost:      cost,
				StateAfterMove: clone,
			})
		}
	}

	return movesFromRoom
}

func getValidMovesFromHall(data amphipod.Burrow) []Move {
	movesFromHall := make([]Move, 0)

	for i := range data.Hall {
		// if empty
		if data.Hall[i] == '.' {
			continue
		}

		pod := data.Hall[i]
		roomNum := amphipod.RoomByName(pod)
		room := data.Rooms[roomNum]

		depth := amphipod.RoomSize - 1
		for depth >= 0 {
			if room[depth] == pod {
				depth++
				continue
			}

			if room[depth] == '.' {
				break
			}

			depth = -1
		}

		// cannot move if other types of amphipod present in room
		if depth == -1 {
			continue
		}

		cost := getMoveCost(i, roomNum, depth, data, pod)

		// no path or no place in room
		if cost == math.MaxUint64 {
			continue
		}

		clone := data.Clone()
		clone.Hall[i] = '.'
		clone.Rooms[roomNum][depth] = pod

		movesFromHall = append(movesFromHall, Move{
			TotalCost:      cost,
			StateAfterMove: clone,
		})
	}

	return movesFromHall
}

func getMoveCost(hall, room, depth int, burrow amphipod.Burrow, r rune) uint64 {
	var start, end int
	if hall/2 < room+1 {
		start = hall
		end = room + 1
	} else {
		start = room + 1
		end = hall
	}

	for _, r := range burrow.Hall[start:end] {
		// hall is blocked
		if r != '.' {
			return math.MaxUint64
		}
	}

	cost := (end - start + (depth + 1)) * amphipod.Cost[r]

	return uint64(cost)
}
