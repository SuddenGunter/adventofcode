package main

import (
	"aoc-2021-day23/task1"
	"aoc-2021-day23/task2"
	"fmt"
)

func main() {
	t1 := task1.Solve("data.txt")
	fmt.Printf("task 1 answer: %v\n", t1)

	t2 := task2.Solve("data.txt")
	fmt.Printf("task 2 answer: %v\n", t2)
}
