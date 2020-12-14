package day6

import (
	"fmt"
	"strings"
)

// CountDistinctAnswers count answers given by anyone
func CountDistinctAnswers(group []string) int {
	answerSet := map[rune]struct{}{}

	for _, answer := range strings.Join(group, "") {
		answerSet[answer] = struct{}{}
	}

	return len(answerSet)
}

// CountIntersectingAnswers count answers given by everyone
func CountIntersectingAnswers(group []string) int {
	amountOfGroups := len(group)
	answerCountMap := map[rune]int{}
	countAnsweredByEveryone := 0

	for _, answer := range strings.Join(group, "") {
		answerCountMap[answer]++
		if answerCountMap[answer] == amountOfGroups {
			countAnsweredByEveryone++
		}
	}

	return countAnsweredByEveryone
}

func sum(input *[][]string, countFunc func([]string) int) int {
	count := 0
	for _, answers := range *input {
		count += countFunc(answers)
	}
	return count
}

// SolveTask1 solve aoc task 1
func SolveTask1(input *[][]string) int {
	return sum(input, CountDistinctAnswers)

}

// SolveTask2 solve aoc task 1
func SolveTask2(input *[][]string) int {
	return sum(input, CountIntersectingAnswers)
}

// Run solve and print solutions
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(input))
	fmt.Println("solution task 2:", SolveTask2(input))
}
