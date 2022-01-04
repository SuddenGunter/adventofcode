package task1

import (
	"aoc-2021-day23/task1/amphipod"
	"math"
)

type Move struct {
	TotalCost      float64
	StateAfterMove amphipod.Burrow
}

var cache = make(map[amphipod.Burrow]float64, 3000)

func Solve(data amphipod.Burrow, oldStates []amphipod.Burrow) float64 {
	if done(data.Rooms) {
		return 0
	}

	res, found := cache[data]
	if found {
		return res
	}

	best := math.Inf(+1)

	moves := getValidMoves(data)

	//if data.HallEmpty() && (oldStates == nil || len(oldStates) == 0) {
	//	if len(moves) > 1 {
	//		for _, m := range moves {
	//			if m.StateAfterMove.Hall[3] == 66 && m.StateAfterMove.Rooms[2][0] == '.' {
	//				moves = []Move{m}
	//				break
	//			}
	//		}
	//	}
	//}
	//
	//if oldStates != nil && len(oldStates) == 1 {
	//	if len(moves) > 1 {
	//		for _, m := range moves {
	//			if m.StateAfterMove.Hall[3] == 66 && m.StateAfterMove.Hall[5] == 67 {
	//				moves = []Move{m}
	//				break
	//			}
	//		}
	//	}
	//
	//}

	//fmt.Println("valid moves")
	//for _, m := range moves {
	//	fmt.Println()
	//	fmt.Println(m.StateAfterMove)
	//}

	for _, move := range moves {
		cost := move.TotalCost

		newStates := make([]amphipod.Burrow, 0)
		if oldStates != nil {
			for _, v := range oldStates {
				newStates = append(newStates, v)
			}
		}

		//if len(newStates) >= 30 {
		//	fmt.Println("STEPS")
		//	fmt.Println(newStates)
		//	panic("panic")
		//}

		newStates = append(newStates, move.StateAfterMove)

		result := Solve(move.StateAfterMove, newStates)

		cache[move.StateAfterMove] = result
		// todo: what if result is max?
		cost += result

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

		nbr := amphipod.NameByRoom(i)
		containsOnlyOwners := true
		for z := j; z < amphipod.RoomSize; z++ {
			if room[z] != nbr {
				containsOnlyOwners = false
			}
		}

		if containsOnlyOwners {
			continue
		}

		for h := range data.Hall {
			if !amphipod.EnterableHallCell(h) || data.Hall[h] != '.' {
				continue
			}

			cost := getMoveCost(h, i, j, data, room[j])

			// no path or no place in cell
			if math.IsInf(cost, +1) {
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
				depth--
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
		if math.IsInf(cost, +1) {
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

func getMoveCost(hall, room, depth int, burrow amphipod.Burrow, r rune) float64 {
	var start, end int
	if hall < 2*(room+1) {
		start = hall + 1
		end = 2 * (room + 1)
	} else {
		start = 2 * (room + 1)
		end = hall - 1
	}

	for _, r := range burrow.Hall[start:end] {
		// hall is blocked
		if r != '.' {
			return math.Inf(+1)
		}
	}

	if hall < 2*(room+1) {
		start--
	} else {
		end++
	}

	cost := (end - start + (depth + 1)) * amphipod.Cost[r]

	return float64(cost)
}
