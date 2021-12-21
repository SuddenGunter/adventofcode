package input

import (
	"aoc-2021-day21/game"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Players [2]game.Player
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	lines = lines[:len(lines)-1]

	p1, err := parsePlayer(lines[0])
	if err != nil {
		return Data{}, err
	}

	p1.ID = 1

	p2, err := parsePlayer(lines[1])
	if err != nil {
		return Data{}, err
	}

	p2.ID = 2

	return Data{Players: [2]game.Player{
		p1, p2,
	}}, nil
}

func parsePlayer(s string) (game.Player, error) {
	runes := []rune(s)

	posStr := runes[len(runes)-2:]

	posInt, err := strconv.Atoi(strings.Trim(string(posStr), "\n "))
	if err != nil {
		return game.Player{}, err
	}

	return game.Player{
		Position: uint8(posInt),
		Score:    0,
	}, err
}
