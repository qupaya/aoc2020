package day14

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

// ApplyMask apply strange bit mask (X ignored, 0 and to set)
func ApplyMask(n uint64, mask string) uint64 {
	nApplied := new(big.Int).SetUint64(n)
	for i, c := range mask {
		if c != 'X' {
			bit := uint(0)
			if c == '1' {
				bit = uint(1)
			}
			nApplied.SetBit(nApplied, len(mask)-1-i, bit)
		}
	}

	return nApplied.Uint64()
}

// SolveTask1 solve aoc task 1
func SolveTask1(input *[]string) uint64 {
	mask := ""
	mem := map[uint64]uint64{}

	for _, instruction := range *input {
		if strings.HasPrefix(instruction, "mask") {
			mask = strings.Fields(instruction)[2]
		} else if strings.HasPrefix(instruction, "mem") {
			address, _ := strconv.Atoi(strings.Split(strings.Split(instruction, "[")[1], "]")[0])
			n, _ := strconv.Atoi(strings.Fields(instruction)[2])
			mem[uint64(address)] = ApplyMask(uint64(n), mask)
		}
	}

	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	return sum
}

// ApplyMaskV2 apply 2nd version of strange bit mask (0 ignored, 1 set, X set both 0 and 1)
func ApplyMaskV2(n uint64, mask string) []uint64 {
	nApplied := new(big.Int).SetUint64(n)

	floating := []int{}
	for i, c := range mask {
		if c == '1' {
			nApplied.SetBit(nApplied, len(mask)-1-i, uint(1))
		} else if c == 'X' {
			floating = append(floating, len(mask)-1-i)
		}
	}

	result := []uint64{}
	amount := int(math.Pow(float64(2), float64(len(floating))))
	for i := 0; i < amount; i++ {
		bits := new(big.Int).SetInt64(int64(i))
		newValue := new(big.Int).Set(nApplied)
		for j, f := range floating {
			newValue.SetBit(newValue, f, bits.Bit(len(floating)-1-j))
		}
		result = append(result, newValue.Uint64())
	}

	return result
}

// SolveTask2 solve aoc task 2
func SolveTask2(input *[]string) uint64 {
	mask := ""
	mem := map[uint64]uint64{}

	for _, instruction := range *input {
		if strings.HasPrefix(instruction, "mask") {
			mask = strings.Fields(instruction)[2]
		} else if strings.HasPrefix(instruction, "mem") {
			origAddress, _ := strconv.Atoi(strings.Split(strings.Split(instruction, "[")[1], "]")[0])
			addresses := ApplyMaskV2(uint64(origAddress), mask)
			n, _ := strconv.Atoi(strings.Fields(instruction)[2])
			for _, a := range addresses {
				mem[a] = uint64(n)
			}
		}
	}

	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	return sum
}

// Run solve and print solutions
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(input))
	fmt.Println("solution task 2:", SolveTask2(input))
}
