package input

import (
	"aoc-2021-day22/geometry"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Operations []geometry.Operation
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	lines = lines[:len(lines)-1]

	operations := make([]geometry.Operation, 0, len(lines))

	for _, l := range lines {
		op, err := parseLine(l)
		if err != nil {
			return Data{}, err
		}

		operations = append(operations, op)
	}

	return Data{Operations: operations}, nil
}

func parseLine(l string) (geometry.Operation, error) {
	split := strings.Split(l, ",")
	x := split[0]
	y := split[1]
	z := split[2]

	op := geometry.Operation{}
	if strings.HasPrefix(x, "on") {
		op.On = true
		x = strings.TrimPrefix(x, "on ")
	} else {
		x = strings.TrimPrefix(x, "off ")
	}

	x1, x2, err := parseNumbers(x, "x")
	if err != nil {
		return geometry.Operation{}, err
	}

	y1, y2, err := parseNumbers(y, "y")
	if err != nil {
		return geometry.Operation{}, err
	}

	z1, z2, err := parseNumbers(z, "z")
	if err != nil {
		return geometry.Operation{}, err
	}

	cuboid := geometry.Cuboid{
		LowerX: x1,
		UpperX: x2,
		LowerY: y1,
		UpperY: y2,
		LowerZ: z1,
		UpperZ: z2,
	}

	op.Cuboid = cuboid

	return op, nil
}

func parseNumbers(numbers, letter string) (int64, int64, error) {
	numbers = strings.TrimPrefix(numbers, fmt.Sprintf("%v=", letter))
	split := strings.Split(numbers, "..")

	first, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, err
	}

	second, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, err
	}

	fi64, si64 := minFirst(int64(first), int64(second))

	return fi64, si64, nil
}

func minFirst(a, b int64) (int64, int64) {
	if a < b {
		return a, b
	}

	return b, a
}
