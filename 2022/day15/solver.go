package main

import (
	"golang.org/x/exp/constraints"
	"math"
	"sort"
)

const (
	row          = 2000000          //10 for test data
	sizeX, sizeY = 4000000, 4000000 //20, 20 for test data
)

func SolveTask(t *Task, num int) int {
	switch num {
	case 1:
		intervals := getForbiddenIntervals(t)
		intervals = mergeIntersecting(intervals)
		beacons := cutOutBeacons(t, intervals)

		forbiddenPointsCount := countForbiddenPoints(intervals)
		return forbiddenPointsCount - beacons
	case 2:
		coordinates := findBeaconCoordinates(t)
		return coordinates.X*4000000 + coordinates.Y
	default:
		panic("unexpected task")
	}
}

// inspired by https://github.com/jasontconnell/advent/blob/master/2022/15/main.go
func findBeaconCoordinates(t *Task) Point {
	for y := 0; y <= sizeY; y++ {
		for x := 0; x <= sizeX; x++ {
			forbiddenByScanner := false
			for _, p := range t.Pairs {
				distToCurrent := manhattanDistance(p.Sensor, Point{X: x, Y: y})
				// distToBeacon acts like a sensor "strenght"
				distToBeacon := manhattanDistance(p.Sensor, p.Beacon)
				if distToCurrent <= distToBeacon {
					forbiddenByScanner = true
					skip := distToBeacon - distToCurrent
					x += skip // skipping part of current row covered by sensor
					break
				}
			}

			if !forbiddenByScanner {
				return Point{X: x, Y: y}
			}
		}
	}

	panic("solution not found")
}

func cutOutBeacons(t *Task, intervals []Interval) int {
	toRemove := make(map[int]struct{})
	for _, x := range t.Pairs {
		if x.Beacon.Y == row {
			toRemove[x.Beacon.X] = struct{}{}
		}
	}

	beaconsInsideInterval := 0
	for beacon := range toRemove {
		for _, x := range intervals {
			if beacon >= x.From && beacon <= x.To {
				beaconsInsideInterval++
				break
			}
		}

	}

	return beaconsInsideInterval
}

// based on https://www.geeksforgeeks.org/merging-intervals/
func mergeIntersecting(intervals []Interval) []Interval {
	result := make([]Interval, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].From < intervals[j].From
	})
	tmp := newStack[Interval]()
	tmp.push(intervals[0])
	for i := 1; i < len(intervals); i++ {
		top := tmp.peek()
		if top.To < intervals[i].From {
			tmp.push(intervals[i])
			continue
		}

		if top.To < intervals[i].To {
			top.To = intervals[i].To
			tmp.pop()
			tmp.push(top)
		}
	}

	for !tmp.isEmpty() {
		result = append(result, tmp.pop())
	}

	return result
}

func countForbiddenPoints(intervals []Interval) int {
	sum := 0
	for _, x := range intervals {
		sum += diff(x.From, x.To) + 1
	}

	return sum
}

func getForbiddenIntervals(t *Task) []Interval {
	forbidden := make([]Interval, 0)
	for _, p := range t.Pairs {
		sameColumn := p.Sensor
		sameColumn.Y = row

		distToRow := manhattanDistance(p.Sensor, sameColumn)
		distToBeacon := manhattanDistance(p.Sensor, p.Beacon)
		if distToRow > distToBeacon {
			continue
		}
		dist := diff(distToBeacon, distToRow)
		left := sameColumn
		left.X -= dist
		right := sameColumn
		right.X += dist

		forbidden = append(forbidden, Interval{
			From: left.X,
			To:   right.X,
		})
	}

	return forbidden
}

func min[T constraints.Ordered](a, b T) T {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max[T constraints.Ordered](a, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

func manhattanDistance(from Point, to Point) int {
	x := from.X - to.X
	y := from.Y - to.Y

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

type Interval struct {
	From, To int
}

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
