package point

import (
	"errors"
	"strconv"
	"strings"
)

type Point3d struct {
	X, Y, Z int
}

func (p Point3d) Plus(p2 Point3d) Point3d {
	return Point3d{
		X: p.X + p2.X,
		Y: p.Y + p2.Y,
		Z: p.Z + p2.Z,
	}
}

func (p Point3d) Minus(p2 Point3d) Point3d {
	return Point3d{
		X: p.X - p2.X,
		Y: p.Y - p2.Y,
		Z: p.Z - p2.Z,
	}
}

func (p Point3d) Face(facing int) (Point3d, error) {
	var newPoint Point3d
	switch facing {
	case 0:
		newPoint = p
	case 1:
		newPoint = Point3d{
			X: p.X,
			Y: -p.Y,
			Z: -p.Z,
		}
	case 2:
		newPoint = Point3d{
			X: p.X,
			Y: -p.Z,
			Z: p.Y,
		}
	case 3:
		newPoint = Point3d{
			X: -p.Y,
			Y: -p.Z,
			Z: p.X,
		}
	case 4:
		newPoint = Point3d{
			X: p.Y,
			Y: -p.Z,
			Z: -p.X,
		}
	case 5:
		newPoint = Point3d{
			X: -p.X,
			Y: -p.Z,
			Z: -p.Y,
		}

	default:
		return Point3d{}, errors.New("unsupported facing")
	}

	return newPoint, nil
}

func (p Point3d) Rotate(rotating int) (Point3d, error) {
	var newPoint Point3d
	switch rotating {
	case 0:
		newPoint = p
	case 1:
		newPoint = Point3d{
			X: -p.Y,
			Y: p.X,
			Z: p.Z,
		}
	case 2:
		newPoint = Point3d{
			X: -p.X,
			Y: -p.Y,
			Z: p.Z,
		}
	case 3:
		newPoint = Point3d{
			X: p.Y,
			Y: -p.X,
			Z: p.Z,
		}

	default:
		return Point3d{}, errors.New("unsupported rotation")
	}

	return newPoint, nil
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
