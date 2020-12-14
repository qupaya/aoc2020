package day14

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyMask(t *testing.T) {
	assert.Equal(t, uint64(73), ApplyMask(11, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"))
	assert.Equal(t, uint64(101), ApplyMask(101, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"))
}

func TestSolveTask1(t *testing.T) {
	input := strings.Split(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`, "\n")
	assert.Equal(t, uint64(165), SolveTask1(&input))
}

func TestApplyMaskV2(t *testing.T) {
	assert.Equal(t, []uint64{26, 27, 58, 59}, ApplyMaskV2(42, "000000000000000000000000000000X1001X"))
}

func TestSolveTask2(t *testing.T) {
	input := strings.Split(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`, "\n")
	assert.Equal(t, uint64(208), SolveTask2(&input))
}
