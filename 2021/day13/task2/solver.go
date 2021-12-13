package task2

import (
	"aoc-2021-day13/fold"
	"aoc-2021-day13/input"
	"fmt"
)

func Solve(data input.Data) (int, error) {
	dots := copyDots(data.Dots)
	for _, f := range data.Folds {
		dots = execute(dots, f)
	}

	prettyPrint(dots)
	return countHashes(dots), nil
}

func countHashes(res [][]string) int {
	counter := 0
	for i := range res {
		for _, v := range res[i] {
			if v == "#" {
				counter++
			}
		}
	}

	return counter
}

func execute(dots [][]string, f fold.Fold) [][]string {
	if f.Direction == fold.X {
		for i := range dots {
			for j := 0; j < f.Value; j++ {
				if dots[i][len(dots[i])-(j+1)] == "#" {
					dots[i][j] = "#"
				}
			}

			dots[i] = dots[i][:f.Value]
		}
	}

	if f.Direction == fold.Y {
		for i := 0; i < f.Value; i++ {
			for j := range dots[i] {
				if dots[len(dots)-(i+1)][j] == "#" {
					dots[i][j] = "#"
				}
			}

		}

		dots = dots[:f.Value]
	}

	return dots
}

func copyDots(dots [][]string) [][]string {
	tmp := make([][]string, 0, len(dots))
	for i := range dots {
		slice := make([]string, len(dots[i]))
		for j := range dots[i] {
			slice[j] = dots[i][j]
		}

		tmp = append(tmp, slice)
	}

	return tmp
}

func prettyPrint(dots [][]string) {
	for i := range dots {
		for _, c := range dots[i] {
			fmt.Print(c)
		}

		fmt.Println()
	}
}
