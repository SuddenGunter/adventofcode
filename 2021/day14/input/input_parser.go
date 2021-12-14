package input

import (
	"os"
	"strings"
)

type PairInsertionRule struct {
	From string
	To   string
}

type Data struct {
	InitialState string
	Changes      []PairInsertionRule
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")

	data := Data{
		InitialState: lines[0],
		Changes:      make([]PairInsertionRule, 0, len(lines)-2),
	}

	for _, l := range lines[2 : len(lines)-1] {
		split := strings.Split(strings.Trim(l, "\n "), " -> ")
		data.Changes = append(data.Changes, PairInsertionRule{
			From: split[0],
			To:   split[1],
		})
	}

	return data, nil
}
