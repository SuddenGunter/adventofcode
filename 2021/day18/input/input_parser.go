package input

import (
	"aoc-2021-day18/tree"
	"os"
	"strings"
)

type Data struct {
	Numbers []tree.Node
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	nodes, err := tree.ParseNodes(lines)
	if err != nil {
		return Data{}, err
	}

	return Data{Numbers: nodes}, err
}
