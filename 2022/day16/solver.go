package main

import "fmt"

func SolveTask(data *Data, task int) int {
	switch task {
	case 1:
		fmt.Println(data)
		return 0 // solveTask1(data)
	case 2:
		return 0 // solveTask2(data)
	default:
		return -1
	}
}
