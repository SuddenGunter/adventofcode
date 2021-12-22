package geometry

type Operation struct {
	Cuboid Cuboid
	On     bool
}

type Cuboid struct {
	LowerX, UpperX int64
	LowerY, UpperY int64
	LowerZ, UpperZ int64
}

func (c Cuboid) Volume() int64 {
	vol := float64(c.UpperX-c.LowerX+1) *
		float64(c.UpperY-c.LowerY+1) *
		float64(c.UpperZ-c.LowerZ+1)

	return int64(vol)
}
