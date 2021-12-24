package task1

import "aoc-2021-day10/input"

var (
	openBrackets = map[rune]struct{}{
		'[': {},
		'<': {},
		'(': {},
		'{': {},
	}

	closeToOpenBrackets = map[rune]rune{
		']': '[',
		'>': '<',
		')': '(',
		'}': '{',
	}

	costPerBracket = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
)

func Solve(data input.Data) (int, error) {
	wrongBracketsCount := make(map[rune]int, 0)

	for _, line := range data.Lines {
		brackets := stack{data: make([]rune, 0, len(data.Lines))}

		for _, r := range line {
			if isOpenBracket(r) {
				brackets.push(r)
				continue
			}

			if brackets.isEmpty() || closeToOpenBrackets[r] != brackets.pop() {
				wrongBracketsCount[r]++
				break
			}
		}

		if !brackets.isEmpty() {
			// todo: line is not corrupted, but incomplete
		}
	}

	cost := analyzeCost(wrongBracketsCount)

	return cost, nil
}

func analyzeCost(count map[rune]int) int {
	totalCost := 0
	for k, v := range count {
		totalCost += costPerBracket[k] * v
	}

	return totalCost
}

func isOpenBracket(r rune) bool {
	_, found := openBrackets[r]
	return found
}
