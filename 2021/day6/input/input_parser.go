package input

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Fishes         []Fish
	DaysToSimulate int
}

type Fish struct {
	DaysBeforeBirth int8
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

		num, err := strconv.Atoi(strings.Trim(char, " \n"))
		if err != nil {
			return Data{}, fmt.Errorf("failed to parse as number '%v': %w", char, err)
		}

		data.Fishes = append(data.Fishes, Fish{DaysBeforeBirth: int8(num)})
	}

	data.DaysToSimulate = 80

	return data, nil
}
