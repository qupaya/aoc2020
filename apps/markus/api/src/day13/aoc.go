package day13

import (
	"fmt"
	"math"
	"strconv"
)

func calcWaitTime(now uint64, bus uint64) uint64 {
	return (bus - (now % bus)) % bus
}

// SolveTask1 solve aoc task 1
func SolveTask1(timestamp uint64, IDs *[]int) uint64 {
	minID := uint64(0)
	minWaitTime := timestamp

	for _, id := range *IDs {
		waitTime := calcWaitTime(timestamp, uint64(id))
		if waitTime < minWaitTime {
			minID = uint64(id)
			minWaitTime = waitTime
		}
	}
	return minID * minWaitTime
}

// SolveTask2BruteForce solve aoc task 2 (brute force); not feasible for real input
func SolveTask2BruteForce(startTimestamp uint64, input []string) uint64 {
	indexOfMaxBusID := 0
	constraints := [][2]uint64{} // [[time, id],[time, id], ...]
	for i, s := range input {
		if s != "x" {
			id, _ := strconv.Atoi(s)
			c := [2]uint64{uint64(i), uint64(id)}
			constraints = append(constraints, c)

			if i > 0 && uint64(id) > constraints[indexOfMaxBusID][1] {
				indexOfMaxBusID = len(constraints) - 1
			}
		}
	}

	fmt.Println("constraints", constraints)
	fmt.Println("max constraint", constraints[indexOfMaxBusID])

	maxModule := constraints[indexOfMaxBusID][1]
	maxModuleTime := constraints[indexOfMaxBusID][0]

	timestamp := uint64(0)
	if startTimestamp > startTimestamp%maxModule {
		timestamp = startTimestamp - startTimestamp%maxModule
	} else {
		timestamp = maxModule
	}
	// optimization: only test every <maxModule> time stamps
	for ; timestamp < math.MaxUint64-1; timestamp += maxModule {
		gotcha := true
		for _, c := range constraints {
			if (timestamp-(maxModuleTime-c[0]))%c[1] != 0 {
				gotcha = false
				break
			}
		}

		if gotcha {
			return timestamp - maxModuleTime
		}
	}
	panic("did not find solution")
}

// SolveTask2ChineseRemainderTheorem solve aoc task 2 (using Chinese Remainder Theorem)
func SolveTask2ChineseRemainderTheorem(input []string) int64 {
	crtInput := SimultaneousCongruence{}
	for i, s := range input {
		if s != "x" {
			modulus, _ := strconv.Atoi(s)
			remainder := (modulus - i) % modulus
			c := CongruenceConstraint{int64(remainder), int64(modulus)}
			crtInput = append(crtInput, c)

		}
	}

	return SolveSimultaneousCongruence(crtInput)
}

// Run solve and pruint64 solutions
func Run() {
	// input := Input()

	// fmt.Println("solution task 1:", SolveTask1(1006697, &[]int{13, 41, 641, 19, 17, 29, 661, 37, 23}))
	// // fmt.Println("solution task 2:", SolveTask2BruteForce(100000000000000, input))
	// fmt.Println("solution task 2:", SolveTask2ChineseRemainderTheorem(input))

}
