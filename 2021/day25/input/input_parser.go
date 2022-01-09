package input

import (
	"os"
	"strings"
)

type Data struct {
	Cucumbers [][]rune
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	lines = lines[:len(lines)-1]

	cucumbers := make([][]rune, 0, len(lines))

	for i := range lines {
		row := make([]rune, len(lines[i]))

		for j := range lines[i] {
			row[j] = rune(lines[i][j])
		}

		cucumbers = append(cucumbers, row)
	}

	return Data{
		cucumbers,
	}, nil
}
