package input

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	Fishes []Fish
}

type Fish struct {
	DaysBeforeBirth byte
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

		data.Fishes = append(data.Fishes, Fish{DaysBeforeBirth: byte(num)})
	}

	return data, nil
}
