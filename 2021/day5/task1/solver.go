package task1

import "aoc-2021-day5/input"

const maxSize = 1000

func Solve(data input.Data) (int, error) {
	filtered := filter(data.Lines)
	sorted := sort(filtered)
	board := draw(sorted)

	return countAboveOne(board)
}

func countAboveOne(board [][]int16) (int, error) {
	counter := 0

	for _, arr := range board {
		for _, x := range arr {
			if x > 1 {
				counter++
			}
		}
	}

	return counter, nil
}

func draw(lines []input.Line) [][]int16 {
	board := make([][]int16, 0, maxSize)

	for i := 0; i < maxSize; i++ {
		board = append(board, make([]int16, maxSize))
	}

	for _, l := range lines {
		if l.P1.X == l.P2.X {
			drawByY(board, l)
		} else {
			drawByX(board, l)
		}
	}

	return board
}

func drawByX(board [][]int16, l input.Line) {
	if l.P1.X == l.P2.X {
		board[l.P1.X][l.P1.Y]++
		return
	}

	for posX := l.P1.X; posX <= l.P2.X; posX++ {
		board[posX][l.P1.Y]++
	}
}

func drawByY(board [][]int16, l input.Line) {
	if l.P1.Y == l.P2.Y {
		board[l.P1.X][l.P1.Y]++
		return
	}

	for posY := l.P1.Y; posY <= l.P2.Y; posY++ {
		board[l.P1.X][posY]++
	}
}

func sort(filtered []input.Line) []input.Line {
	sorted := make([]input.Line, 0, len(filtered))

	for _, l := range filtered {
		tmp := l
		if l.P1.X > l.P2.X {
			tmp.P1, tmp.P2 = l.P2, l.P1
		} else {
			if l.P1.Y > l.P2.Y {
				tmp.P1, tmp.P2 = l.P2, l.P1
			}
		}

		sorted = append(sorted, tmp)
	}

	return sorted
}

func filter(lines []input.Line) []input.Line {
	filtered := make([]input.Line, 0)

	for _, l := range lines {
		if l.P1.X == l.P2.X {
			filtered = append(filtered, l)
		}

		if l.P1.Y == l.P2.Y {
			filtered = append(filtered, l)
		}
	}

	return filtered
}
