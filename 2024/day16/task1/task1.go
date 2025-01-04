package task1

import (
	"log"
	"maps"
	"math"
	"slices"

	"aoc-2024-day16/heap"
)

const (
	up    = "up"
	down  = "down"
	left  = "left"
	right = "right"
)

type vertex struct {
	p         point
	neighbors []*vertex
}

type move struct {
	accumulated_path []move
	accumulated_cost int
	vertex           *vertex
	entry_direction  string
}

type vertexWithDir struct {
	vertex          *vertex
	entry_direction string
}

type point struct {
	row, col int
}

func Solve(input []byte) int {
	graph, start, finish := build_graph(input)
	return find_cheapest_path(graph, start, finish)
}

func find_cheapest_path(graph []*vertex, start point, finish point) int {
	h := heap.NewHeap[move]()
	startV := must_find(graph, start)

	visited := map[vertexWithDir]int{}

	h.Insert(heap.Entity[move]{
		Move: move{
			vertex:          startV,
			entry_direction: right,
		},
	})

	return traverse(h, graph, visited, finish)
}

func traverse(h *heap.Heap[move], graph []*vertex, visited map[vertexWithDir]int, finish point) int {
	current_move, err := h.TakeTop()
	if err != nil {
		// todo shandle invalid path
		return math.MaxInt
	}

	visited[vertexWithDir{
		vertex:          current_move.Move.vertex,
		entry_direction: current_move.Move.entry_direction,
	}] = current_move.Move.accumulated_cost
	if current_move.Move.vertex.p == finish {
		// fmt.Println(len(current_move.Move.accumulated_path))
		// for _, x := range current_move.Move.accumulated_path {
		// 	fmt.Println(x.vertex.p)
		// }
		return current_move.Move.accumulated_cost
	}

	for _, x := range current_move.Move.vertex.neighbors {
		cost, new_direction := get_move_cost(current_move.Move.vertex, x, current_move.Move.entry_direction)
		if cost == -1 {
			continue // invalid move: 180 rotations are not allowed
		}

		copyMoves := make([]move, len(current_move.Move.accumulated_path))
		copy(copyMoves, current_move.Move.accumulated_path)

		e := heap.Entity[move]{
			Move: move{
				accumulated_path: append(copyMoves, current_move.Move), // append(current_move.Move.accumulated_path, current_move.Move),
				accumulated_cost: current_move.Move.accumulated_cost + cost,
				entry_direction:  new_direction,
				vertex:           x,
			},
			Priority: current_move.Move.accumulated_cost + cost,
		}

		if already_v_cost, already_visited := visited[vertexWithDir{
			entry_direction: new_direction,
			vertex:          x,
		}]; already_visited && already_v_cost <= current_move.Move.accumulated_cost+cost {
			continue // skip loop
		}

		h.Insert(e)
	}

	return traverse(h, graph, visited, finish)
}

func get_move_cost(from, to *vertex, current_direction string) (int, string) {
	new_dir_point := point{
		row: to.p.row - from.p.row,
		col: to.p.col - from.p.col,
	}

	// assert
	if math.Abs(float64(new_dir_point.row))+math.Abs(float64(new_dir_point.col)) != 1.0 {
		panic("unsupported move")
	}

	dir := as_direction(new_dir_point)
	if dir == opposite(current_direction) {
		return -1, ""
	}
	if dir != current_direction {
		return 1000 + 1, dir
	} else {
		return 0 + 1, dir
	}
}

func opposite(dir string) string {
	switch dir {
	case up:
		return down
	case down:
		return up
	case left:
		return right
	case right:
		return left
	default:
		panic("unknown direction")
	}
}

func as_point(dir string) point {
	switch dir {
	case down:
		return point{
			row: 1,
			col: 0,
		}
	case up:
		return point{
			row: -1,
			col: 0,
		}
	case right:
		return point{
			row: 0,
			col: 1,
		}
	case left:
		return point{
			row: 0,
			col: -1,
		}
	default:
		panic("unexpected direction")
	}
}

func as_direction(p point) string {
	switch p {
	case point{row: 1, col: 0}:
		return down
	case point{row: -1, col: 0}:
		return up
	case point{row: 0, col: 1}:
		return right
	case point{row: 0, col: -1}:
		return left

	default:
		panic("unexpected direction")
	}
}

func must_find(graph []*vertex, start point) *vertex {
	for _, x := range graph {
		if x.p == start {
			return x
		}
	}

	log.Fatalf("%v not found in graph", start)
	return nil
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
