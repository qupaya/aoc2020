package day9

import (
	"fmt"
	"strconv"
)

func convertToInts(input *[]string) []int {
	numbers := make([]int, len(*input))
	for i, n := range *input {
		numbers[i], _ = strconv.Atoi(n)
	}
	return numbers
}

// FindFirstIllegalNumber find first number, that is not the sum of any two of the <preambleLength> previous numbers
func FindFirstIllegalNumber(input *[]string, preambleLength int) int {
	numbers := convertToInts(input)

	for i := preambleLength; i < len(numbers); i++ {
		testNumber := numbers[i]
		testRange := numbers[i-preambleLength : i]
		found := false
		for _, a := range testRange {
			for _, b := range testRange {
				if a == b {
					continue
				}
				if a+b == testNumber {
					found = true
				}
			}
		}
		if !found {
			return testNumber
		}
	}

	panic("nothing found")
}

func minMax(numbers *[]int) (int, int) {
	min := (*numbers)[0]
	max := (*numbers)[0]
	for _, x := range *numbers {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return min, max
}

// CalculateContiguousNumbersCode calculate min+max of the contiguous numbers that add up exactly to <testNumber>
func CalculateContiguousNumbersCode(input *[]string, testNumber int) int {
	numbers := convertToInts(input)

	for i := range numbers {
		sum := 0
		addedNumbers := []int{}
		for j := i; j < len(numbers); j++ {
			if sum > testNumber {
				continue
			}
			sum += numbers[j]
			addedNumbers = append(addedNumbers, numbers[j])
			if sum == testNumber {
				min, max := minMax(&addedNumbers)
				return min + max
			}
		}
	}

	panic("nothing found")
}

// SolveTask1 solve aoc task 1
func SolveTask1(input *[]string) int {
	return FindFirstIllegalNumber(input, 25)
}

// SolveTask2 solve aoc task 1
func SolveTask2(input *[]string, testNumber int) int {
	return CalculateContiguousNumbersCode(input, testNumber)
}

// Run solve and print solutions
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(input))
	fmt.Println("solution task 2:", SolveTask2(input, 70639851))
}
