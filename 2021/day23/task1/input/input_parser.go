package input

import (
	"aoc-2021-day23/task1/amphipod"
	"errors"
	"os"
	"strings"
)

func ParseInput(name string) (amphipod.Burrow, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return amphipod.Burrow{}, err
	}

	lines := strings.Split(string(file), "\n")

	lines = lines[1 : len(lines)-2]
	hall, err := parseHall(lines)
	if err != nil {
		return amphipod.Burrow{}, err
	}

	rooms, err := parseRooms(lines)
	if err != nil {
		return amphipod.Burrow{}, err
	}

	return amphipod.Burrow{
		Hall:  hall,
		Rooms: rooms,
	}, nil
}

func parseRooms(lines []string) ([amphipod.Rooms][amphipod.RoomSize]rune, error) {
	var rooms [amphipod.Rooms][amphipod.RoomSize]rune

	line := 1
	for i := 0; i < amphipod.RoomSize; i++ {
		j := 0

		for _, r := range []rune(lines[line+i]) {
			switch r {
			case 'A', 'B', 'C', 'D', '.':
				rooms[j][i] = r
				j++
			default:
				continue
			}
		}

		if j != amphipod.Rooms {
			return rooms, errors.New("wrong amount of rooms found")
		}

	}

	return rooms, nil
}

func parseHall(lines []string) ([amphipod.HallSize]rune, error) {
	var hall [amphipod.HallSize]rune

	i := 0
	for _, r := range []rune(lines[0]) {
		switch r {
		case 'A', 'B', 'C', 'D', '.':
			hall[i] = r
			i++
		default:
			continue
		}
	}

	if i != amphipod.HallSize {
		return [amphipod.HallSize]rune{}, errors.New("wrong hallway length")
	}

	return hall, nil
}
