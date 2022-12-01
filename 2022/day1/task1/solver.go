package task1

func Solve(input map[int][]int) (int, error) {
	sum := sumByElf(input)
	return maxByElf(sum), nil
}

func maxByElf(sum map[int]int) int {
	max := sum[0]
	for _, v := range sum {
		if v > max {
			max = v
		}
	}

	return max
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
