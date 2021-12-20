package task1

import (
	"aoc-2021-day20/image"
	"aoc-2021-day20/image/algorithm"
	"aoc-2021-day20/input"
	"fmt"
)

const (
	stepSize = 2
	simSize  = 4
)

func Solve(data input.Data) (int, error) {
	outOfBounds := byte(0)

	stepResult := data.Image.Clone()
	for i := 0; i < stepSize; i++ {
		resized, err := image.FromExisting(stepResult, len(stepResult.Pixels)+simSize, len(stepResult.Pixels[0])+simSize, simSize/2, simSize/2, outOfBounds)
		if err != nil {
			return 0, err
		}

		stepResult = enhance(resized, data.Algorithm, outOfBounds)
		if outOfBounds == 0 {
			outOfBounds = data.Algorithm[0]
		} else {
			outOfBounds = data.Algorithm[511]
		}
	}

	return calculateLit(stepResult), nil
}

func prettyPrint(img image.Image) {
	for i := range img.Pixels {
		for _, v := range img.Pixels[i] {
			if v == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}

		fmt.Println()
	}
}

func calculateLit(img image.Image) int {
	sum := 0
	for _, line := range img.Pixels {
		for _, v := range line {
			sum += int(v)
		}
	}

	return sum
}

func enhance(resized image.Image, algorithm algorithm.Algorithm, outOfBounds byte) image.Image {
	center := position{
		x: 0,
		y: 0,
	}

	newImage := resized.Clone()

	for ; center.x < len(resized.Pixels); center.x++ {

		for ; center.y < len(resized.Pixels[0]); center.y++ {
			number := getNumericVal(resized, algorithm, center, outOfBounds)
			algorithmVal := algorithm[number]
			newImage.Pixels[center.x][center.y] = algorithmVal
		}

		center.y = 0
	}

	return newImage
}

func getNumericVal(resized image.Image, alg algorithm.Algorithm, center position, outOfBounds byte) int {
	num := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			num = num << 1
			num += getVal(resized, alg, center.x+i, center.y+j, outOfBounds)
		}
	}

	return num
}

func getVal(resized image.Image, alg algorithm.Algorithm, i, j int, outOfBounds byte) int {
	if i < 0 || i >= len(resized.Pixels) || j < 0 || j >= len(resized.Pixels[0]) {
		return int(outOfBounds)
	} else {
		return int(resized.Pixels[i][j])
	}
}

type position struct {
	x, y int
}
