package day3

import (
	"fmt"
)

func countTrees(input *[]string, down int, right int) int {
	width := len((*input)[0])
	countTrees := 0
	x := 0
	for y, row := range *input {
		if y%down == 0 {
			if row[x%width] == 35 {
				countTrees++
			}
			x += right
		}
	}
	return countTrees
}

// SolveTask1 of day 3
func SolveTask1(input *[]string) int {
	return countTrees(input, 1, 3)
}

// SolveTask2 of day 3
func SolveTask2(input *[]string) int {
	return countTrees(input, 1, 1) * countTrees(input, 1, 3) * countTrees(input, 1, 5) * countTrees(input, 1, 7) * countTrees(input, 2, 1)
}

// Run day3
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(&input))
	fmt.Println("solution task 2:", SolveTask2(&input))

}
