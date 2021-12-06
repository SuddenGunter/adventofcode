package input

import (
	"aoc-2021-day6/input/age"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Fishes []Fish
}

type Fish struct {
	DaysBeforeBirth int8
	Age             age.FishAge
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	chars := strings.Split(string(file), ",")

	data := Data{
		Fishes: make([]Fish, 0, len(chars)),
	}

	for _, char := range chars {
		if char == "" {
			continue
		}

		num, err := strconv.Atoi(char)
		if err != nil {
			return Data{}, fmt.Errorf("failed to parse as number '%v': %w", char, err)
		}

		data.Fishes = append(data.Fishes, Fish{Age: age.Old, DaysBeforeBirth: int8(num)})
	}

	return data, nil
}
