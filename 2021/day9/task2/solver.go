package task2

import (
	"aoc-2021-day9/input"
	"sort"
)

type position struct {
	X int
	Y int
}

func Solve(data input.Data) (int, error) {
	minByLine := minimumsByLine(data.Numbers)
	minByLineAndCol := minimums(minByLine, data.Numbers)
	basins := basins(minByLineAndCol, data.Numbers)

	return multiply(basins), nil
}

func basins(mins map[position]struct{}, numbers [][]byte) []int {
	basins := make([]int, 0, len(mins))

	for k := range mins {
		visitedNodes := make(map[position]struct{})
		toVisit := &stack{data: make([]position, 0)}

		toVisit.push(k)

		for !toVisit.isEmpty() {
			current := toVisit.pop()
			visitedNodes[current] = struct{}{}

			byXY := append(directionsToGoByX(current, numbers), directionsToGoByY(current, numbers)...)
			for _, pos := range byXY {
				if _, found := visitedNodes[pos]; found {
					continue
				}

				if numbers[pos.X][pos.Y] == 9 {
					continue
				}

				toVisit.push(pos)
			}
		}

		basins = append(basins, len(visitedNodes))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	return basins[0:3]
}

func multiply(basins []int) int {
	power := basins[0]
	for _, v := range basins[1:] {
		power *= v
	}

	return power
}

func minimums(minByLine map[position]struct{}, numbers [][]byte) map[position]struct{} {
	localMinimums := make(map[position]struct{})

	for k, v := range minByLine {
		if trueMinimum(k, numbers) {
			localMinimums[k] = v
		}
	}

	return localMinimums
}

func directionsToGoByX(k position, numbers [][]byte) []position {
	directions := make([]position, 0)

	switch k.X {
	case 0:
		directions = append(directions, position{X: k.X + 1, Y: k.Y})
	case len(numbers) - 1:
		directions = append(directions, position{X: k.X - 1, Y: k.Y})
	default:
		directions = append(directions, position{X: k.X - 1, Y: k.Y}, position{X: k.X + 1, Y: k.Y})
	}

	return directions
}

func directionsToGoByY(k position, numbers [][]byte) []position {
	directions := make([]position, 0)

	switch k.Y {
	case 0:
		directions = append(directions, position{X: k.X, Y: k.Y + 1})
	case len(numbers[k.X]) - 1:
		directions = append(directions, position{X: k.X, Y: k.Y - 1})
	default:
		directions = append(directions, position{X: k.X, Y: k.Y - 1}, position{X: k.X, Y: k.Y + 1})
	}

	return directions
}

func trueMinimum(k position, numbers [][]byte) bool {
	switch k.X {
	case 0:
		return numbers[k.X][k.Y] < numbers[k.X+1][k.Y]
	case len(numbers) - 1:
		return numbers[k.X][k.Y] < numbers[k.X-1][k.Y]
	default:
		return numbers[k.X][k.Y] < numbers[k.X-1][k.Y] && numbers[k.X][k.Y] < numbers[k.X+1][k.Y]
	}
}

func minimumsByLine(lines [][]byte) map[position]struct{} {
	localMinimums := make(map[position]struct{})

	for i := 0; i < len(lines); i++ {
		// ignore edge columns, treat them as special case
		for j := 1; j < len(lines[i])-1; j++ {
			x := lines[i][j]

			if x < lines[i][j-1] && x < lines[i][j+1] {
				localMinimums[position{
					X: i,
					Y: j,
				}] = struct{}{}

				// if x is local minimum - next element never can be a local minimum
				j++
			}
		}
	}

	// Compare edge columns in rows. Will only work if all lines have equal length
	for i, lCol, rCol := 0, 0, len(lines[0])-1; i < len(lines); i++ {
		left := lines[i][lCol]

		if left < lines[i][lCol+1] {
			localMinimums[position{
				X: i,
				Y: lCol,
			}] = struct{}{}
		}

		right := lines[i][rCol]
		if right < lines[i][rCol-1] {
			localMinimums[position{
				X: i,
				Y: rCol,
			}] = struct{}{}
		}
	}

	return localMinimums
}
