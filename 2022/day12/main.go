package main

import (
	"aoc-2022-day12/graph"
	"aoc-2022-day12/solver"
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	data, err := readInput()
	if err != nil {
		fmt.Println(err)
	}

	g := graph.NewGraph(data)

	fmt.Printf("task 1 answer: %v\n", solver.SolveTask(g, 1))
	fmt.Printf("task 2 answer: %v\n", solver.SolveTask(g, 2))
}

func readInput() ([][]rune, error) {
	f, err := os.Open("1data.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	lines := make([][]rune, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}

		lines = append(lines, bytes.Runes(s.Bytes()))
	}

	if err = s.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %w", err)
	}

	return lines, nil
}
