package task1

import (
	"aoc-2021-day20/image"
	"aoc-2021-day20/image/algorithm"
	"aoc-2021-day20/input"
)

func Solve(data input.Data) (int, error) {
	resized, err := image.FromExisting(data.Image, len(data.Image.Pixels)+4, len(data.Image.Pixels[0])+4, 2, 2)
	if err != nil {
		return 0, err
	}

	step1 := enhance(resized, data.Algorithm)

	resizedS1, err := image.FromExisting(step1, len(step1.Pixels)+4, len(step1.Pixels[0])+4, 2, 2)
	if err != nil {
		return 0, err
	}

	step2 := enhance(resizedS1, data.Algorithm)

	return calculateLit(step2), nil
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

func enhance(resized image.Image, algorithm algorithm.Algorithm) image.Image {
	center := position{
		x: 1,
		y: 1,
	}

	newImage := resized.Clone()

	for ; center.x < len(resized.Pixels)-1; center.x++ {
		for ; center.y < len(resized.Pixels[0])-1; center.y++ {
			number := getNumericVal(resized, center)
			algorithmVal := algorithm[number]
			newImage.Pixels[center.x][center.y] = algorithmVal
		}

		center.y = 1
	}

	return newImage
}

func getNumericVal(resized image.Image, center position) int {
	num := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			num = num << 1
			num += int(resized.Pixels[center.x+i][center.y+j])
		}
	}

	return num
}

type position struct {
	x, y int
}
