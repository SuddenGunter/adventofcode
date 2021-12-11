package input

import (
	"os"
	"strings"
)

type Data struct {
	OctopusesP1 [][]byte
	OctopusesP2 [][]byte
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	data := Data{
		OctopusesP1: make([][]byte, 0, len(lines)-1),
		OctopusesP2: make([][]byte, 0, len(lines)-1),
	}

	for _, l := range lines[:len(lines)-1] {
		if l == "" {
			continue
		}

		lineP1 := parseLine(l)
		lineP2 := make([]byte, len(lineP1), len(lineP1))
		copy(lineP2, lineP1)

		data.OctopusesP1 = append(data.OctopusesP1, lineP1)
		data.OctopusesP2 = append(data.OctopusesP2, lineP2)
	}

	return data, nil
}

func parseLine(l string) []byte {
	numbers := make([]byte, 0, len(l)-1)
	for _, n := range l {
		numbers = append(numbers, byte(n-'0'))
	}

	return numbers
}
