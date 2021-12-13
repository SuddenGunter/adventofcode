package main

import (
	"aoc-2021-day13/input"
	"aoc-2021-day13/task1"
	"aoc-2021-day13/task2"
	"fmt"
)

func main() {
	data, err := input.ParseInput("data.txt")
	if err != nil {
		panic(err)
	}

	t1, err := task1.Solve(data)
	fmt.Printf("task 1 answer: %v, err: %v\n", t1, err)

	fmt.Println("task 2 answer: ")

	_, err = task2.Solve(data)
	fmt.Printf("err: %v\n", err)
}
