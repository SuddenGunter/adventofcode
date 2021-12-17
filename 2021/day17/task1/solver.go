package task1

import (
	"aoc-2021-day17/input"
	"errors"
	"math"
)

func Solve(data input.Data) (float64, error) {
	if data.YLow >= 0 {
		return 0, errors.New("this solution will not work for your input")
	}

	lastJumpSize := math.Abs(float64(data.YLow)) - 1

	// gauss summation
	return lastJumpSize * (lastJumpSize + 1) / 2, nil
}
