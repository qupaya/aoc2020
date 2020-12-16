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

	return numbers[len(numbers)-1]
}

func solveBruteForce2(input *[]int, max int) int {
	numbers := make([]int, max)
	backtrack := make([]int, max)
	copy(numbers, *input)

	lastOccurrence := make(map[int]int, max)

	for i := range *input {
		backtrack[i] = i
		lastOccurrence[(*input)[i]] = i
	}

	for i := len(*input); i < max; i++ {
		numbers[i] = (i - 1 - backtrack[i-1]) % i

		lastOcc, has := lastOccurrence[numbers[i]]
		if has {
			backtrack[i] = lastOcc
		} else {
			backtrack[i] = i
		}
		lastOccurrence[numbers[i]] = i
	}

	return numbers[len(numbers)-1]
}

func solveBruteForce3(input *[]int, max int) int {
	numbers := make([]int, max)
	backtrack := make([]int, max)
	copy(numbers, *input)
	lastOccurrence := make([]int, max)

	for i := range *input {
		backtrack[i] = i + 1
		lastOccurrence[(*input)[i]] = i + 1
	}

	for i := len(*input); i < max; i++ {
		numbers[i] = (i - backtrack[i-1])

		lastOcc := lastOccurrence[numbers[i]]
		if lastOcc != 0 {
			backtrack[i] = lastOcc
		} else {
			backtrack[i] = i + 1
		}
		lastOccurrence[numbers[i]] = i + 1
	}

	return numbers[len(numbers)-1]
}

// SolveTask1 solve aoc task 1
func SolveTask1(input *[]int) int {
	// return solveBruteForce(input, 2020)
	// return solveBruteForce2(input, 2020)
	return solveBruteForce3(input, 2020)
}

// SolveTask2 solve aoc task 2
func SolveTask2(input *[]int) int {
	// return solveBruteForce(input, 30000000)
	// return solveBruteForce2(input, 30000000)
	return solveBruteForce3(input, 30000000)
}

// Run solve and print solutions
func Run() {
	input := []int{6, 3, 15, 13, 1, 0}

	fmt.Println("solution task 1:", SolveTask1(&input))
	fmt.Println("solution task 2:", SolveTask2(&input))
}
