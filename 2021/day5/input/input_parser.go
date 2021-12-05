package input

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const bitSize = 16

type Data struct {
	Lines []Line
}

type Line struct {
	P1 Position
	P2 Position
}

type Position struct {
	X int16
	Y int16
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	// todo: what if line empty? (last line)

	data := Data{
		Lines: make([]Line, 0, len(lines)),
	}

	for i, l := range lines {
		if l == "" {
			continue
		}

		line, err := parseLine(l)
		if err != nil {
			return Data{}, fmt.Errorf("failed to parse line %v: %w", i, err)
		}

		data.Lines = append(data.Lines, line)
	}

	return data, nil
}

func parseLine(l string) (Line, error) {
	halves := strings.Split(l, "->")

	if halves[0] == "" {
		return Line{}, errors.New("line doesn't match format")
	}

	left, err := parseHalf(halves[0])
	if err != nil {
		return Line{}, err
	}

	right, err := parseHalf(halves[1])
	if err != nil {
		return Line{}, err
	}

	return Line{
		P1: left,
		P2: right,
	}, nil
}

func parseHalf(h string) (Position, error) {
	nums := strings.Split(h, ",")

	x, err := strconv.ParseInt(strings.Trim(nums[0], " "), 10, bitSize)
	if err != nil {
		return Position{}, err
	}

	y, err := strconv.ParseInt(strings.Trim(nums[1], " "), 10, bitSize)
	if err != nil {
		return Position{}, err
	}

	return Position{X: int16(x), Y: int16(y)}, nil
}
