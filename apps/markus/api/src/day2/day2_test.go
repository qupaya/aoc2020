package day2

import (
	"testing"
)

func TestSolveTask1(t *testing.T) {
	input := [...]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	result := SolveTask1(input[:])
	if result != 2 {
		t.Error("Expected 2 but was", result)
	}
}

func TestSolveTask2(t *testing.T) {
	input := [...]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	result := SolveTask2(input[:])
	if result != 1 {
		t.Error("Expected 1 but was", result)
	}
}
