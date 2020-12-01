package main

import (
	"aoc/apps/markus/api/src/day1"
	"testing"
)

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
