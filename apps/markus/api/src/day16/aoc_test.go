package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFields(t *testing.T) {
	input := `class: 1-3 or 5-7
	row: 6-11 or 33-44
	seat: 13-40 or 45-50

	your ticket:
	7,1,14

	nearby tickets:
	7,3,47
	40,4,50
	55,2,20
	38,6,12`

	assert.Equal(t, 3, len(*ParseFields(&input)))
}

func TestParseTickets(t *testing.T) {
	input := `class: 1-3 or 5-7
	row: 6-11 or 33-44
	seat: 13-40 or 45-50

	your ticket:
	7,1,14

	nearby tickets:
	7,3,47
	40,4,50
	55,2,20
	38,6,12`

	assert.Equal(t, 5, len(*ParseTickets(&input)))
}

func TestSolveTask1(t *testing.T) {
	input := `class: 1-3 or 5-7
	row: 6-11 or 33-44
	seat: 13-40 or 45-50

	your ticket:
	7,1,14

	nearby tickets:
	7,3,47
	40,4,50
	55,2,20
	38,6,12`

	assert.Equal(t, 71, SolveTask1(ParseFields(&input), ParseTickets(&input)))
}

func TestSolveTask2_givenExample(t *testing.T) {
	input := `departure class: 0-1 or 4-19
	departure row: 0-5 or 8-19
	seat: 0-13 or 16-19

	your ticket:
	11,12,13

	nearby tickets:
	100,1,1
	3,9,18
	15,1,5
	1,1000,1
	5,14,9
	1,1,15676587`

	assert.Equal(t, 12*11, SolveTask2(ParseFields(&input), ParseTickets(&input)))
}

func TestSolveTask2_withUnclearMatch(t *testing.T) {
	input := `	departure class: 0-1 or 3-6
				departure row: 0-1 or 3-5

				your ticket:
				11,12

				nearby tickets:
				1,1
				1,1
				1,6`

	assert.Equal(t, 12*11, SolveTask2(ParseFields(&input), ParseTickets(&input)))
}

func TestSolveTask2_realInput(t *testing.T) {
	fields, tickets := Input()

	assert.Equal(t, 410460648673, SolveTask2(fields, tickets))
}
