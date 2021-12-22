package input

import (
	"aoc-2021-day22/geometry"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Cuboids []geometry.Cuboid
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	lines = lines[:len(lines)-1]

	cuboids := make([]geometry.Cuboid, 0, len(lines))
	for _, l := range lines {
		c, err := parseLine(l)
		if err != nil {
			return Data{}, err
		}

		cuboids = append(cuboids, c)
	}

	return Data{Cuboids: cuboids}, nil
}

func parseLine(l string) (geometry.Cuboid, error) {
	split := strings.Split(l, ",")
	x := split[0]
	y := split[1]
	z := split[2]

	cuboid := geometry.Cuboid{}
	if strings.HasPrefix(x, "on") {
		cuboid.On = true
		x = strings.TrimPrefix(x, "on ")
	} else {
		x = strings.TrimPrefix(x, "off ")
	}

	x1, x2, err := parseNumbers(x, 'x')
	if err != nil {
		return geometry.Cuboid{}, err
	}

	y1, y2, err := parseNumbers(y, 'y')
	if err != nil {
		return geometry.Cuboid{}, err
	}

	z1, z2, err := parseNumbers(z, 'z')
	if err != nil {
		return geometry.Cuboid{}, err
	}

	cuboid.X1 = x1
	cuboid.X2 = x2
	cuboid.Y1 = y1
	cuboid.Y2 = y2
	cuboid.Z1 = z1
	cuboid.Z2 = z2

	return cuboid, nil
}

func parseNumbers(numbers string, letter rune) (int, int, error) {
	numbers = strings.Trim(numbers, fmt.Sprintf(" %v=", letter))
	split := strings.Split(numbers, "..")

	first, err := strconv.Atoi(split[0])
	if err != nil {
		return 0, 0, err
	}

	second, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, err
	}

	first, second = minFirst(first, second)
	return first, second, nil
}

func minFirst(a, b int) (int, int) {
	if a < b {
		return a, b
	}

	return b, a
}
