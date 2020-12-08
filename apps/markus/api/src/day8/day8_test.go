package day8

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseProgram(t *testing.T) {
	input := strings.Split(`nop +0
	acc +1
	jmp +4
	acc +3
	jmp -3
	acc -99
	acc +1
	jmp -4
	acc +6`, "\n")

	expectedProgram := []Instruction{
		{NOP, 0},
		{ACC, 1},
		{JMP, 4},
		{ACC, 3},
		{JMP, -3},
		{ACC, -99},
		{ACC, 1},
		{JMP, -4},
		{ACC, 6}}

	assert.Equal(t, expectedProgram, *ParseProgram(&input))
}

func TestRunProgram(t *testing.T) {
	input := strings.Split(`nop +0
	acc +1
	jmp +4
	acc +3
	jmp -3
	acc -99
	acc +1
	jmp -4
	acc +6`, "\n")

	acc, ip, lastIP := RunProgram(ParseProgram(&input), 100)
	assert.Equal(t, 5, acc)
	assert.Equal(t, 1, ip)
	assert.Equal(t, 4, lastIP)
}

func TestSolveTask2(t *testing.T) {
	input := strings.Split(`nop +0
	acc +1
	jmp +4
	acc +3
	jmp -3
	acc -99
	acc +1
	jmp -4
	acc +6`, "\n")

	assert.Equal(t, 8, SolveTask2(&input))
}
