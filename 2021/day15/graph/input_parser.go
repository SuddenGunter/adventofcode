package graph

import (
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

		// first line
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

	rows := make([][]int, 0)

	for rowMod := 0; rowMod < 5; rowMod++ {
		for i := range lines[:len(lines)-1] {
			row := make([]int, 0)
			numbersInLine := parseLine(lines[i])

			for colMod := 0; colMod < 5; colMod++ {
				for j := range numbersInLine {
					num := numbersInLine[j] + colMod + rowMod
					if num > 9 {
						num = num % 9
					}

					row = append(row, num)
				}
			}

			rows = append(rows, row)
		}
	}

	//for _, r := range rows {
	//	fmt.Println(r)
	//}

	g := NewGraph()
	var start, finish *Node

	for i := range rows {
		for j := range rows[i] {
			_, err = g.Add(Position{
				X: i,
				Y: j,
			}, rows[i][j])

			if err != nil {
				return Data{}, err
			}
		}

		// first line
		if i == 0 {
			start = g.Nodes[Position{X: 0, Y: 0}]
		}
		// last line
		if i == len(rows)-1 {
			finish = g.Nodes[Position{X: i, Y: len(rows[i]) - 1}]
		}
	}

	return Data{
		Graph:  g,
		Start:  start,
		Finish: finish,
	}, nil

	return Data{}, nil
}

func parseLine(l string) []int {
	numbers := make([]int, 0, len(l)-1)
	for _, n := range l {
		numbers = append(numbers, int(n-'0'))
	}

	return numbers
}
