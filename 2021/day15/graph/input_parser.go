package graph

import (
	"os"
	"strings"
)

type Data struct {
	Lines []string
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	data := Data{
		Lines: lines[:len(lines)-1],
	}

	return data, nil
}
