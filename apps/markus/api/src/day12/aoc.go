package day12

import (
	"fmt"
	"math"
	"strconv"
)

// SolveTask1 solve aoc task 1
func SolveTask1(input *[]string) int {
	x := 0
	y := 0
	facing := 90 // e

	for _, inst := range *input {
		arg, _ := strconv.Atoi(inst[1:])

		switch command := inst[0]; command {
		case 'E':
			x += arg
		case 'W':
			x -= arg
		case 'N':
			y -= arg
		case 'S':
			y += arg
		case 'L':
			facing = (facing - arg + 360) % 360
		case 'R':
			facing = (facing + arg) % 360
		case 'F':
			switch facing {
			case 0: // N
				y -= arg
			case 90: // E
				x += arg
			case 180: // S
				y += arg
			case 270: // W
				x -= arg
			default:
				fmt.Println("unknown F", facing)
				panic("unknown F")
			}
		default:
			panic("unknown command")
		}
		//fmt.Println(x, y, facing)
	}

	return int(math.Abs(float64(x))) + int(math.Abs(float64(y)))
}

// RotateWaypoint rotate waypoint around ship
func RotateWaypoint(waypoint *[2]int, arg int) {
	degrees := (arg + 360) % 360
	newWaypoint := [2]int{}
	switch degrees {
	case 90:
		newWaypoint[0] = -(*waypoint)[1]
		newWaypoint[1] = (*waypoint)[0]
	case 180:
		newWaypoint[0] = -(*waypoint)[0]
		newWaypoint[1] = -(*waypoint)[1]
	case 270:
		newWaypoint[0] = (*waypoint)[1]
		newWaypoint[1] = -(*waypoint)[0]
	default:
		panic("unknown degrees")
	}
	(*waypoint)[0] = newWaypoint[0]
	(*waypoint)[1] = newWaypoint[1]
}

// SolveTask2 solve aoc task 2
func SolveTask2(input *[]string) int {
	x := 0
	y := 0
	waypoint := [2]int{10, -1}

	for _, inst := range *input {
		arg, _ := strconv.Atoi(inst[1:])

		switch command := inst[0]; command {
		case 'E':
			waypoint[0] += arg
		case 'W':
			waypoint[0] -= arg
		case 'N':
			waypoint[1] -= arg
		case 'S':
			waypoint[1] += arg
		case 'L':
			RotateWaypoint(&waypoint, -arg)
		case 'R':
			RotateWaypoint(&waypoint, arg)
		case 'F':
			x += arg * waypoint[0]
			y += arg * waypoint[1]
		default:
			panic("unknown command")
		}
		//fmt.Println(x, y, waypoint)
	}

	return int(math.Abs(float64(x))) + int(math.Abs(float64(y)))
}

// Run solve and print solutions
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(input))
	fmt.Println("solution task 2:", SolveTask2(input))
}
