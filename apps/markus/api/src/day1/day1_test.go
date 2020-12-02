package day1

import (
	"testing"
)


func TestSolveTask1(t *testing.T) {
	input := [...]int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	result := SolveTask1(input[:])
	if result != 514579 {
		t.Error("Expected 514579 but was", result)
	}
}

func TestSolveTask2(t *testing.T) {
	input := [...]int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	result := SolveTask2(input[:])
	if result != 241861950 {
		t.Error("Expected 241861950 but was", result)
	}
}
