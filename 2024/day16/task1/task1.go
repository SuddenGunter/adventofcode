package task1

import (
	"fmt"
	"maps"
	"slices"
)

type vertex struct {
	p         point
	neighbors []*vertex
}

type point struct {
	row, col int
}

func Solve(input []byte) int {
	graph, start, finish := build_graph(input)
	fmt.Println(graph)
	fmt.Println(start)
	fmt.Println(finish)
	return 0
}

func build_graph(input []byte) ([]*vertex, point, point) {
	tiles := make(map[point]byte)
	p := point{}
	for _, x := range input {
		switch x {
		case '\n':
			p.row += 1
			p.col = 0
			continue
		case '#':
			p.col += 1
			continue
		default:
			tiles[p] = x
			p.col += 1
		}
	}

	vertexes := make(map[point]*vertex)
	mr, mc := 0, 0
	for p := range maps.Keys(tiles) {
		if p.row > mr {
			mr = p.row
		}
		if p.col > mc {
			mc = p.col
		}
	}
	for pos := range tiles {
		vertexes[pos] = &vertex{
			p:         pos,
			neighbors: make([]*vertex, 0),
		}

		pn := possible_neighbors(pos, mr, mc)
		for _, x := range pn {
			if n, ok := vertexes[x]; ok {
				vertexes[pos].neighbors = append(vertexes[pos].neighbors, n)
				vertexes[x].neighbors = append(vertexes[x].neighbors, vertexes[pos])
			}
		}
	}

	startPos, finPos := point{}, point{}
	for pos, val := range tiles {
		switch val {
		case 'E':
			finPos = pos
		case 'S':
			startPos = pos
		}
	}

	return slices.Collect(maps.Values(vertexes)), startPos, finPos
}

func possible_neighbors(pos point, maxRows int, maxCols int) []point {
	points := make([]point, 0, 4)
	if pos.row-1 >= 0 {
		points = append(points, point{row: pos.row - 1, col: pos.col})
	}
	if pos.row+1 <= maxRows {
		points = append(points, point{row: pos.row + 1, col: pos.col})
	}
	if pos.col-1 >= 0 {
		points = append(points, point{row: pos.row, col: pos.col - 1})
	}
	if pos.col+1 <= maxCols {
		points = append(points, point{row: pos.row, col: pos.col + 1})
	}

	return points
}
