package input

import (
	"aoc-2021-day13/fold"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Dots  [][]rune
	Folds []fold.Fold
}

type point struct {
	X, Y int
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	dots, nextLine, err := parseDots(lines[:len(lines)-1])
	if err != nil {
		return Data{}, err
	}

	folds, err := parseFolds(lines[nextLine : len(lines)-1])
	if err != nil {
		return Data{}, err
	}

	return Data{
		Dots:  dots,
		Folds: folds,
	}, nil
}

func parseFolds(lines []string) ([]fold.Fold, error) {
	folds := make([]fold.Fold, 0, len(lines))
	for i, l := range lines {
		importantPart := strings.TrimPrefix(l, "fold along ")
		split := strings.Split(strings.Trim(importantPart, "\n "), "=")
		x, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, fmt.Errorf("parseFolds failed on line: %v with error: %w", i, err)
		}

		folds = append(folds, fold.Fold{
			Direction: fold.Direction(split[0]),
			Value:     x,
		})
	}

	return folds, nil
}

func parseDots(lines []string) ([][]rune, int, error) {
	points := make([]point, 0, len(lines))
	line := 0
	for ; lines[line] != ""; line++ {
		l := lines[line]
		split := strings.Split(strings.Trim(l, "\n "), ",")

		x, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, 0, fmt.Errorf("parseDots failed on line: %v with error: %w", line, err)
		}

		y, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, 0, err
		}

		points = append(points, point{
			X: x,
			Y: y,
		})
	}

	maxSizeX := maxX(points) + 1
	maxSizeY := maxY(points) + 1

	dots := make([][]rune, 0, maxSizeY)
	for i := 0; i < maxSizeY; i++ {
		dots = append(dots, make([]rune, maxSizeX))
		for j := range dots[i] {
			dots[i][j] = '.'
		}
	}
	for _, p := range points {
		dots[p.Y][p.X] = '#'
	}

	return dots, line + 1, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxX(points []point) int {
	maxVal := math.MinInt
	for _, p := range points {
		maxVal = max(maxVal, p.X)
	}

	return maxVal
}

func maxY(points []point) int {
	maxVal := math.MinInt
	for _, p := range points {
		maxVal = max(maxVal, p.Y)
	}

	return maxVal
}
