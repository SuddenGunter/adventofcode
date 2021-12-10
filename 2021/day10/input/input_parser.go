package input

import (
	"os"
)

type Data struct {
	Lines []string
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	//lines := strings.Split(string(file), "\n")

	// todo: what if line empty? (last line)

	data := Data{
		Lines: nil,
	}

	//for i, l := range lines {
	//	if l == "" {
	//		continue
	//	}
	//
	//	line, err := parseLine(l)
	//	if err != nil {
	//		return Data{}, fmt.Errorf("failed to parse line %v: %w", i, err)
	//	}
	//
	//	data.Lines = append(data.Lines, line)
	//}

	return data, nil
}
