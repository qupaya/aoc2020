package day11

import (
	"fmt"
	"strings"
)

// PrepareBoardWithBorders add border with free space ('.')
func PrepareBoardWithBorders(input *[]string) *[][]rune {
	board := make([][]rune, len(*input)+2)

	board[0] = []rune(strings.Repeat(".", len((*input)[0])+2))
	for y, line := range *input {
		newLine := "."
		newLine += line
		newLine += "."
		board[y+1] = []rune(newLine)
	}

	board[len(*input)+1] = []rune(strings.Repeat(".", len((*input)[0])+2))

	return &board
}

// CountAdjacentOccupied count '#' directly adjacent to the given seat
func CountAdjacentOccupied(b *[][]rune, x int, y int) int {
	count := 0
	if (*b)[y][x] == '#' {
		count--
	}

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if (*b)[i][j] == '#' {
				count++
			}
		}
	}

	return count
}

// CountOccupiedInLineOfSight count '#' in all lines of sight
func CountOccupiedInLineOfSight(b *[][]rune, x int, y int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				count += isOccupiedInLineOfSight(b, x, y, j, i)
			}
		}
	}

	return count

}

func isOccupiedInLineOfSight(b *[][]rune, x int, y int, sightDeltaX int, sightDeltaY int) int {
	for {
		x += sightDeltaX
		y += sightDeltaY
		if x < 0 || y < 0 || x >= len((*b)[0]) || y >= len((*b)) || (*b)[y][x] == 'L' {
			return 0
		}
		if (*b)[y][x] == '#' {
			return 1
		}
	}
}

func performSeatingRound(input *[][]rune, countAdjacentOccupied func(*[][]rune, int, int) int, neededFree int) bool {
	seatsToOccupy := [][2]int{}
	seatsToFree := [][2]int{}

	for y, line := range *input {
		for x, c := range line {
			if c == '.' {
				continue
			}
			occupied := countAdjacentOccupied(input, x, y)
			if c == 'L' && occupied == 0 {
				seatsToOccupy = append(seatsToOccupy, [2]int{x, y})
			}
			if c == '#' && occupied >= neededFree {
				seatsToFree = append(seatsToFree, [2]int{x, y})
			}
		}
	}

	for _, seats := range seatsToOccupy {
		(*input)[seats[1]][seats[0]] = '#'
	}
	for _, seats := range seatsToFree {
		(*input)[seats[1]][seats[0]] = 'L'
	}

	return len(seatsToOccupy)+len(seatsToFree) != 0
}

func printBoard(b *[][]rune) {
	for _, line := range *b {
		fmt.Println(string(line))
	}
	fmt.Println()
}

func countOccupied(b *[][]rune) int {
	count := 0
	for _, line := range *b {
		for _, s := range line {
			if s == '#' {
				count++
			}
		}
	}
	return count
}

// SolveTask1 solve aoc task 1
func SolveTask1(input *[]string) int {
	board := PrepareBoardWithBorders(input)

	for performSeatingRound(board, CountAdjacentOccupied, 4) {
		// printBoard(board)
	}

	return countOccupied(board)
}

// SolveTask2 solve aoc task 2
func SolveTask2(input *[]string) int {
	board := PrepareBoardWithBorders(input)

	for performSeatingRound(board, CountOccupiedInLineOfSight, 5) {
		// printBoard(board)
	}

	return countOccupied(board)
}

// Run solve and print solutions
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(&input))
	fmt.Println("solution task 2:", SolveTask2(&input))
}
