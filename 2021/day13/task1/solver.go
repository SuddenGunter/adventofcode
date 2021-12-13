package task1

import (
	"aoc-2021-day13/fold"
	"aoc-2021-day13/input"
)

func Solve(data input.Data) (int, error) {
	dots := copyDots(data.Dots)
	res := execute(dots, data.Folds[0])

	return countHashes(res), nil
}

func countHashes(res [][]rune) int {
	counter := 0
	for i := range res {
		for _, v := range res[i] {
			if v == '█' {
				counter++
			}
		}
	}

	return counter
}

func execute(dots [][]rune, f fold.Fold) [][]rune {
	if f.Direction == fold.X {
		for i := range dots {
			for j := 0; j < f.Value; j++ {
				if dots[i][len(dots[i])-(j+1)] == '█' {
					dots[i][j] = '█'
				}
			}

			dots[i] = dots[i][:f.Value]
		}
	}

	if f.Direction == fold.Y {
		for i := 0; i < f.Value; i++ {
			for j := range dots[i] {
				if dots[len(dots)-(i+1)][j] == '█' {
					dots[i][j] = '█'
				}
			}
		}

		dots = dots[:f.Value]
	}

	return dots
}

func copyDots(dots [][]rune) [][]rune {
	tmp := make([][]rune, 0, len(dots))
	for i := range dots {
		slice := make([]rune, len(dots[i]))
		for j := range dots[i] {
			slice[j] = dots[i][j]
		}

		tmp = append(tmp, slice)
	}

	return tmp
}
