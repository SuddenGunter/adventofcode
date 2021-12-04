package task1

import "aoc-2021-day4/input"

const markedNumber = -1

// Solve simulates bingo game until first winner and return sum of all unmarked numbers on
// winners board.
func Solve(data input.Data) (int, error) {
	winners := make([]input.Player, 0)
	hasWinner := false
	lastWinningNumber := -1

	for _, v := range data.WinningSequence {
		for _, p := range data.Players {
			won := markNumber(p, v)
			if won {
				hasWinner = true
				winners = append(winners, p)
			}
		}

		if hasWinner {
			lastWinningNumber = v
			break
		}
	}

	maxScore := getBiggestWinner(winners)
	return maxScore * lastWinningNumber, nil
}

func getBiggestWinner(winners []input.Player) int {
	maxScore := -1
	for _, w := range winners {
		score := getScore(w)
		maxScore = max(maxScore, score)
	}

	return maxScore
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getScore(p input.Player) int {
	score := 0

	for i := range p.Data {
		for _, v := range p.Data[i] {
			if v != markedNumber {
				score += v
			}
		}
	}

	return score
}

func markNumber(p input.Player, num int) bool {
	pos, found := p.Index[num]
	if !found {
		return false
	}

	p.Data[pos.X][pos.Y] = markedNumber

	wonByRow := checkIfWon(p.Data[pos.X])
	wonByCol := checkIfWon(getCol(p.Data, pos.Y))

	return wonByCol || wonByRow
}

func getCol(data [][]int, y int) []int {
	res := make([]int, input.Size)

	for i := range data {
		res[i] = data[i][y]
	}

	return res
}

func checkIfWon(data []int) bool {
	for _, v := range data {
		if v != markedNumber {
			return false
		}
	}

	return true
}
