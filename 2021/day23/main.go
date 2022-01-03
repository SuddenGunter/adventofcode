package main

import (
	"aoc-2021-day23/task1"
	"aoc-2021-day23/task1/input"
	"fmt"
)

func main() {
	data, err := input.ParseInput("data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	t1 := task1.Solve(data, nil)
	fmt.Printf("task 1 answer: %v\n", t1)

	// t2, err := task2.Solve(data)
	// fmt.Printf("task 2 answer: %v, err: %v\n", t2, err)
}
