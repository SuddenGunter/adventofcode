package point

type Transform struct {
	Scanner Point3d
	Beacons map[Point3d]struct{}
}
