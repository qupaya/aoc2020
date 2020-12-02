package day2

import (
	"fmt"
	"strconv"
	"strings"
)

// SolveTask1 of day 2
func SolveTask1(input []string) int {
	count := 0
	for _, v := range input {
		s := strings.Fields(v)
		minmax := strings.Split(s[0], "-")
		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])
		letter := strings.Split(s[1], ":")[0]
		password := s[2]
		passwordWithoutLetter := strings.Join(strings.Split(password, letter), "")
		charCount := len(password) - len(passwordWithoutLetter)

		if charCount >= min && charCount <= max {
			count++
		}
	}

	return count
}

// SolveTask2 of day 2
func SolveTask2(input []string) int {
	count := 0
	for _, v := range input {
		// 1-3 a: abcde
		s := strings.Fields(v)
		minmax := strings.Split(s[0], "-")
		first, _ := strconv.Atoi(minmax[0])
		second, _ := strconv.Atoi(minmax[1])
		letter := s[1][0]
		password := s[2]

		foundTimes := 0
		if len(password) >= first && password[first-1] == letter {
			foundTimes++
		}
		if len(password) >= second && password[second-1] == letter {
			foundTimes++
		}

		if foundTimes == 1 {
			count++
		}
	}

	return count
}

// Run day2
func Run() {
	input := Input()
	fmt.Println("solution task 1:", SolveTask1(input))
	fmt.Println("solution task 2:", SolveTask2(input))
}
