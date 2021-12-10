package main

import (
	"aoc-2021-day10/input"
	"aoc-2021-day10/task1"
	"aoc-2021-day10/task2"
	"fmt"
)

func main() {
	data, err := input.ParseInput("data.txt")
	if err != nil {
		panic(err)
	}

	t1, err := task1.Solve(data)
	fmt.Printf("task 1 answer: %v, err: %v\n", t1, err)

	t2, err := task2.Solve(data)
	fmt.Printf("task 2 answer: %v, err: %v\n", t2, err)
}
