package main

type Wall struct {
	Points []Point
}

func (w Wall) Clone() Wall {
	points := make([]Point, 0, len(w.Points))
	for _, x := range w.Points {
		points = append(points, Point{
			X: x.X,
			Y: x.Y,
		})
	}

	return Wall{Points: points}
}
