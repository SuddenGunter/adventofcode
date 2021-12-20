package task1

import (
	"aoc-2021-day20/image"
	"aoc-2021-day20/image/algorithm"
	"aoc-2021-day20/input"
	"fmt"
)

func Solve(data input.Data) (int, error) {
	resized, err := image.FromExisting(data.Image, len(data.Image.Pixels)+4, len(data.Image.Pixels[0])+4, 2, 2)
	if err != nil {
		return 0, err
	}

	fmt.Println("original image")
	prettyPrint(resized)
	fmt.Println()

	step0 := enhance(resized, data.Algorithm, 0)

	fmt.Println("step0")
	prettyPrint(step0)
	fmt.Println()

	resizedS1, err := image.FromExisting(step0, len(step0.Pixels)+4, len(step0.Pixels[0])+4, 2, 2)
	if err != nil {
		return 0, err
	}

	step1 := enhance(resizedS1, data.Algorithm, 1)

	fmt.Println("step1")
	prettyPrint(step1)
	fmt.Println()

	return calculateLit(step1), nil
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

func enhance(resized image.Image, algorithm algorithm.Algorithm, step int) image.Image {
	center := position{
		x: 0,
		y: 0,
	}

	newImage := resized.Clone()

	for ; center.x < len(resized.Pixels); center.x++ {
		for ; center.y < len(resized.Pixels[0]); center.y++ {
			number := getNumericVal(resized, algorithm, center, step)
			algorithmVal := algorithm[number]
			newImage.Pixels[center.x][center.y] = algorithmVal
		}

		center.y = 1
	}

	return newImage
}

func getNumericVal(resized image.Image, alg algorithm.Algorithm, center position, step int) int {
	num := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			num = num << 1
			num += getVal(resized, alg, center.x+i, center.y+j, step)
		}
	}

	return num
}

func getVal(resized image.Image, alg algorithm.Algorithm, i int, j int, step int) int {
	if i < 0 || i >= len(resized.Pixels) || j < 0 || j >= len(resized.Pixels[0]) {
		switch {
		case step == 0:
			return 0
		case step%2 == 0:
			return int(alg[0])
		default:
			if alg[0] == 0 {
				return 0
			} else {
				return int(alg[511])
			}
		}
	} else {
		return int(resized.Pixels[i][j])
	}
}

type position struct {
	x, y int
}
