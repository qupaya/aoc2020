package day13

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveTask1(t *testing.T) {
	assert.Equal(t, uint64(3966), SolveTask1(1006697, &[]int{13, 41, 641, 19, 17, 29, 661, 37, 23}))

}

func TestSolveTask2BruteForce(t *testing.T) {
	assert.Equal(t, uint64(3417), SolveTask2BruteForce(0, strings.Split(`17,x,13,19`, ",")))
	assert.Equal(t, uint64(3417), SolveTask2BruteForce(3416, strings.Split(`17,x,13,19`, ",")))
	assert.Equal(t, uint64(1068781), SolveTask2BruteForce(0, strings.Split(`7,13,x,x,59,x,31,19`, ",")))
	assert.Equal(t, uint64(1068781), SolveTask2BruteForce(1068780, strings.Split(`7,13,x,x,59,x,31,19`, ",")))
}

func TestSolveTask2ChineseRemainderTheorem(t *testing.T) {
	assert.Equal(t, int64(3417), SolveTask2ChineseRemainderTheorem(strings.Split(`17,x,13,19`, ",")))
	assert.Equal(t, int64(1068781), SolveTask2ChineseRemainderTheorem(strings.Split(`7,13,x,x,59,x,31,19`, ",")))
	assert.Equal(t, int64(800177252346225), SolveTask2ChineseRemainderTheorem(Input()))
}
