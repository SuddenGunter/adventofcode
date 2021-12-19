package input

import (
	"aoc-2021-day19/point"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	Scanners []map[point.Point3d]struct{}
}

func ParseInput(name string) (Data, error) {
	file, err := os.ReadFile(name)
	if err != nil {
		return Data{}, err
	}

	lines := strings.Split(string(file), "\n")
	lines = lines[:len(lines)-1]

	scanners := make([]map[point.Point3d]struct{}, 0, len(lines))
	var scanner map[point.Point3d]struct{}
	for _, l := range lines {
		if l == "" {
			continue
		}

		if l[0:2] == "--" {
			if scanner != nil {
				scanners = append(scanners, scanner)
			}

			scanner = make(map[point.Point3d]struct{})
			continue
		}

		p, err := point.FromString(l)
		if err != nil {
			return Data{}, fmt.Errorf("cannot parse string '%s' as point: %w", l, err)
		}

		scanner[p] = struct{}{}
	}

	return Data{
		Scanners: scanners,
	}, nil
}
