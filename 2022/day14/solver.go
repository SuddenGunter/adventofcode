package main

import (
	"fmt"
	"log"
	"math"

	"golang.org/x/exp/constraints"
)

const (
	sand  = 'o'
	wall  = '#'
	empty = '.'

	notMoved  = "notMoved"
	moving    = "moving"
	finalized = "finalized"
)

type MapState struct {
	Rows                   [][]rune
	maxX, maxY, minX, minY int

	topLeft, bottomRight Point
}

func (m *MapState) At(p Point) rune {
	return m.Rows[p.Y][p.X]
}

func (m *MapState) SetAt(p Point, r rune) {
	m.Rows[p.Y][p.X] = r
}

func (m *MapState) Draw() {
	for _, y := range m.Rows {
		fmt.Println(string(y))
	}
}

func (m *MapState) Clone() *MapState {
	clone := MapState{
		Rows:        nil,
		maxX:        m.maxX,
		maxY:        m.maxY,
		minX:        m.minX,
		minY:        m.minY,
		bottomRight: m.bottomRight,
		topLeft:     m.topLeft,
	}

	rows := make([][]rune, len(m.Rows))
	for i, r := range m.Rows {
		row := make([]rune, len(r))
		for j, x := range r {
			row[j] = x
		}

		rows[i] = row
	}

	clone.Rows = rows

	return &clone
}

func (m *MapState) Map(p Point) Point {
	x := p.X - m.minX
	y := p.Y - m.minY
	return Point{
		X: x,
		Y: y,
	}
}

func (m *MapState) Equal(to *MapState) bool {
	if len(m.Rows) != len(to.Rows) {
		return false
	}

	for i := range m.Rows {
		for j := range m.Rows[i] {
			if len(m.Rows[i]) != len(to.Rows[i]) {
				return false
			}

			if m.Rows[i][j] != to.Rows[i][j] {
				return false
			}
		}
	}

	return true
}

func (m *MapState) Legal(p Point) bool {
	return p.X >= m.topLeft.X && p.Y >= m.topLeft.Y && p.X <= m.bottomRight.X && p.Y <= m.bottomRight.Y
}

func SolveTask(t *Task, num int) int {
	switch num {
	case 1:
		state := generateMap(t)
		buildWalls(&state, t.Walls)
		state = *simulateSand(&state, state.Map(t.SandSpawn))

		return countSand(&state)
	case 2:
		return 0
	default:
		panic("unexpected task")
	}
}

func countSand(m *MapState) int {
	counter := 0
	for _, r := range m.Rows {
		for _, x := range r {
			if x == sand {
				counter++
			}
		}
	}

	return counter
}

func simulateSand(m *MapState, spawn Point) *MapState {
	temp := m.Clone()
	for {
		withSand := spawnSandUnit(temp.Clone(), spawn)
		dropSandUnit(withSand, spawn)

		if withSand.Equal(temp) {
			break
		} else {
			temp = withSand
		}

	}

	return temp
}

func spawnSandUnit(m *MapState, spawn Point) *MapState {
	if m.At(spawn) != empty {
		m.Draw()
		log.Fatal("spawn not empty")
	}

	m.SetAt(spawn, sand)
	return m
}

func dropSandUnit(m *MapState, pos Point) {
	moveTo := pos
	moveTo.Y++

	for {
		res := tryMoveSandUnit(m, pos, moveTo)
		switch res {
		case finalized:
			return
		case moving:
			pos = moveTo
			moveTo.Y++
			continue
		case notMoved:
		}

		moveTo.X--
		res = tryMoveSandUnit(m, pos, moveTo)
		switch res {
		case finalized:
			return
		case moving:
			pos = moveTo
			moveTo.Y++
			continue
		case notMoved:
		}

		moveTo.X += 2
		res = tryMoveSandUnit(m, pos, moveTo)
		switch res {
		case finalized:
			return
		case moving:
			pos = moveTo
			moveTo.Y++
			continue
		case notMoved:
			return
		}

	}
}

func tryMoveSandUnit(m *MapState, from, moveTo Point) string {
	// sand felt below FoV
	if !m.Legal(moveTo) {
		m.SetAt(from, empty)
		return finalized
	}

	if m.At(moveTo) == empty {
		m.SetAt(moveTo, sand)
		m.SetAt(from, empty)
		return moving
	}

	return notMoved
}

func buildWalls(state *MapState, walls []Wall) {
	for _, w := range walls {
		for i := 0; i < len(w.Points); i++ {
			point := w.Points[i]
			next := nextPoint(w.Points, i+1)
			drawLine(state.Map(point), state.Map(next), state)
		}
	}
}

func drawLine(point Point, next Point, state *MapState) {
	fromY := min(point.Y, next.Y)
	toY := max(point.Y, next.Y)
	fromX := min(point.X, next.X)
	toX := max(point.X, next.X)

	for y := fromY; y <= toY; y++ {
		for x := fromX; x <= toX; x++ {
			state.Rows[y][x] = wall
		}
	}
}

func nextPoint(points []Point, i int) Point {
	if len(points) <= i {
		return points[i-1]
	} else {
		return points[i]
	}
}

func generateMap(t *Task) MapState {
	minX, maxX, minY, maxY := getEdges(t.Walls)
	state := MapState{
		Rows: nil,
		maxX: maxX + 1,
		maxY: maxY + 1,
		minX: minX - 1,
		minY: minY,
	}

	state.topLeft = state.Map(Point{
		X: minX,
		Y: minY,
	})
	state.bottomRight = state.Map(Point{
		X: maxX,
		Y: maxY,
	})

	rows := make([][]rune, state.maxY-state.minY+1)
	for y := state.minY; y <= state.maxY; y++ {
		row := make([]rune, state.maxX-state.minX+1)
		for x := state.minX; x <= state.maxX; x++ {
			row[state.Map(Point{
				X: x,
			}).X] = empty
		}

		rows[state.Map(Point{
			Y: y,
		}).Y] = row
	}

	state.Rows = rows

	return state
}

func getEdges(walls []Wall) (int, int, int, int) {
	minX, minY := math.MaxInt, 0
	maxX, maxY := math.MinInt, math.MinInt
	for _, w := range walls {
		for _, p := range w.Points {
			minX = min(minX, p.X)
			maxX = max(maxX, p.X)
			maxY = max(maxY, p.Y)
		}
	}

	return minX, maxX, minY, maxY
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
