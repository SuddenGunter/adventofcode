package task2

import (
	"aoc-2021-day21/game"
	"aoc-2021-day21/input"
	"testing"
)

func Benchmark_Solve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, _ := Solve(
			input.Data{
				P1: game.Player{
					Position: 4,
					Score:    0,
				},
				P2: game.Player{
					Position: 8,
					Score:    0,
				},
			})

		if result != uint64(444356092776315) {
			b.Fatal("invalid result")
		}
	}
}

func Benchmark_SolveCached(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, _ := SolveCached(
			input.Data{
				P1: game.Player{
					Position: 4,
					Score:    0,
				},
				P2: game.Player{
					Position: 8,
					Score:    0,
				},
			})

		if result != uint64(444356092776315) {
			b.Fatal("invalid result")
		}
	}
}
