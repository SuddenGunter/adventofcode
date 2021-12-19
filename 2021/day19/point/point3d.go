package point

import (
	"strconv"
	"strings"
)

type Point3d struct {
	X, Y, Z int
}

func FromString(input string) (Point3d, error) {
	split := strings.Split(strings.Trim(input, "\n "), ",")

	x, err := strconv.Atoi(split[0])
	if err != nil {
		return Point3d{}, err
	}

	y, err := strconv.Atoi(split[1])
	if err != nil {
		return Point3d{}, err
	}

	z, err := strconv.Atoi(split[2])
	if err != nil {
		return Point3d{}, err
	}

	return Point3d{
		X: x,
		Y: y,
		Z: z,
	}, nil
}
