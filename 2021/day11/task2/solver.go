package task2

import "aoc-2021-day11/input"

const steps = 999

type position struct {
	x, y int
}

func Solve(data input.Data) (int, error) {
	for i := 0; i < steps; i++ {
		if simulateStep(data.Octopuses) {
			return i + 1, nil
		}
	}

	return -1, nil
}

func simulateStep(octopuses [][]byte) bool {
	flashed := make(map[position]struct{})

	shouldFlash := &queue{elements: make([]position, 0)}
	for i := range octopuses {
		for j := range octopuses[i] {
			octopuses[i][j]++
			if octopuses[i][j] > 9 {
				shouldFlash.enqueue(position{
					x: i,
					y: j,
				})
			}
		}
	}

	for !shouldFlash.isEmpty() {
		pos := shouldFlash.dequeue()

		_, alreadyFlashed := flashed[pos]
		if alreadyFlashed {
			continue
		}

		flashed[pos] = struct{}{}

		nb := neighbours(pos, len(octopuses), len(octopuses[pos.x]))
		for _, n := range nb {
			octopuses[n.x][n.y]++
			if octopuses[n.x][n.y] > 9 {
				_, flashed := flashed[position{
					x: n.x,
					y: n.y,
				}]

				if !flashed {
					shouldFlash.enqueue(n)
				}
			}
		}
	}

	for k := range flashed {
		octopuses[k.x][k.y] = 0
	}
	if len(flashed) == len(octopuses)*len(octopuses[0]) {
		return true
	}

	return false
}

func neighbours(p position, lenX, lenY int) []position {
	nXY := append(neighboursX(p, lenX), neighboursY(p, lenY)...)

	canAddPrevRow := false
	canAddNextRow := false
	canAddPrevCol := false
	canAddNextCol := false

	for _, v := range nXY {
		if v.x < p.x {
			canAddPrevRow = true
		}
		if v.x > p.x {
			canAddNextRow = true
		}
		if v.y < p.y {
			canAddPrevCol = true
		}
		if v.y > p.y {
			canAddNextCol = true
		}
	}

	if canAddPrevCol && canAddPrevRow {
		nXY = append(nXY, position{
			x: p.x - 1,
			y: p.y - 1,
		})
	}
	if canAddNextCol && canAddNextRow {
		nXY = append(nXY, position{
			x: p.x + 1,
			y: p.y + 1,
		})
	}
	if canAddNextRow && canAddPrevCol {
		nXY = append(nXY, position{
			x: p.x + 1,
			y: p.y - 1,
		})
	}
	if canAddNextCol && canAddPrevRow {
		nXY = append(nXY, position{
			x: p.x - 1,
			y: p.y + 1,
		})
	}

	return nXY
}

func neighboursX(p position, lenX int) []position {
	result := make([]position, 0)
	if p.x > 0 {
		result = append(result, position{
			x: p.x - 1,
			y: p.y,
		})
	}

	if p.x < lenX-1 {
		result = append(result, position{
			x: p.x + 1,
			y: p.y,
		})
	}

	return result
}

func neighboursY(p position, lenY int) []position {
	result := make([]position, 0)
	if p.y > 0 {
		result = append(result, position{
			x: p.x,
			y: p.y - 1,
		})
	}

	if p.y < lenY-1 {
		result = append(result, position{
			x: p.x,
			y: p.y + 1,
		})
	}

	return result
}
