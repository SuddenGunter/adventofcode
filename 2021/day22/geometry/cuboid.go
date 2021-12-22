package geometry

import "math"

type Cuboid struct {
	X1, X2 int
	Y1, Y2 int
	Z1, Z2 int
}

func (c Cuboid) Volume() int {
	vol := math.Abs(float64(c.X1-c.X2)) *
		math.Abs(float64(c.Y1-c.Y2)) *
		math.Abs(float64(c.Z1-c.Z2))

	return int(vol)
}
