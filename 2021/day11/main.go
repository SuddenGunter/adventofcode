package main

import (
	"aoc-2021-day11/input"
	"aoc-2021-day11/task1"
	"aoc-2021-day11/task2"
	"fmt"
)

func main() {
	data, err := input.ParseInput("data.txt")
	if err != nil {
		panic(err)
	}

	t2octopuses := make([][]byte, 0, len(data.Octopuses))
	for i := range data.Octopuses {
		line := make([]byte, 0, len(data.Octopuses[i]))
		for _, v := range data.Octopuses[i] {
			line = append(line, v)
		}

		t2octopuses = append(t2octopuses, line)
	}
	t2data := input.Data{Octopuses: t2octopuses}

	t1, err := task1.Solve(data)
	fmt.Printf("task 1 answer: %v, err: %v\n", t1, err)

	t2, t2f, err := task2.Solve(t2data)
	fmt.Printf("task 2 answer: %v, total flashes: %v, err: %v\n", t2, t2f, err)
}
