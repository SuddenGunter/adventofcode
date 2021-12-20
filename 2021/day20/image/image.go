package image

import "errors"

type Image struct {
	Pixels [][]byte
}

func FromLines(lines []string) (Image, error) {
	pixels := make([][]byte, 0, len(lines))

	for _, l := range lines {
		line := make([]byte, 0, len(l))

		for _, char := range []rune(l) {
			switch char {
			case '#':
				line = append(line, 1)
			case '.':
				line = append(line, 0)
			default:
				return Image{}, errors.New("unknown pixel type")
			}
		}

		pixels = append(pixels, line)
	}

	return Image{Pixels: pixels}, nil
}
