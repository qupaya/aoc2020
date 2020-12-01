package main

import (
	"aoc/apps/markus/api/src/day1"
	"fmt"
)

func main() {
	input := day1.Input()
	fmt.Println("day1, task 1:", day1.SolveTask1(input))
	fmt.Println("day1, task 2:", day1.SolveTask2(input))
}
