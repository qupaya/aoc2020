package day9

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstIllegalNumber(t *testing.T) {
	input := strings.Fields(`35
	20
	15
	25
	47
	40
	62
	55
	65
	95
	102
	117
	150
	182
	127
	219
	299
	277
	309
	576`)

	assert.Equal(t, 127, FindFirstIllegalNumber(&input, 5))
}

func TestFindContiguousNumbers(t *testing.T) {
	input := strings.Fields(`35
	20
	15
	25
	47
	40
	62
	55
	65
	95
	102
	117
	150
	182
	127
	219
	299
	277
	309
	576`)

	assert.Equal(t, 62, CalculateContiguousNumbersCode(&input, 127))
}
