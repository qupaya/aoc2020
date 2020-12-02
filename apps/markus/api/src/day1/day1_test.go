package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, 514579, result)
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
	assert.Equal(t, 241861950, result)
}
