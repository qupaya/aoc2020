package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveTask1(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	result := SolveTask1(&input)
	assert.Equal(t, 2, result)
}

func TestSolveTask2(t *testing.T) {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	result := SolveTask2(&input)
	assert.Equal(t, 1, result)
}
