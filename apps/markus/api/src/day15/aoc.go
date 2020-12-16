package day15

import (
	"fmt"
)

func solveBruteForce(input *[]int, max int) int {
	numberTracker := map[int][]int{}
	numbers := make([]int, max)
	copy(numbers, *input)
	for i, n := range *input {
		numberTracker[n] = []int{i}
	}

	for i := len(*input); i < max; i++ {
		nLast := numbers[i-1]
		indices := numberTracker[nLast]
		if len(indices) == 1 {
			numberTracker[0] = append(numberTracker[0], i)
		} else {
			nNew := indices[len(indices)-1] - indices[len(indices)-2]
			numberTracker[nNew] = append(numberTracker[nNew], i)
			numbers[i] = nNew
		}
	}

	// fmt.Println(numbers)
	return numbers[len(numbers)-1]
}

// func solveBruteForce2(input *[]int, max int) int {
// 	numbers := make([]int, max)
// 	copy(numbers, *input)
// 	lastIndices := make([]int, max)

// 	for i := len(*input); i < max; i++ {
// 		nLast := numbers[i-1]
// 		iLast := lastIndices[i-1]
// 		if iLast == 0 {
// 			numberTracker[0] = append(numberTracker[0], i)
// 		} else {
// 			nNew := indices[len(indices)-1] - indices[len(indices)-2]
// 			numberTracker[nNew] = append(numberTracker[nNew], i)
// 			numbers[i] = nNew
// 		}
// 	}

// 	return numbers[len(numbers)-1]
// }

// SolveTask1 solve aoc task 1
func SolveTask1(input *[]int) int {
	return solveBruteForce(input, 2020)
}

// SolveTask2 solve aoc task 2
func SolveTask2(input *[]int) int {
	return solveBruteForce(input, 30000000)
}

// Run solve and print solutions
func Run() {
	input := []int{6, 3, 15, 13, 1, 0}

	fmt.Println("solution task 1:", SolveTask1(&input))
	fmt.Println("solution task 2:", "needs optimization ;-)" /*SolveTask2(&input)*/)
}
