package main

import (
	"aoc-2022-day1/task1"
	"aoc-2022-day1/task2"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := readInput()
	if err != nil {
		fmt.Println(err)
	}

	t1, err := task1.Solve(data)
	fmt.Printf("task 1 answer: %v, err: %v\n", t1, err)

	t2, err := task2.Solve(data)
	fmt.Printf("task 2 answer: %v, err: %v\n", t2, err)
}

func readInput() (map[int][]int, error) {
	data := make(map[int][]int)

	f, err := os.Open("data.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	elf := 0
	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "" {
			elf++
			continue
		}

		num, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, fmt.Errorf("failed to convert string to int: %w", err)
		}

		data = addTo(data, elf, num)
	}

	if err = s.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %w", err)
	}

	return data, nil
}

func addTo(data map[int][]int, elf int, val int) map[int][]int {
	if data[elf] == nil {
		data[elf] = make([]int, 0, 1)
	}

	data[elf] = append(data[elf], val)

	return data
}
