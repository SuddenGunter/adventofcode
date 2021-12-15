package main

import (
	"aoc-2021-day15/graph"
	"aoc-2021-day15/task2"
	"fmt"
)

func main() {
	//dataT1, err := graph.ParseInputForTask1("data.txt")
	//if err != nil {
	//	panic(err)
	//}
	//
	//t1, err := task1.Solve(dataT1)
	//fmt.Printf("task 1 answer: %v, err: %v\n", t1, err)

	dataT2, err := graph.ParseInputForTask2("demo.data.txt")
	if err != nil {
		panic(err)
	}

	t2, err := task2.Solve(dataT2)
	fmt.Printf("task 2 answer: %v, err: %v\n", t2, err)
}
