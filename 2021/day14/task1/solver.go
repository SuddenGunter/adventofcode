package task1

import (
	"aoc-2021-day14/input"
	"fmt"
	"math"
)

const steps = 10

func Solve(data input.Data) (int, error) {
	counter := make(map[string]int)
	for _, c := range data.InitialState {
		counter[string(c)]++
	}

	pairs := getMap(data.InitialState)

	for i := 0; i < steps; i++ {
		copyOfPairs := copyM(pairs)
		for _, rule := range data.Changes {
			val := pairs[rule.From]

			copyOfPairs[rule.From] -= val
			counter[string(rule.To)] += val
			copyOfPairs[fmt.Sprintf("%s%s", string(rule.From[0]), rule.To)] += val
			copyOfPairs[fmt.Sprintf("%s%s", rule.To, string(rule.From[1]))] += val
		}

		pairs = copyOfPairs
	}

	max, min := math.MinInt, math.MaxInt
	for _, v := range counter {
		if v >= max {
			max = v
		}
		if v <= min {
			min = v
		}
	}

	return max - min, nil
}

func getMap(state string) map[string]int {
	pairs := make(map[string]int)

	for i := 0; i < len(state)-1; i++ {
		pairs[string(state[i:i+2])]++
	}

	return pairs
}

func copyM(m map[string]int) map[string]int {
	newMap := make(map[string]int, len(m))
	for k, v := range m {
		newMap[k] = v
	}

	return newMap
}
