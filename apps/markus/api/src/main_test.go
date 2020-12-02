package main

import (
	"aoc/apps/markus/api/src/day1"
	"aoc/apps/markus/api/src/day2"
	"testing"
)

func TestDay2_Task1(t *testing.T) {
	input := [...]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	result := day2.SolveTask1(input[:])
	if result != 2 {
		t.Error("Expected 2 but was", result)
	}
}

func TestDay2_Task2(t *testing.T) {
	input := [...]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	result := day2.SolveTask2(input[:])
	if result != 1 {
		t.Error("Expected 1 but was", result)
	}
}

func TestDay1_Task1(t *testing.T) {
	input := [...]int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	result := day1.SolveTask1(input[:])
	if result != 514579 {
		t.Error("Expected 514579 but was", result)
	}
}

func TestDay1_Task2(t *testing.T) {
	input := [...]int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	result := day1.SolveTask2(input[:])
	if result != 241861950 {
		t.Error("Expected 241861950 but was", result)
	}
}
