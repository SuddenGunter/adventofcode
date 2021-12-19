package main

import (
	"aoc-2021-day19/input"
	"aoc-2021-day19/task1"
	"aoc-2021-day19/task2"
	"fmt"
)

func main() {
	data, err := input.ParseInput("data.txt")
	if err != nil {
		panic(err)
	}

	t1, scanners, err := task1.Solve(data)
	fmt.Printf("task 1 answer: %v, err: %v\n", t1, err)

	t2, err := task2.Solve(scanners)
	fmt.Printf("task 2 answer: %v, err: %v\n", t2, err)
}
