package task1

import (
	"aoc-2021-day19/input"
	"aoc-2021-day19/point"
)

func Solve(data input.Data) (int, error) {
	baseSector := data.Scanners[0]

	foundScanners := make(map[point.Point3d]struct{})
	foundScanners[point.Point3d{}] = struct{}{}

	unmappedSectors := queue{elements: make([]map[point.Point3d]struct{}, 0)}

	for _, sector := range data.Scanners[1:] {
		unmappedSectors.enqueue(sector)
	}

	for !unmappedSectors.isEmpty() {
		sector := unmappedSectors.dequeue()
		result, intersects := findTransformIfIntersects(baseSector, sector)
		if intersects {
			for beacon := range result.Beacons {
				baseSector[beacon] = struct{}{}
			}

			foundScanners[result.Scanner] = struct{}{}
		} else {
			// try again later
			unmappedSectors.enqueue(sector)
		}
	}

	return len(baseSector), nil
}

func findTransformIfIntersects(left, right map[point.Point3d]struct{}) (point.Transform, bool) {
	for face := 0; face <= 5; face++ {
		for rotation := 0; rotation <= 3; rotation++ {
			rightReoriented, err := orient(right, face, rotation)
			if err != nil {
				panic(err)
			}

			for l := range left {
				for r := range rightReoriented {
					diff := l.Minus(r)

					moved := move(rightReoriented, diff)

					if intersectCount(moved, left) >= 12 {
						return point.Transform{
							Scanner: diff,
							Beacons: moved,
						}, true
					}
				}
			}

		}
	}

	return point.Transform{}, false
}

func intersectCount(moved, left map[point.Point3d]struct{}) int {
	intersectCounter := 0
	for k := range moved {
		if _, found := left[k]; found {
			intersectCounter++
		}
	}

	return intersectCounter
}

func move(nodes map[point.Point3d]struct{}, diff point.Point3d) map[point.Point3d]struct{} {
	result := make(map[point.Point3d]struct{})

	for k := range nodes {
		result[k.Plus(diff)] = struct{}{}
	}

	return result
}

func orient(nodes map[point.Point3d]struct{}, face, rotation int) (map[point.Point3d]struct{}, error) {
	result := make(map[point.Point3d]struct{})

	for k := range nodes {
		faced, err := k.Face(face)
		if err != nil {
			return nil, err
		}

		rotated, err := faced.Rotate(rotation)
		if err != nil {
			return nil, err
		}

		result[rotated] = struct{}{}
	}

	return result, nil
}
