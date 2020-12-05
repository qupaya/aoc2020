package day5

import (
	"fmt"
	"math"
	"sort"
)

func binarySpacePartition(input string, lowerChar rune) int {
	min := 0
	max := int(math.Pow(2, float64(len(input))))

	for _, char := range input {
		half := ((max - min) + 1) / 2
		if char == lowerChar {
			max -= half
		} else {
			min += half
		}
	}
	return min
}

// ComputeRow get row
func ComputeRow(input string) int {
	return binarySpacePartition(input[0:7], 'F')
}

// ComputeColumn get col
func ComputeColumn(input string) int {
	return binarySpacePartition(input[len(input)-3:], 'L')
}

// ComputeSeatID compute seat ID
func ComputeSeatID(input string) int {
	return ComputeRow(input)*8 + ComputeColumn(input)
}

// SolveTask1 of day 5
func SolveTask1(input *[]string) int {
	maxSeatID := 0
	for _, pass := range *input {
		seatID := ComputeSeatID(pass)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	return maxSeatID

}

// SolveTask2 of day 5
func SolveTask2(input *[]string) int {
	seatIds := make([]int, len(*input))
	for i, pass := range *input {
		seatIds[i] = ComputeSeatID(pass)
	}

	sort.Ints(seatIds)
	for i, seatID := range seatIds[1 : len(seatIds)-1] {
		if seatID-seatIds[i] > 1 {
			return seatIds[i] + 1
		}
	}

	return 0
}

// Run day5
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(&input))
	fmt.Println("solution task 2:", SolveTask2(&input))
}
