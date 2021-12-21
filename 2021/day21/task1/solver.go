package task1

import (
	"aoc-2021-day21/game"
	"aoc-2021-day21/input"
	"aoc-2021-day21/rng"
)

var rand, rolls = rng.Deterministic()

func Solve(data input.Data) (uint64, error) {
	p1Turn := true

	for {
		if p1Turn {
			Step(&data.P1)

			if Won(data.P1) {
				return rolls() * data.P2.Score, nil
			}
		} else {
			Step(&data.P2)

			if Won(data.P2) {
				return rolls() * data.P1.Score, nil
			}
		}

		p1Turn = !p1Turn
	}
}

func Step(p *game.Player) {
	r1 := rand()
	r2 := rand()
	r3 := rand()
	move := (r1 + r2 + r3)

	if move >= 10 {
		move = move % 10
	}

	p.Position += move

	if p.Position > 10 {
		p.Position = p.Position % 10
	}

	p.Score += p.Position
}

func Won(p game.Player) bool {
	return p.Score >= 1000
}
