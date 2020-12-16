package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveTask1(t *testing.T) {
	input := []int{0, 3, 6}
	assert.Equal(t, 436, SolveTask1(&input))
}

// func TestSolveTask2(t *testing.T) {
// 	input := []int{0, 3, 6}
// 	assert.Equal(t, -1, SolveTask2(&input))
// }
