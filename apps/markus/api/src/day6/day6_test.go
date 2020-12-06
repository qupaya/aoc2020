package day6

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDistinct(t *testing.T) {
	assert.Equal(t, 3, CountDistinctAnswers([]string{"abc"}))
	assert.Equal(t, 3, CountDistinctAnswers([]string{"a", "b", "c"}))
	assert.Equal(t, 3, CountDistinctAnswers([]string{"ab", "ac"}))
}

func TestCountUnion(t *testing.T) {
	assert.Equal(t, 3, CountIntersectingAnswers([]string{"abc"}))
	assert.Equal(t, 0, CountIntersectingAnswers([]string{"a", "b", "c"}))
	assert.Equal(t, 1, CountIntersectingAnswers([]string{"ab", "ac"}))
	assert.Equal(t, 1, CountIntersectingAnswers([]string{"a", "a", "a", "a"}))
	assert.Equal(t, 1, CountIntersectingAnswers([]string{"b"}))
}

func TestSolveTask1(t *testing.T) {
	input := strings.Split(`abc

	a
	b
	c

	ab
	ac

	a
	a
	a
	a

	b`, "\n\n")
	assert.Equal(t, 11, SolveTask1(&input))
}

func TestSolveTask2(t *testing.T) {
	input := strings.Split(`abc

	a
	b
	c

	ab
	ac

	a
	a
	a
	a

	b`, "\n\n")
	assert.Equal(t, 6, SolveTask2(&input))
}
