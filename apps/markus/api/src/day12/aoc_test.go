package day12

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveTask1(t *testing.T) {
	input := strings.Fields(`F10
	N3
	F7
	R90
	F11`)

	assert.Equal(t, 25, SolveTask1(&input))

	input = strings.Fields(`F10
	R90
	F10
	R90
	F10
	R90
	F10

	F10
	R90
	F10
	R90
	F10
	R90
	F10`)

	assert.Equal(t, 0, SolveTask1(&input))

	input = strings.Fields(`F10
	L90
	F10
	L90
	F10
	L90
	F10`)

	assert.Equal(t, 0, SolveTask1(&input))

}

func TestRotateWaypoint1(t *testing.T) {
	waypoint := [2]int{3, -1}

	RotateWaypoint(&waypoint, 90)
	assert.Equal(t, [2]int{1, 3}, waypoint)

	RotateWaypoint(&waypoint, 90)
	assert.Equal(t, [2]int{-3, 1}, waypoint)

	RotateWaypoint(&waypoint, 90)
	assert.Equal(t, [2]int{-1, -3}, waypoint)

	RotateWaypoint(&waypoint, 90)
	assert.Equal(t, [2]int{3, -1}, waypoint)

	RotateWaypoint(&waypoint, -90)
	RotateWaypoint(&waypoint, -90)
	RotateWaypoint(&waypoint, -90)
	RotateWaypoint(&waypoint, -90)
	assert.Equal(t, [2]int{3, -1}, waypoint)

	RotateWaypoint(&waypoint, 180)
	RotateWaypoint(&waypoint, 180)
	assert.Equal(t, [2]int{3, -1}, waypoint)

	RotateWaypoint(&waypoint, 270)
	RotateWaypoint(&waypoint, 90)
	assert.Equal(t, [2]int{3, -1}, waypoint)
}

func TestRotateWaypoint2(t *testing.T) {
	waypoint := [2]int{0, 1}

	RotateWaypoint(&waypoint, 90)
	assert.Equal(t, [2]int{-1, 0}, waypoint)

}

func TestSolveTask2(t *testing.T) {
	input := strings.Fields(`F10
	N3
	F7
	R90
	F11`)

	assert.Equal(t, 286, SolveTask2(&input))

}
