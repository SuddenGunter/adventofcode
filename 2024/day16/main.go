package main

import (
	"fmt"
	"log"
	"os"

	"aoc-2024-day16/task1"
	"aoc-2024-day16/task2"
)

func main() {
	f := readFile()
	fmt.Printf("task #1 solution: %v\n", task1.Solve(f))
	fmt.Printf("task #2 solution: %v\n", task2.Solve(f))
}

func readFile() []byte {
	res, err := os.ReadFile("demo.data.txt")
	if err != nil {
		log.Fatal(err)
	}

	return res
}
