package task2

import "sort"

func Solve(input map[int][]int) (int, error) {
	sum := sumByElf(input)
	return maxByTopThreeElves(sum), nil
}

func maxByTopThreeElves(sum map[int]int) int {
	values := make([]int, 0, len(sum))
	for _, v := range sum {
		values = append(values, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	return values[0] + values[1] + values[2]
}

func sumByElf(input map[int][]int) map[int]int {
	result := make(map[int]int)
	for elf, data := range input {
		for _, v := range data {
			result[elf] += v
		}
	}

	return result
}
