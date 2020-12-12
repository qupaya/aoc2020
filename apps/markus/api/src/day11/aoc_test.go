package day11

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveTask1(t *testing.T) {
	input := strings.Fields(`L.LL.LL.LL
	LLLLLLL.LL
	L.L.L..L..
	LLLL.LL.LL
	L.LL.LL.LL
	L.LLLLL.LL
	..L.L.....
	LLLLLLLLLL
	L.LLLLLL.L
	L.LLLLL.LL`)

	assert.Equal(t, 37, SolveTask1(&input))
}

func TestCountAdjacentOccupied(t *testing.T) {
	input := strings.Fields(`
	#.#
	###
	#.#
	###`)

	assert.Equal(t, 2, CountAdjacentOccupied(PrepareBoardWithBorders(&input), 1, 1))
	assert.Equal(t, 3, CountAdjacentOccupied(PrepareBoardWithBorders(&input), 1, 2))
	assert.Equal(t, 5, CountAdjacentOccupied(PrepareBoardWithBorders(&input), 2, 1))
	assert.Equal(t, 6, CountAdjacentOccupied(PrepareBoardWithBorders(&input), 2, 2))
}

func TestCountOccupiedInLineOfSight(t *testing.T) {
	input := strings.Fields(`
	#.#
	###
	#.#
	###
	`)

	assert.Equal(t, 3, CountOccupiedInLineOfSight(PrepareBoardWithBorders(&input), 1, 1))
	assert.Equal(t, 4, CountOccupiedInLineOfSight(PrepareBoardWithBorders(&input), 1, 2))
	assert.Equal(t, 5, CountOccupiedInLineOfSight(PrepareBoardWithBorders(&input), 2, 1))
	assert.Equal(t, 7, CountOccupiedInLineOfSight(PrepareBoardWithBorders(&input), 2, 2))
}

func TestSolveTask2(t *testing.T) {
	input := strings.Fields(`L.LL.LL.LL
	LLLLLLL.LL
	L.L.L..L..
	LLLL.LL.LL
	L.LL.LL.LL
	L.LLLLL.LL
	..L.L.....
	LLLLLLLLLL
	L.LLLLLL.L
	L.LLLLL.LL`)

	assert.Equal(t, 26, SolveTask2(&input))
}
