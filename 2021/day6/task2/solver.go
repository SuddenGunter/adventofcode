package task2

import "aoc-2021-day6/input"

const (
	defaultAfterBirth = 6
	defaultForNewFish = 8
	daysToSimulate    = 256
)

func Solve(data input.Data) (int64, error) {
	// counts how may fish we have. Element index represents how many days we have before new fish spawns.
	// inspired from: https://github.com/DarthSharkie/advent_of_code_2021/
	fishCounter := make([]int64, 9)

	for _, f := range data.Fishes {
		fishCounter[f.DaysBeforeBirth]++
	}

	for day := 1; day <= daysToSimulate; day++ {
		newFishes := fishCounter[0]

		for i := 0; i < len(fishCounter)-1; i++ {
			fishCounter[i] = fishCounter[i+1]
		}

		fishCounter[defaultAfterBirth] += newFishes
		fishCounter[defaultForNewFish] = newFishes
	}

	sum := int64(0)
	for _, v := range fishCounter {
		sum += v
	}

	return sum, nil
}
