package main

import (
	"aoc-2015-day9/task2"
	"fmt"

	"aoc-2015-day9/parser"
	"aoc-2015-day9/task1"
)

func main() {
	data, err := parser.Data("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("task #1 solution: ", task1.Solve(data))
	fmt.Println("task #2 solution: ", task2.Solve(data))
}
