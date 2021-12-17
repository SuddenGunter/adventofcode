package task2

import (
	"aoc-2021-day17/input"
	"errors"
	"math"
)

type velocity struct {
	x, y int
}

func Solve(data input.Data) (int, int, error) {
	err := validateInvariants(data)
	if err != nil {
		return 0, 0, err
	}

	xlow, err := getLowerXBound(data)
	if err != nil {
		return 0, 0, err
	}

	xhigh := getUpperXBound(data)
	ylow := getLowerYBound(data)
	yhigh := getUpperYBound(data)

	simulationsCounter := 0
	validStartingVelocity := make(map[velocity]struct{})

	for x := xlow; x <= xhigh; x++ {
		for y := ylow; y <= yhigh; y++ {

			v := velocity{
				x: x,
				y: y,
			}
			if reachedTargetArea(v, data) {
				validStartingVelocity[v] = struct{}{}
			}

			simulationsCounter++
		}
	}

	return len(validStartingVelocity), simulationsCounter, nil
}

func reachedTargetArea(v velocity, data input.Data) bool {
	currentX := 0
	currentY := 0

	// while not missed area
	for currentX <= data.XHigh && currentY >= data.YLow {
		if currentX >= data.XLow && currentX <= data.XHigh &&
			currentY >= data.YLow && currentY <= data.YHigh {
			return true
		}

		currentX += v.x
		currentY += v.y
		v.y--

		if v.x != 0 {
			v.x -= int(math.Copysign(1, float64(v.x)))
		}
	}

	return false
}

func validateInvariants(data input.Data) error {
	if data.XHigh < 0 || data.XLow < 0 {
		return errors.New("not tested on negative Y values, probably will not work because of some assumptions")
	}

	if data.YLow > 0 || data.YHigh > 0 {
		return errors.New("not tested on positive Y values, probably will not work because of some assumptions")
	}

	return nil
}

// getLowerXBound returns min velocity by X. With any velocity below it - we will never reach target area.
// Would need some tinkering for XLow equal to zero or below it.
// Basically we are just solving quadratic equation starting with the nearest target area column. If no whole positive
// number found as one of the roots - it means we cannot reach this column with any velocity - arithmetic progression
// of our movement will never allow it. Think of it if we start moving from that column:
// could we reach our x=0 starting point?
// Even better: we only need *positive* root, so just forget about calculating a negative one.
// Inspired by: https://math.stackexchange.com/questions/238047/reversing-an-arithmetic-sequence.
func getLowerXBound(data input.Data) (int, error) {
	a := 1.
	b := 1.
	minReachableX := data.XLow

	for minReachableX <= data.XHigh {
		c := -float64(minReachableX) * 2
		posSolution := (-b + math.Sqrt(b*b-4*a*c)) / 2
		if isInt(posSolution) {
			return int(posSolution), nil
		}

		minReachableX++
	}

	return 0, errors.New("unsolvable solution: target area unreachable")
}

// getUpperXBound returns max velocity by X. With any velocity above it - we will miss the target area.
// would need some tinkering for XHigh equal to zero or below it.
func getUpperXBound(data input.Data) int {
	return data.XHigh
}

// getLowerYBound returns min velocity by Y which allows reaching target area.
// With any velocity below this value we swim under the area.
func getLowerYBound(data input.Data) int {
	return data.YLow
}

// getUpperYBound returns max velocity by Y which allows reaching target area.
// With any velocity above it - negative arithmetic progression of our velocity leads to jumping over the area.
// Basically this is the same logic as part1: there would be a moment when we are going down, with big velocity and
// reach y=0 (position, not velocity). We are guaranteed to get there - we started there, reached great height - and
// went down with the same progression. So, our last step after it - would be even bigger. Any starting velocity
// bigger than distance_between(startPoint, targetArea.lowerY) - leads to us missing the area.
func getUpperYBound(data input.Data) int {
	return int(math.Abs(float64(data.YLow)) - 1)
}

func isInt(val float64) bool {
	return val == float64(int(val))
}
