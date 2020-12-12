package day10

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveTask1(t *testing.T) {
	input := strings.Fields(`16 10 15 5 1 11 7 19 6 12 4`)

	assert.Equal(t, 7*5, SolveTask1(&input))
}

func TestSolveTask2_example1(t *testing.T) {
	input := strings.Fields(`16 10 15 5 1 11 7 19 6 12 4`)

	assert.Equal(t, 8, SolveTask2(&input))
}

func TestSolveTask2_example2(t *testing.T) {
	input := strings.Fields(`28 33 18 42 31 14 46 20 48 47 24 23 49 45 19 38 39 11 1 32 25 35 8 17 7 9 4 2 34 10 3`)

	assert.Equal(t, 19208, SolveTask2(&input))
}

func TestSolveTask2_input(t *testing.T) {
	input := Input()
	assert.Equal(t, 1322306994176, SolveTask2(&input))
}
