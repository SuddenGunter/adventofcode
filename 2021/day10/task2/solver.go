package task2

import (
	"aoc-2021-day10/input"
	"sort"
)

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
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
)

func Solve(data input.Data) (int, error) {
	costs := make([]int, 0)

	for _, line := range data.Lines {
		brackets := &stack{data: make([]rune, 0, len(data.Lines))}
		validString := true

		for _, r := range line {
			if isOpenBracket(r) {
				brackets.push(r)
				continue
			}

			if brackets.isEmpty() || closeToOpenBrackets[r] != brackets.pop() {
				validString = false
				break
			}
		}

		if validString {
			costs = append(costs, analyzeCost(brackets))
		}
	}

	sort.Ints(costs)
	// get middle. costs could only be odd (assignment limitation).
	cost := costs[len(costs)/2]

	return cost, nil
}

func analyzeCost(brackets *stack) int {
	totalCost := 0

	for !brackets.isEmpty() {
		totalCost *= 5

		// we do not actually need to get what closing bracket we need for auto-completion. we could infer
		// cost from opening bracket.
		totalCost += costPerBracket[brackets.pop()]
	}

	return totalCost
}

func isOpenBracket(r rune) bool {
	_, found := openBrackets[r]
	return found
}
