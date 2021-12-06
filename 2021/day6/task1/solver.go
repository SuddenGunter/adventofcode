package task1

import (
	"aoc-2021-day6/input"
)

const (
	defaultAfterBirth = 6
	defaultForNewFish = 8
	daysToSimulate    = 80
)

func Solve(data input.Data) (int, error) {
	sum := 0
	for _, f := range data.Fishes {
		res := simulateFishAndDescendants(f, daysToSimulate)
		sum += res
	}

	return sum, nil
}

func simulateFishAndDescendants(f input.Fish, days int) int {
	descendants := []input.Fish{f}

	for day := 1; day <= days; day++ {
		addFishes := 0

		for i := range descendants {
			// no new fish spawns
			if descendants[i].DaysBeforeBirth > 0 {
				descendants[i].DaysBeforeBirth--
				continue
			}

			// new fish spawns
			descendants[i].DaysBeforeBirth = defaultAfterBirth
			addFishes++
		}

		for addFishes > 0 {
			descendants = append(descendants, input.Fish{DaysBeforeBirth: defaultForNewFish})
			addFishes--
		}
	}

	return len(descendants)
}
