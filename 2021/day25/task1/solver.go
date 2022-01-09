package task1

import (
	"aoc-2021-day25/input"
)

func Solve(data input.Data) (int, error) {
	old, new := clone(data.Cucumbers), emptyWithSizeOf(data.Cucumbers)
	moves := 0

	for !identical(old, new) {
		if moves > 0 {
			old = clone(new)
		} else {
			new = clone(old)
		}

		moveEast(old, new)

		movedEast := clone(new)

		moveSouth(movedEast, new)
		moves++
	}

	return moves, nil
}

func moveSouth(old, new [][]rune) {
	for j := range old[0] {
		for i := 1; i < len(old); i++ {
			prev := i - 1
			if old[prev][j] == 'v' && old[i][j] == '.' {
				new[i][j] = 'v'
				new[prev][j] = '.'
			}
		}

		if old[0][j] == '.' && old[len(old)-1][j] == 'v' {
			new[0][j] = 'v'
			new[len(old)-1][j] = '.'
		}
	}
}

func moveEast(old, new [][]rune) {
	for i := range old {
		for j := 1; j < len(old[i]); j++ {
			prev := j - 1
			if old[i][prev] == '>' && old[i][j] == '.' {
				new[i][j] = '>'
				new[i][prev] = '.'
			}
		}

		if old[i][0] == '.' && old[i][len(old[i])-1] == '>' {
			new[i][0] = '>'
			new[i][len(old[i])-1] = '.'
		}
	}
}

func identical(old, new [][]rune) bool {
	for i := range old {
		for j := range old[i] {
			if old[i][j] != new[i][j] {
				return false
			}
		}
	}

	return true
}

func clone(old [][]rune) [][]rune {
	new := make([][]rune, len(old))

	for i := range old {
		row := make([]rune, len(old[i]))
		for j := range old[i] {
			row[j] = old[i][j]
		}

		new[i] = row
	}

	return new
}

func emptyWithSizeOf(old [][]rune) [][]rune {
	new := make([][]rune, len(old))

	for i := range old {
		row := make([]rune, len(old[i]))
		new[i] = row
	}

	return new
}
