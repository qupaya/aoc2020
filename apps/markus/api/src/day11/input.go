package day11

import "strings"

// Input aoc input
func Input() []string {
	return strings.Fields(`LLLLLLLLLLLLLLL.LL.L.LLLLLLLLLLLLLLL.LLLL.LLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLL
	LLLLLLLL..L.LLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLL.LLLLLLLL.LLLLLLLLL.LLLLL.LLLLLL.LLLLL.L.LLLLLL.
	LLLLLLLLLLL.LLLLLLLL.LLLLLLLLL.LLLLL.LLLL.LLLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LL.LL.LLLLLLLLLLLLLLL
	LLLLL.LLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLL.LLLL.LLLLLLLL.L.LLLLLLLLLLLLLL.
	LLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLL.LLL.LLLLLLLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLL.LLLLLLL.L
	LLLL..LLLLL.LLLLLLLL..LLLLLLLLLLLLLL..LLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL..LLLL.LLL.LLLLLLLLLLL
	LLLLL.LLLLLLLLLLLLLL.LLLLLLLLL.LLL.L.LLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.LLLLL.LLLLLLL.L
	.LLLLL..LLL.LLLL.LLLLLLLLLLLLL.LLLLL.LLLLLLLLL.LLLLL.LLLLLLLL.LLLL.LLLLLL.LLLLL.LLLLL.LLLLLLLLL
	LLLLLLLLLLL.LLL.LLLL.LLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLL.L.LLLLLLLLLLLLLLLLLL.LLLLLLLLL
	LLLLL.L.LLL.LLLLLLLL.LLLLLLLLL..LL.LLLLLLLL.LLLLLLLLLLLL.LLLLLLLLL.LLLLL..LLLLL.LLLLL.LLL.LLLLL
	L..L.L...LL.LL...L....LL......L.L.L...L....LL........LL.L..L...L.......L.L...L...LL......LL...L
	LLLLL.LLLLLLLLLLLL.LLLLLLLLLLL.LLLLL.LLLLL.LLL.LLLLLLL.LLLLLLLLLLL.LLLLLL.LLLLLLLLLLL.LLLLLLLLL
	LLLLL.LLLLL.LLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLL.LLLL.LLLL.LLLLLLLLL.LLLLLLLLLLLLLL.LLL.LLLLLLLL.
	LL.LL.LL.LL.LLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLL.L.LLLLLLLLLLLLLL.LLLL..LLLLLLLLLLLLLLL
	LLLL..LLLLLLLLLLLLLL.LLLLLLLLLLLL.LL.LLLLLLLLL.LLLLLL.LLLLLLL.LL.L.LLLLLL...LLLLLLLLL.LLLLLLLLL
	LLLLL..LLLL.LLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLL.LLLLLLLLL
	L.LLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLLL.LLLLLLLLL.LLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLL.LLLLLLLLL
	LL........LL....L.L.....LLL......LLL..L.L.LLLLL...LL....L.....LL..LL.L.......L.L.L.LL.L.L.L....
	.LLLLLLLLLLLL.LL.LLLLLLLLLLLLL.LLLLL.LLLLLLLLL.LLLLLLLLL.LLL.LLLL..LLLLLLLLLLLL..LLLLLLLLLLLLLL
	LLLLL.LLLLL.LLLLLLLL.LLLLLLLLL..LLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.L.L.LLLLLLLLLLLLL.LLLLLLL
	LLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LL..LLLLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.LLLLLLLLLLLLLLL
	LLLLL..LLLL.LLLLLLLL.L.LLLLLLL.LLL.L..LLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LLLLL.LLLLL.LLL
	LLLL..LLL.L...LLLLLL.LL.LLLLLL.LLLL.LLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLL.LLLLL.LL.LL.LLLLLLLLL
	LLLLL.LLLLL.LLLLLLLLLLLLLLLLLL..LLLLLLL.LLLLLL.LLLL.LLLL.LLLLLLLLL.LLLLLL.LLL...LLLLLL.LLLLLLLL
	....LLL.......L...LL.L..........L....LLL.L.LL..L..L........LLL..LL...L......L..LLL..L....L.....
	LLLLL.LLLLLLLL.LLLLL.LLLLLLLLL.LLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLL.LL.LL.LL.LLLLL.LLLLLLLLL
	LLLLL.LLLLL.LLLL.LLL.LLL.LLLLL.LLLLL.LLLL.LLLL.LLLLLLLLLLLLLLLLLLLLLLLLLL.L.LLLLLLLLLLLLLLLLL.L
	LLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLL...LLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLL.LLLLLLLLL
	LLLLL.LLLLL.LLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLL.LLLLLLLLLL.LLLLL.LLLLLLLLL
	LLLLL.LLLLLLLLLLLLLL.LLLLLLLLL.LLLLL.LL.LLLLLLLLLLLLLLLLLLL..LLLLL.LLLLLL.LLLLLLLLLLL.LLLLLLLLL
	LLLLLLLLLLLLLLLLL.LL.LLLLLLLLL.LLLLLLLLLLLLL.LL.LLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.LL.LL.LLLLLLLLL
	LLLLL.LLLLL.LL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLL.L.LLL.L.LLLLL.LLLLLLLLL
	LLLLL.LLLLL.LLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLL.LLLL..LL.LL.LLLLL.LLLLL.LLLLLLLLL
	.LL.L...LLLL..L..L..LL...LL....LL..L...LL.L..L...L...L.................LLL....L..L...L.LLL..L..
	LLLLLLLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLL.LLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLL
	LLLLLLLLLLL.LLLLLLLL.L.LLLLLLL.LLLLL.LLLLLLLLL.LLLLLL.LL.LLL.LLLLLLLLLLLL..L.LLLLLLLLLLLLLLLLLL
	LLL.L.LLL.LLLLLL.LLL.LL.LL.LLLLL.LLL.LLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLL.LL.LLLLL.LLLLLLLLL
	LLLLLLLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLL..LLLLLLL.L.LLLLLLLLL.LLLLLL.LLLLL.LLLL..LLLLLLLLL
	LLLLL.L.LLLLLLLL.LL..LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLL.L.LLLLL.LLL.L.LLLL.LLLLL.LLLLLLLLLLLLLLL
	LLLLLLL.LLLLLLLLLLLL.LLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLL.LLLLLLLLL
	LLLLLLLLLLLLLLLL.LLL.LLLLLLLLL.LLLLL.LLLLLLLLL.LLLLLLLLL.LL.L.LLL..LLLLLLLLLLLL.LLLLL.LLLLLLLLL
	LLLLL.LLLLL.LLLLLLLL.LLLLLLLLL.LLLLL.L.LLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLL.LLLLL.LLLLLLLLL
	LL.L....L.L..........L...L.L.......LL......L.LL...LL.....L.L.L...L....LL.L..L......L..LL..L....
	LLLLL.LLLLL.LLLLLLLL..LLLLLLLL.LLLLLL.LLLLLLL.LLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.L.LLL.LLLLLLLLL
	L.LLL.LLLLL.LLLLLLLL.LLLLLLLLLLLLLLL.LLLLL.LLL.LLLLL.LLLLLLLLLLLLLLLLLLLL.LLL.LLLLLLL.LLLLLLLLL
	LLLL.LLLLLLLLLLLLLLL.LLLLL.LLL.LL.LLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLL.LLLLL.LLLLL..LL
	LLLLL.LLLLL.LL.LLLLL..LLLLL.LL.LLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.LLLLL.LLLLLLLLL
	LLLLLLLLLLLLLLLLLLLL.LLLL.LLLL.LLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.L.LL.LLLLLL.LLLLLLLLL
	LLLLL.LLLLL.LLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLL.L.LLLLL.LLLLL.LLLLLLLLL
	........L....LLL.L....L...L...L.LL.......L..L..L.LL.L.L.L..L...L..L..L.........LLL.....LL.L..L.
	LLLLL.LL.LL.LLLLLLLLLLLLLLLLLL.L.LLL..LLLLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLL..L.LLL.LL
	LLLLLLLLLLL.LLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLL..LLLLLLLL.LL.LLLLL.LLLLLLLLL
	LLLLL.LLLLLLLLLLLLLL.LLLLL..LL.LLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.LLLLL.L.LLLLLLL
	LLLLL.LLLLL.LLLLLL.L.LLLLLLLLLL.LLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLL.L.L.LL.LLLL.LLLLLLLLL
	LLLLL.LLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLL.LLLLLLL.LLLLLLLLL.L.LLLL.LLLLLLLLLLL.LLLLLLLLL
	LLLLLLLLLLL.LLLLLLLL.LL.LLLLLL.LLLLLLLLLLLLLL..LLLLLLLLL.LLLLLLLLL.LLLL.L.LLLLLLLLLLL.LLLLLLLLL
	LLLLL.L.LLL.LLLLLLLLLLL.LLLLLL.LLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLL.LLLLLLLLL.LLLLL.LLL
	LLLLL.LLLLL.LLLLLLLL.L.LLLLLLL.LLLLLLLLLL.LLLL.LLLL.LLLL.LLLLLLLLL.LLLLLL..LLLL.LLLLL.LLLLLLLLL
	...L.....LL.L...L......LL..L.....LL.L...L..L...L....L......L....L......L.LL..L........L.......L
	LLLLLLLL..L.LLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLLLL..LLLLLLLLLLLLLL.LLLLLLL.LLLLL.LLLLL.L.LLLLLLL
	LLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLL.LLLLLL.LL..LLLLLLLL.LLLLLLLLLLLL.LLLLL.LLLLLLL.L
	LLLLLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLLLL..LLLLLLLLLLLLLLL
	LLLLL.LLLLL.LLLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLL.LLLLLLLLLLL
	LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLL
	....L..L....L....L..L..LL.L......LL..LLL..L.L.L.LL......L.LLLL.L...LLL...L.....L.L.L.L.LLLL....
	L.LLL.LLLLL.LLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLL..LLLL.LLL..LLLLLLLLL.LLLLLLLLLLL..LLLLLLLLLLLL.LL
	L.L.L.LLLLLLLLLLLL.LLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLL..LLLLLLLLL.LLLLLL.L..LL.LLLLL.LLLLLLLLL
	LLLLLLLLLLLLLLLLLLLLLL.LLLLL.L.LLLLL.LLLLLLLLL.LL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLL.LLL
	LLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLL..LLLLLLLLLLLL.LLLLL.LLL
	LLLLL.LLLLLLLLLLLL...LLL.LLLLLLLLLLL.LLLLLLLLL.LLLLL.LLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLL.LLLLLLLLL
	LLLLL.LLLLL.LL.LLLLLLLLLLLLL.LLLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLL..LLLLLLLL
	.L.LLLL..L.........LLL...L..L.L..LLL....LL...L..L......LL...L...L...L...L.....LLL.L...L.L..LLLL
	LLLLL.LLLLL.LLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLL.LLLLLLLLL
	LLLLL.LLLLLLLLLLLL.L.LLLLLLLLLL.L.LL.LLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLL.LLLLLLLLL.L.LLLLLLLLL
	LLLLLLL.LLL.L.LLLLLL.LLLLLLLLL.L.LLL.LLLLLLLLL..LLLLLLLL.LLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLL
	LLLLL.LLLLL.LLLLL.LL.LLLLLLLLL.L.LLL.L.LLLLLLLLLL.LLLLLL.LLLLLLL.LLLLLLLL.LLLLLLLLLLLLLLLLLLLLL
	LLLLL.L.L.LLLLL.LLLL.LLLLLLLLLLLLLLL.LLLLLLLLLLLLLL.LL.L.LLL.LLLLLLLLLLLL.LL.LL.LL.LLLLLLLLLLLL
	LLLLL.LLLLL.LLLLLLL..LLLLLLL.L.LLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.LLLLL.LLLLLLLLL
	LLLLLLLLLLL.LLLLLLLL.LLLLLLLLL.LLLL..LLLLLLLLL.LLLLLLLL..LLLL.LLLL.LLLLLLLLL.LL.LLLLL.LLLLLLLLL
	LLLLL.LLLLL.LLLLLLLL.LL.LLL.LL.LLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLL
	LLLLLLLLLLL.LLLLLLLL.LLLLLL.LL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLL
	.L.....L..LL.L.L.L.LL.L.L..L...L...L....LL.L..L.L..L...L..L.L.L..L........L.L.L.L..L.LL.L.L....
	LLLLL.LLLLL.LLLLLLLLLLLLLLL.LLLLLLLL.LLLLLLLLL.LLLLLLLLL..LLLLLLLL.LL.LLLLLLLLL.LLLLLLLLLLLLLLL
	LLLLL.LLLLL.LLLLL.L.LLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LLL.LLLLLLLLLL.L.LLLLL.LL.LL.LLLLLLLLL
	LLLL.LLLLLL.LLLLLLL..LL.LLLLLL.LLLLLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLL.LL.LLL.LLLLL..LLLL.LLLLLLLLL
	.LLLL.LLLLLLLLLLLLLL.LLLLLLLLL.LLLLL.LLLLLLLLLLLLLLLLLLL.LLLL.LLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLL
	LLLLLLLLLLLLLL.LLLLL.L.LLLLLLL.LLLLL..LLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLLLLLLLL.LL..LLLLL
	LLLLL.LLLLLLL.L.LLLL.LLLLLLLL.LLL.LLLLLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLLL.LLLLLLL..LLLLL.LLLL.LLL.
	LLLLLLLLL.L.LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL.LLLLL.LLLLLLLLLLLLL.LLLLLLLLLL.LLLLLLL.LLLLL.LLL
	..L..LLLL....L...L.L.L....L.....LL..L.....L..L............L.L....L.L..L.LL...L.L........L...L..
	LLLLL..LLLLLLLL.LLLLLLLLLLLLLL.LLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLL.LLLLLLLLLLL..LLL.LLLL
	LLLLL.LLLLL.LLLLLLLL.LLLLLL.LL.LLLLLLL.LLLLLLL.LLLLLLLLL.LL.LLL.LL.LLLLLL..LLLL.LLLLLLLLLLLLLLL
	LLLLL.LLLLLLLL.LLLLL.LL.LLLLL..LLLLL.LLLLLLLLL.LLLLLLLLLLLLLLLLLLL.LLLLLLLLLLLL.LLLLLLLLLLLLLLL
	LLLLL.LLL.LLLLLLLLL..LLLLLLLLL.LLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLLLLLLLL.LLLLL.LLLLLLLLL
	LLLLL.LLLLL.L.LLLLLL.LLLLLLLLL.LLLLLLLLLLLLLLL.LLLLLLLLL.LLLLLLLLL.LLLLLL.LLLLL.LL.LL.LLLL.LLLL`)
}
