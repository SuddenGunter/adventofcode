package image

import "errors"

type Image struct {
	Pixels [][]byte
}

func (img Image) Clone() Image {
	pixels := make([][]byte, 0, len(img.Pixels))
	for i := range img.Pixels {
		line := make([]byte, len(img.Pixels[0]))
		for j := range img.Pixels[i] {
			line[j] = img.Pixels[i][j]
		}

		pixels = append(pixels, line)
	}

	return Image{Pixels: pixels}
}

// FromExisting creates new image with sizeX x sizeY and places top left corner of the old
// image to the [posX;posY].
func FromExisting(img Image, sizeX, sizeY, posX, posY int, emptyPixelVal byte) (Image, error) {
	if sizeX < len(img.Pixels)+posX {
		return Image{}, errors.New("can't fit big image to small canvas")
	}

	pixels := make([][]byte, 0, sizeX)
	for i := 0; i < sizeY; i++ {
		line := make([]byte, sizeY)
		for i := range line {
			line[i] = emptyPixelVal
		}
		pixels = append(pixels, line)
	}

	for i := range img.Pixels {

		if sizeY < len(img.Pixels[i])+posY {
			return Image{}, errors.New("can't fit big image to small canvas")
		}

		for j := range img.Pixels[i] {
			pixels[i+posX][j+posY] = img.Pixels[i][j]
		}
	}

	return Image{Pixels: pixels}, nil
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
