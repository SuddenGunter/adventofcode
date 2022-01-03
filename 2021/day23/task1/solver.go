package task1

import (
	"aoc-2021-day23/task1/amphipod"
	"math"
)

type Move struct {
	TotalCost      uint64
	StateAfterMove amphipod.Burrow
}

func Solve(data amphipod.Burrow) (int, error) {
	validMoves := getValidMoves(data)
	return 0, nil
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

			// todo j is the depth!
			_, cost := getMoveCost()

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

		// cannot move if other types of amphipod present in room
		if !containsOnlyOwners(room, pod) {
			continue
		}

		depth, cost := getMoveCost()

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

func getMoveCost() (int, uint64) {
	return 0, math.MaxUint64
}

func containsOnlyOwners(runes [amphipod.RoomSize]rune, pod rune) bool {
	for _, r := range runes {
		if r != '.' && r != pod {
			return false
		}
	}

	return true
}
