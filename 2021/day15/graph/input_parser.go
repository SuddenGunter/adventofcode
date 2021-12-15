package graph

import (
	"os"
	"strings"
)

type Data struct {
	Graph  *Graph
	Finish *Node
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	for i := range lines[:len(lines)-1] {
		numbersInLine := parseLine(lines[i])
		for j := range numbersInLine {

		}
	}

	return data, nil
}

func parseLine(l string) []int {
	numbers := make([]int, 0, len(l)-1)
	for _, n := range l {
		numbers = append(numbers, int(n-'0'))
	}

	return numbers
}
