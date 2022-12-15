package main

import "fmt"

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
