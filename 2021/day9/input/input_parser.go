package input

import (
	"os"
	"strings"
)

type Data struct {
	Numbers [][]byte
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")
	data := Data{
		Numbers: make([][]byte, 0, len(lines)-1),
	}

	for _, l := range lines[:len(lines)-1] {
		line := parseLine(l)

		data.Numbers = append(data.Numbers, line)
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
