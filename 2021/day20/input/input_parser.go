package input

import (
	"aoc-2021-day20/image"
	"aoc-2021-day20/image/algorithm"
	"os"
	"strings"
)

type Data struct {
	Image     image.Image
	Algorithm algorithm.Algorithm
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	lines = lines[:len(lines)-1]

	alg, err := algorithm.FromString(lines[0])
	if err != nil {
		return Data{}, err
	}

	img, err := image.FromLines(lines[2:])
	if err != nil {
		return Data{}, err
	}

	return Data{
		Image:     img,
		Algorithm: alg,
	}, nil
}
