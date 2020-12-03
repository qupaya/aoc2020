package day3

import (
	"fmt"
)

func countTrees(input *[]string, down int, right int) int {
	width := len((*input)[0])
	height := len(*input)
	x := 0
	y := 0
	countTrees := 0

	for y < height {
		if (*input)[y][x%width] == 35 {
			countTrees++
		}
		x += right
		y += down
	}
	return countTrees
}

// SolveTask1 of day 3
func SolveTask1(input *[]string) int {
	return countTrees(input, 1, 3)
}

// SolveTask2 of day 3
func SolveTask2(input *[]string) int {
	slopes := [][]int{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}

	trees := 1
	for _, s := range slopes {
		trees *= countTrees(input, s[0], s[1])
	}
	return trees
}

// Run day3
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(&input))
	fmt.Println("solution task 2:", SolveTask2(&input))

}
