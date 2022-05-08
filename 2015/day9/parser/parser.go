package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"aoc-2015-day9/graph"
)

func Lines(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	data := make([]string, 0)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, nil
}

func Data(fileName string) (*graph.Graph, error) {
	lines, err := Lines(fileName)
	if err != nil {
		return nil, err
	}

	g := graph.New()

	for _, l := range lines {
		from, to, weight, err := parseLine(l)
		if err != nil {
			return nil, err
		}

		g.AddEdge(from, to, weight)
		g.AddEdge(to, from, weight)
	}

	return g, nil
}

func parseLine(l string) (string, string, int, error) {
	split := strings.Split(l, "=")
	w := strings.TrimSpace(split[1])
	weight, err := strconv.Atoi(w)
	if err != nil {
		return "", "", 0, err
	}

	path := strings.Split(split[0], "to")

	return strings.TrimSpace(path[0]), strings.TrimSpace(path[1]), weight, nil
}
