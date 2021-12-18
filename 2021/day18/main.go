package main

import (
	"aoc-2021-day18/input"
	"aoc-2021-day18/task1"
	"aoc-2021-day18/task2"
	"fmt"
)

func main() {
	data, err := input.ParseInput("data.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Parsed: ")
	for _, v := range data.Numbers {
		fmt.Println(v)
	}

	t1, _, err := task1.Solve(data)
	fmt.Printf("task 1 answer: %v, err: %v\n", t1, err)

	t2, _, err := task2.Solve(data)
	fmt.Printf("task 2 answer: %v, err: %v\n", t2, err)
}
