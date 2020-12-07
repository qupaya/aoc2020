package day7

import (
	"fmt"
	"strconv"
	"strings"
)

// Bags map of bag names with contained bags (and its count)
type Bags map[string]map[string]int

// ParseRule parse 1 rule
func ParseRule(rule string) (string, map[string]int) {
	s1 := strings.Split(strings.Trim(rule[:len(rule)-5], " "), " bags contain ")
	target := s1[0]
	content := s1[1]
	contains := map[string]int{}

	if content == "no other" {
		return target, contains
	}

	splittedContent := strings.Split(content, " bag")

	for _, s := range splittedContent {
		if strings.HasPrefix(s, ",") || strings.HasPrefix(s, "s,") {
			s = strings.Trim(s[2:], " ") // should have used regex 🙈️
		}
		n, _ := strconv.Atoi(strings.Trim(s[:2], " "))
		bagType := s[2:]
		contains[bagType] = n
	}

	return target, contains
}

// ParseRules parse all rules
func ParseRules(rules *[]string) Bags {
	parsed := make(Bags, len(*rules))

	for _, rule := range *rules {
		target, content := ParseRule(rule)
		parsed[target] = content
	}

	return parsed
}

func canContain(target string, searchIn string, bags Bags) bool {
	if len(bags[searchIn]) == 0 {
		return false
	}

	for bag := range bags[searchIn] {
		if bag == target {
			return true
		}
		if canContain(target, bag, bags) {
			return true
		}
	}

	return false
}

// CountPossibilities count how many bags can eventually contain target bag
func CountPossibilities(target string, bags Bags) int {
	count := 0

	for bag := range bags {
		if canContain(target, bag, bags) {
			count++
		}
	}

	return count
}

// CountBags count how many bags target bag contains
func CountBags(target string, bags Bags) int {
	count := 0

	for bag, amount := range bags[target] {
		count += amount
		count += amount * CountBags(bag, bags)
	}

	return count
}

// SolveTask1 solve aoc task 1
func SolveTask1(input *[]string) int {
	return CountPossibilities("shiny gold", ParseRules(input))
}

// SolveTask2 solve aoc task 1
func SolveTask2(input *[]string) int {
	return CountBags("shiny gold", ParseRules(input))
}

// Run solve and print solutions
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(&input))
	fmt.Println("solution task 2:", SolveTask2(&input))
}
