package graph

import (
	"fmt"
	"os"
	"strings"
)

type Data struct {
	Graph  *Graph
	Start  *Node
	Finish *Node
}

func ParseInputForTask1(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	g := NewGraph()
	var start, finish *Node

	for i := range lines[:len(lines)-1] {
		numbersInLine := parseLine(lines[i])
		for j := range numbersInLine {
			_, err = g.Add(Position{
				X: i,
				Y: j,
			}, numbersInLine[j])

			if err != nil {
				return Data{}, err
			}
		}

		// last line
		if i == 0 {
			start = g.Nodes[Position{X: 0, Y: 0}]
		}
		// last line
		if i == len(lines)-2 {
			finish = g.Nodes[Position{X: i, Y: len(numbersInLine) - 1}]
		}
	}

	return Data{
		Graph:  g,
		Start:  start,
		Finish: finish,
	}, nil
}

func ParseInputForTask2(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	g := NewGraph()
	var start, finish *Node

	for m := 0; m < 5; m++ {
		for i := range lines[:len(lines)-1] {
			numbersInLine := parseLine(lines[i])
			for n := 0; n < 5; n++ {
				for j := range numbersInLine {
					_, err = g.Add(Position{
						X: i + m,
						Y: j + n,
					}, numbersInLine[j]+m+n)
					fmt.Print(numbersInLine[j] + m + n)
					if err != nil {
						return Data{}, err
					}
				}
			}
			// last line
			if i == 0 && m == 0 {
				start = g.Nodes[Position{X: 0, Y: 0}]
			}

			// last line
			if i == len(lines)-2 && m == 4 {
				finish = g.Nodes[Position{X: i + m, Y: len(numbersInLine) - 1}]
			}
		}
	}

	return Data{
		Graph:  g,
		Start:  start,
		Finish: finish,
	}, nil
}

func parseLine(l string) []int {
	numbers := make([]int, 0, len(l)-1)
	for _, n := range l {
		numbers = append(numbers, int(n-'0'))
	}

	return numbers
}
