package task1

import "aoc-2021-day9/input"

type position struct {
	X int
	Y int
}

func Solve(data input.Data) (int, error) {
	minByLine := minimumsByLine(data.Numbers)
	minByLineAndCol := minimums(minByLine, data.Numbers)

	return sumRisk(minByLineAndCol, data.Numbers), nil
}

func sumRisk(minimums map[position]struct{}, numbers [][]byte) int {
	sum := 0
	for k := range minimums {
		// risk level of a low point is 1 plus its height.
		sum += int(numbers[k.X][k.Y]) + 1
	}

	return sum
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
