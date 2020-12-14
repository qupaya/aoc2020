package day10

import (
	"fmt"
	"sort"
	"strconv"
)

// ConvertToInts ConvertToInts
func ConvertToInts(input *[]string) []int {
	numbers := make([]int, len(*input))
	for i, n := range *input {
		numbers[i], _ = strconv.Atoi(n)
	}
	return numbers
}

// SolveTask1 solve aoc task 1
func SolveTask1(input *[]string) int {
	adapters := ConvertToInts(input)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	currentJolt := 0
	diffOnes := 0
	diffThrees := 0

	for _, a := range adapters {
		diff := a - currentJolt
		if diff == 1 {
			diffOnes++
		}
		if diff == 3 {
			diffThrees++
		}
		currentJolt = a
	}

	return diffOnes * diffThrees
}

func SolveTask2(input *[]string) int {
	adapters := ConvertToInts(input)
	sort.Ints(adapters)
	adapters = append([]int{0}, adapters...)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	digraph := Digraph{}

	incomingEdgesPerVertex := map[Vertex]int{}
	breakpoints := []Vertex{0}

	for i := 0; i < len(adapters); i++ {
		for _, j := range []int{1, 2, 3} {
			if i+j < len(adapters) && adapters[i+j]-adapters[i] <= 3 {
				AddEdge(&digraph, Vertex(adapters[i]), Vertex(adapters[i+j]))
				incomingEdgesPerVertex[Vertex(adapters[i+j])]++
			}
		}
		if i+1 < len(adapters) && adapters[i+1]-adapters[i] == 3 {
			breakpoints = append(breakpoints, Vertex(adapters[i]))
		}
	}

	breakpoints = append(breakpoints, Vertex(adapters[len(adapters)-1]))

	count := 1
	for i := 1; i < len(breakpoints); i++ {
		segmentCount := CountPaths(&digraph, breakpoints[i-1], breakpoints[i])
		count *= segmentCount
	}

	return count
}

// Run solve and print solutions
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(input))
	fmt.Println("solution task 2:", SolveTask2(input))
}
