package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := readInput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("task 1 answer: %v\n", SolveTask(data.Clone(), 1))
	fmt.Printf("task 2 answer: %v\n", SolveTask(data.Clone(), 2))
}

func readInput() (*Data, error) {
	file, err := os.Open("demo.data.txt")
	if err != nil {
		return nil, fmt.Errorf("read input: file open: %w", err)
	}

	data := Data{
		Valves: make([]Valve, 0),
	}

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		line := scan.Text()
		valve, err := parseValve(line)
		if err != nil {
			return nil, fmt.Errorf("parse valve: %w", err)
		}

		data.Valves = append(data.Valves, valve)
	}

	return &data, nil
}

func parseValve(line string) (Valve, error) {
	line = strings.TrimPrefix(line, "Valve ")
	name := line[0:2]
	line = line[2:]
	line = strings.TrimPrefix(line, " has flow rate=")
	idx := strings.Index(line, ";")
	flowRate, err := strconv.Atoi(line[:idx])
	if err != nil {
		return Valve{}, fmt.Errorf("parser failed on line: %v", line)
	}
	line = line[idx:]
	line = strings.TrimPrefix(line, " tunnels lead to valves")
	res := valveName.FindAllStringSubmatch(line, -1)

	tunnels := make([]string, 0)
	for i := range res {
		tunnels = append(tunnels, res[i][0])
	}

	return Valve{
		Name:     name,
		FlowRate: flowRate,
		Tunnels:  tunnels,
	}, nil
}

var valveName = regexp.MustCompile("([A-Z]{2})")
