package day8

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// NOP no op
	NOP = 0
	// ACC increment/decrement accumulator
	ACC = 1
	// JMP jump to instruction
	JMP = 2
)

// Instruction 0: NOP|ACC|JMP 1: argument
type Instruction = [2]int

// ParseProgram parse program
func ParseProgram(input *[]string) *[]Instruction {
	program := make([]Instruction, len(*input))
	for i, line := range *input {
		fields := strings.Fields(line)
		argument, _ := strconv.Atoi(fields[1])
		switch fields[0] {
		case "nop":
			program[i] = Instruction{NOP, argument}
		case "acc":
			program[i] = Instruction{ACC, argument}
		case "jmp":
			program[i] = Instruction{JMP, argument}
		}

	}
	return &program
}

// RunProgram run program, return accumulator, currentIP, lastIP
func RunProgram(program *[]Instruction, maxRuns int) (int, int, int) {
	length := len(*program)
	ran := make(map[int]bool, length)
	lastIP := -1
	currentIP := 0
	accumulator := 0

	for i := 0; i < maxRuns; i++ {
		if ran[currentIP] == true {
			break
		}
		ran[currentIP] = true
		instruction := (*program)[currentIP][0]
		argument := (*program)[currentIP][1]

		lastIP = currentIP
		switch instruction {
		case NOP:
			currentIP = (currentIP + 1) % length
		case ACC:
			accumulator += argument
			currentIP = (currentIP + 1) % length
		case JMP:
			currentIP = (currentIP + argument) % length
		}
	}
	return accumulator, currentIP, lastIP
}

// SolveTask1 solve aoc task 1
func SolveTask1(input *[]string) int {
	acc, _, _ := RunProgram(ParseProgram(input), 100000)
	return acc
}

// SolveTask2 solve aoc task 1
func SolveTask2(input *[]string) int {
	program := *ParseProgram(input)
	for i := range program {
		if program[i][0] == ACC {
			continue
		}
		if program[i][0] == NOP || program[i][0] == JMP {
			orig := program[i][0]

			if orig == NOP {
				program[i][0] = JMP
			} else {
				program[i][0] = NOP
			}

			acc, _, lastIP := RunProgram(&program, 1000000)
			if lastIP == len(program)-1 {
				return acc
			}
			program[i][0] = orig
		}
	}
	panic("did not find solution")
}

// Run solve and print solutions
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(&input))
	fmt.Println("solution task 2:", SolveTask2(&input))
}
