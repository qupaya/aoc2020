package main

import (
	"aoc/apps/markus/api/src/day1"
	"aoc/apps/markus/api/src/day10"
	"aoc/apps/markus/api/src/day11"
	"aoc/apps/markus/api/src/day12"
	"aoc/apps/markus/api/src/day13"
	"aoc/apps/markus/api/src/day14"
	"aoc/apps/markus/api/src/day15"
	"aoc/apps/markus/api/src/day16"
	"aoc/apps/markus/api/src/day2"
	"aoc/apps/markus/api/src/day3"
	"aoc/apps/markus/api/src/day4"
	"aoc/apps/markus/api/src/day5"
	"aoc/apps/markus/api/src/day6"
	"aoc/apps/markus/api/src/day7"
	"aoc/apps/markus/api/src/day8"
	"aoc/apps/markus/api/src/day9"
	"flag"
	"fmt"
	"os"
)

func main() {
	days := []func(){
		day1.Run,
		day2.Run,
		day3.Run,
		day4.Run,
		day5.Run,
		day6.Run,
		day7.Run,
		day8.Run,
		day9.Run,
		day10.Run,
		day11.Run,
		day12.Run,
		day13.Run,
		day14.Run,
		day15.Run,
		day16.Run,
	}

	day := flag.Uint("day", 0, "[OPTIONAL] number of day to run or 0 to run all days")
	flag.Parse()

	if int(*day) != 0 {
		if int(*day) > len(days) {
			fmt.Fprintf(os.Stderr, "error: invalid day %v\n", *day)
			os.Exit(1)
		}
		fmt.Println("day", *day)
		days[*day-1]()

	} else {
		for i, day := range days {
			fmt.Println("day", i+1)
			day()
		}
	}
}
