package task2

import (
	"aoc-2021-day21/game"
	"aoc-2021-day21/input"
)

var totalRolled2Realitites = map[uint64]uint64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func Solve(data input.Data) (uint64, error) {
	var p1wins, p2wins uint64

	for k, v := range totalRolled2Realitites {
		p1w, p2w := Game(data.P1, data.P2, true, k, v)

		p1wins += p1w * v
		p2wins += p2w * v
	}

	return max(p1wins, p2wins), nil
}

func Game(p1, p2 game.Player, p1turn bool, roll, rollMP uint64) (uint64, uint64) {
	var p1wins, p2wins uint64

	if p1turn {
		p1 = Step(p1, roll)
	} else {
		p2 = Step(p2, roll)
	}

	p1turn = !p1turn

	if Won(p1) {
		return 1, 0
	}

	if Won(p2) {
		return 0, 1
	}

	for k, v := range totalRolled2Realitites {
		p1w, p2w := Game(p1, p2, p1turn, k, v)

		p1wins += p1w * v
		p2wins += p2w * v
	}

	return p1wins, p2wins
}

func Won(p game.Player) bool {
	return p.Score >= 21
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	}

	return b
}

func Step(p game.Player, roll uint64) game.Player {
	move := roll

	if move >= 10 {
		move = move % 10
	}

	p.Position += move

	if p.Position > 10 {
		p.Position = p.Position % 10
	}

	p.Score += p.Position

	return p
}
