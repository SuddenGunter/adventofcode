package main

import (
	"aoc-2021-day25/input"
	"aoc-2021-day25/task1"
	"fmt"
)

func main() {
	data, err := input.ParseInput("data.txt")
	if err != nil {
		panic(err)
	}

	t1, err := task1.Solve(data)
	fmt.Printf("task 1 answer: %v, err: %v\n", t1, err)
}
