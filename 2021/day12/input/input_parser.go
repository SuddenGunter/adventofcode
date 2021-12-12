package input

import (
	"aoc-2021-day12/cave"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	Caves map[string]cave.Cave
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	data := Data{
		Caves: make(map[string]cave.Cave),
	}

	for i, l := range lines[:len(lines)-1] {
		from, to, err := parseLine(l)
		if err != nil {
			return Data{}, fmt.Errorf("failed to parse line %v: %w", i, err)
		}

		appendConnection(&data, from, to)
		appendConnection(&data, to, from)
	}

	return data, nil
}

func appendConnection(d *Data, from, to string) {
	_, found := d.Caves[from]
	if !found {
		d.Caves[from] = cave.Cave{
			Size:        cave.SizeByName(from),
			ConnectedTo: make([]string, 0, 1),
		}
	}

	fromCave := d.Caves[from]
	fromCave.ConnectedTo = append(fromCave.ConnectedTo, to)
	d.Caves[from] = fromCave
}

func parseLine(l string) (string, string, error) {
	split := strings.Split(l, "-")
	return split[0], split[1], nil
}
