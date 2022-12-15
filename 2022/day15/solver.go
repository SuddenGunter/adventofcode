package main

import (
	"golang.org/x/exp/constraints"
	"math"
	"sort"
)

const row = 10 // 2000000

func SolveTask(t *Task, num int) int {
	switch num {
	case 1:
		intervals := getForbiddenIntervals(t)
		intervals = mergeIntersecting(intervals)
		intervals = cutOutBeacons(t, intervals)

		forbiddenPointsCount := countForbiddenPoints(intervals)
		return forbiddenPointsCount
	case 2:
		return 0
	default:
		panic("unexpected task")
	}
}

func cutOutBeacons(t *Task, intervals []Interval) []Interval {
	beaconsOnRow := newStack[Point]()
	for _, x := range t.Pairs {
		if x.Beacon.Y == row {
			beaconsOnRow.push(x.Beacon)
		}
	}

	for !beaconsOnRow.isEmpty() {
		beacon := beaconsOnRow.pop()
		splitInterval := -1
		for i, x := range intervals {
			if beacon.X >= x.From.X && beacon.X <= x.To.X {
				splitInterval = i
				break
			}
		}

		if splitInterval != -1 {
			newIntervals := make([]Interval, 0, len(intervals)+1)
			intervalsToAdd := cutInterval(intervals[splitInterval], beacon)
			if splitInterval > 0 {
				newIntervals = append(newIntervals, intervals[0:splitInterval-1]...)
			}
			if len(intervalsToAdd) > 0 {
				newIntervals = append(newIntervals, intervalsToAdd...)
			}
			if splitInterval < len(intervals)-1 {
				newIntervals = append(newIntervals, intervals[splitInterval+1:]...)
			}

			intervals = newIntervals
		}
	}

	return intervals
}

func cutInterval(interval Interval, beacon Point) []Interval {
	leftX := diff(interval.From.X, beacon.X)
	intervals := make([]Interval, 0)
	if leftX > 0 {
		intervals = append(intervals, Interval{
			From: Point{
				X: interval.From.X,
				Y: interval.From.Y,
			},
			To: Point{
				X: beacon.X - 1,
				Y: beacon.Y,
			},
		})
	}

	rightX := diff(interval.To.X, beacon.X)
	if rightX > 0 {
		intervals = append(intervals, Interval{
			From: Point{
				X: beacon.X + 1,
				Y: beacon.Y,
			},
			To: Point{
				X: interval.To.X,
				Y: interval.To.Y,
			},
		})
	}

	return intervals
}

// based on https://www.geeksforgeeks.org/merging-intervals/
func mergeIntersecting(intervals []Interval) []Interval {
	result := make([]Interval, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].From.X < intervals[j].From.X
	})
	tmp := newStack[Interval]()
	tmp.push(intervals[0])
	for i := 1; i < len(intervals); i++ {
		top := tmp.peek()
		if top.To.X < intervals[i].From.X {
			tmp.push(intervals[i])
			continue
		}

		if top.To.X < intervals[i].To.X {
			top.To.X = intervals[i].To.X
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
		sum += manhattanDistance(x.From, x.To) + 1
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
		dist := diff(distToBeacon, distToRow)
		left := sameColumn
		left.X -= dist
		right := sameColumn
		right.X += dist

		forbidden = append(forbidden, Interval{
			From: left,
			To:   right,
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
	From, To Point
}

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
