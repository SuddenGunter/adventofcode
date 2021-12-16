package main

import (
	"aoc-2021-day16/task1"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := getInput()
	if err != nil {
		panic(err)
	}

	t1, err := task1.Solve(data)
	fmt.Printf("task 1 answer: %v, err: %v\n", t1, err)

	//t2, err := task2.Solve(data)
	//fmt.Printf("task 2 answer: %v, err: %v\n", t2, err)
}

func getInput() (string, error) {
	file, err := os.ReadFile("data.txt")
	if err != nil {
		return "", err
	}

	return strings.Trim(string(file), "\n "), nil
}
