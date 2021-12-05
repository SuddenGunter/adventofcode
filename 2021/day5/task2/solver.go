package task2

import "aoc-2021-day5/input"

const maxSize = 1000

func Solve(data input.Data) (int, error) {
	board := draw(data.Lines)

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
		drawByXY(board, l)
	}

	return board
}

func drawByXY(board [][]int16, l input.Line) {
	if samePoint(l.P1, l.P2) {
		board[l.P1.X][l.P1.Y]++
		return
	}

	pos := input.Position{
		X: l.P1.X,
		Y: l.P1.Y,
	}

	for !samePoint(pos, l.P2) {
		board[pos.X][pos.Y]++
		pos = moveTowards(pos, l.P2)
	}

	// line ending also must be drawn.
	board[pos.X][pos.Y]++
}

func moveTowards(from, to input.Position) input.Position {
	switch compare(from.X, to.X) {
	case 1:
		from.X--
	case -1:
		from.X++
	default:
		// already in the same point
	}

	switch compare(from.Y, to.Y) {
	case 1:
		from.Y--
	case -1:
		from.Y++
	default:
		// already in the same point
	}

	return from
}

func compare(a, b int16) int {
	if a == b {
		return 0
	}

	if a > b {
		return 1
	} else {
		return -1
	}
}

func samePoint(p1, p2 input.Position) bool {
	return p1.X == p2.X && p1.Y == p2.Y
}
