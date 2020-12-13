package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveSimultaneousCongruence(t *testing.T) {
	input := SimultaneousCongruence{
		CongruenceConstraint{remainder: 2, modulus: 3},
		CongruenceConstraint{remainder: 3, modulus: 4},
		CongruenceConstraint{remainder: 2, modulus: 5}}

	assert.Equal(t, int64(47), SolveSimultaneousCongruence(input))
}
