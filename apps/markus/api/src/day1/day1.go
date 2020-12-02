package day1

import "fmt"

// SolveTask1 of day 1
func SolveTask1(input []int) int {
	for i, a := range input {
		for j, b := range input {
			if i != j && a+b == 2020 {
				return a * b
			}
		}
	}

	return -1
}

// SolveTask2 of day 1
func SolveTask2(input []int) int {
	for i, a := range input {
		for j, b := range input {
			for k, c := range input {
				if i != j && i != k && j != k && a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}

	return -1
}

// Run day1
func Run() {
	input := Input()
	fmt.Println("solution task 1:", SolveTask1(input))
	fmt.Println("solution task 2:", SolveTask2(input))
}
