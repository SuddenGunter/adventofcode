package main

import (
	"aoc-2021-day24/task1"
	"aoc-2021-day24/task2"
	"fmt"
)

func main() {
	t1, err := task1.Solve("data.txt")
	fmt.Printf("task 1 answer: %v, err: %v\n", t1, err)

	t2, err := task2.Solve("data.txt")
	fmt.Printf("task 1 answer: %v, err: %v\n", t2, err)
}
