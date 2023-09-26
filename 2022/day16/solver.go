package main

import "fmt"

func SolveTask(data *Data, task int) int {
	switch task {
	case 1:
		fmt.Println(data)
		// 1. Find shortest distance to all [not yet enabled] valves [with flowRate > 0]
		// 2. Calculate "max total throughput valve" -> max from v := (30-distance to valve-1)*flow rate
		// 3. Go to "max total throughput valve", enable it.
		// 4. If going through  [not yet enabled] valve [with flowRate > 0] - enable
		// 5. When selecting route to "max total throughput valve" OR making choise between multiple "max total throughput valve"
		//  if multiple options are available - simulate both, chose best in the end.
		return 0 // solveTask1(data)
	case 2:
		return 0 // solveTask2(data)
	default:
		return -1
	}
}
