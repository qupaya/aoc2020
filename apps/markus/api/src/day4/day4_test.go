package day4

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveTask1(t *testing.T) {
	input := strings.Fields(`
  `)

	result := SolveTask1(&input)
	assert.Equal(t, -1, result)
}

func TestSolveTask2(t *testing.T) {
	input := strings.Fields(``)

	result := SolveTask2(&input)
	assert.Equal(t, -1, result)
}
