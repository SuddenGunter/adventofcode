package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := readInput()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("task 1 answer: %v\n", SolveTask(data.Clone(), 1))
	fmt.Printf("task 2 answer: %v\n", SolveTask(data.Clone(), 2))
}

func readInput() (*Task, error) {
	f, err := os.Open("data.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	walls := make([]Wall, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}

		walls = append(walls, parseWall(s.Text()))
	}

	if err = s.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan file: %w", err)
	}

	return &Task{
		Walls: walls,
		SandSpawn: Point{
			X: 500,
			Y: 0,
		},
	}, nil
}

func parseWall(text string) Wall {
	wall := Wall{}
	points := strings.Split(text, " -> ")
	wall.Points = make([]Point, 0, len(points))

	for _, p := range points {
		wall.Points = append(wall.Points, parsePoint(p))
	}

	return wall
}

func parsePoint(p string) Point {
	coordinate := strings.Split(p, ",")
	x, err := strconv.Atoi(coordinate[0])
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(coordinate[1])
	if err != nil {
		log.Fatal(err)
	}

	return Point{X: x, Y: y}
}
