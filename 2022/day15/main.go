package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := readInput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("task 1 answer: %v\n", SolveTask(data.Clone(), 1))
	fmt.Printf("task 2 answer: %v\n", SolveTask(data.Clone(), 2))
}

func readInput() (*Task, error) {
	f, err := os.Open("data.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	pairs := make([]Pair, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}

		pairs = append(pairs, parsePair(s.Text()))
	}

	if err = s.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %w", err)
	}

	return &Task{
		Pairs: pairs,
	}, nil
}

func parsePair(text string) Pair {
	rgx := regexp.MustCompile(`-?\d+`)
	matches := rgx.FindAllSubmatch([]byte(text), -1)
	nums := make([]int, 0, 4)
	for _, x := range matches {
		num, err := strconv.Atoi(string(x[0]))
		if err != nil {
			log.Fatal(err)
		}

		nums = append(nums, num)
	}

	return Pair{
		Sensor: Point{
			X: nums[0],
			Y: nums[1],
		}, Beacon: Point{
			X: nums[2],
			Y: nums[3],
		}}
}
