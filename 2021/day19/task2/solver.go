package task2

import (
	"aoc-2021-day19/point"
	"math"
)

func Solve(scanners map[point.Point3d]struct{}) (int, error) {
	max := math.MinInt
	for k := range scanners {
		for j := range scanners {
			if k == j {
				continue
			}

			distance := getManhattanDistance(k, j)
			if distance > max {
				max = distance
			}
		}
	}

	return max, nil
}

func getManhattanDistance(k, j point.Point3d) int {
	r := math.Abs(float64(k.X-j.X)) + math.Abs(float64(k.Y-j.Y)) + math.Abs(float64(k.Z-j.Z))

	return int(r)
}
