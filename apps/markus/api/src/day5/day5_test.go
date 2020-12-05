package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeRow(t *testing.T) {
	assert.Equal(t, 44, ComputeRow("FBFBBFFRLR"))
}

func TestComputeColumn(t *testing.T) {
	assert.Equal(t, 5, ComputeColumn("FBFBBFFRLR"))
}

func TestComputeSeatId(t *testing.T) {
	assert.Equal(t, 357, ComputeSeatID("FBFBBFFRLR"))
	assert.Equal(t, 567, ComputeSeatID("BFFFBBFRRR"))
	assert.Equal(t, 119, ComputeSeatID("FFFBBBFRRR"))
	assert.Equal(t, 820, ComputeSeatID("BBFFBBFRLL"))
}

func TestSolveTask1(t *testing.T) {
	input := Input()
	assert.Equal(t, 951, SolveTask1(&input))
}

func TestSolveTask2(t *testing.T) {
	input := Input()
	assert.Equal(t, 653, SolveTask2(&input))
}
