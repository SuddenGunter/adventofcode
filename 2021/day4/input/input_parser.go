package input

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const Size = 5

type Data struct {
	WinningSequence []int
	Players         []Player
}

type Player struct {
	Index map[int]Position
	Data  [][]int
}

type Position struct {
	X int
	Y int
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	winningSeq, err := parseIntLine(lines[0], ',')
	if err != nil {
		return Data{}, err
	}

	data := Data{
		WinningSequence: winningSeq,
		Players:         make([]Player, 0),
	}

	lines = lines[2:]
	for i := 0; i < len(lines); i += Size + 1 {
		player := Player{
			Index: nil,
			Data:  make([][]int, 0),
		}

		for j := 0; j < Size; j++ {
			l, err := parseIntLine(lines[i+j], ' ')
			if err != nil {
				return Data{}, err
			}

			player.Data = append(player.Data, l)
		}

		player.Index, err = buildIndex(player.Data)
		if err != nil {
			return Data{}, err
		}

		data.Players = append(data.Players, player)
	}

	return data, nil
}

func buildIndex(data [][]int) (map[int]Position, error) {
	index := make(map[int]Position)

	for i := range data {
		for j := range data[i] {
			_, found := index[data[i][j]]
			if found {
				return nil, errors.New("duplicate found, unable to build index")
			}

			index[data[i][j]] = Position{X: i, Y: j}
		}
	}

	return index, nil
}

func parseIntLine(s string, sep rune) ([]int, error) {
	f := func(c rune) bool {
		return c == sep
	}

	// strings.Split would not be able to skip empty entries
	seq := strings.FieldsFunc(s, f)

	seqInt := make([]int, len(seq))

	for i := range seq {
		val, err := strconv.Atoi(seq[i])
		if err != nil {
			return nil, err
		}

		seqInt[i] = val
	}

	return seqInt, nil
}
