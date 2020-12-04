package day4

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// SolveTask1 of day 4
func SolveTask1(input *[]string) int {
	count := 0

	for _, maybePassport := range *input {
		fields := strings.Fields(maybePassport)
		sort.Strings(fields)

		if len(fields) == 8 &&
			strings.HasPrefix(fields[0], "byr") &&
			strings.HasPrefix(fields[1], "cid") &&
			strings.HasPrefix(fields[2], "ecl") &&
			strings.HasPrefix(fields[3], "eyr") &&
			strings.HasPrefix(fields[4], "hcl") &&
			strings.HasPrefix(fields[5], "hgt") &&
			strings.HasPrefix(fields[6], "iyr") &&
			strings.HasPrefix(fields[7], "pid") {
			count++
		} else if len(fields) == 7 &&
			strings.HasPrefix(fields[0], "byr") &&
			strings.HasPrefix(fields[1], "ecl") &&
			strings.HasPrefix(fields[2], "eyr") &&
			strings.HasPrefix(fields[3], "hcl") &&
			strings.HasPrefix(fields[4], "hgt") &&
			strings.HasPrefix(fields[5], "iyr") &&
			strings.HasPrefix(fields[6], "pid") {
			count++
		}
	}

	return count
}

// SolveTask2 of day 4
func SolveTask2(input *[]string) int {
	valid4Digits := regexp.MustCompile(`^[0-9]{4}$`)
	validHgt := regexp.MustCompile(`^[0-9]{3}cm|[0-9]{2}in$`)
	validHcl := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	validEcl := regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
	validPid := regexp.MustCompile(`^[0-9]{9}$`)

	count := 0

	for _, maybePassport := range *input {
		fields := strings.Fields(maybePassport)
		sort.Strings(fields)

		var passport map[string]string
		passport = make(map[string]string)

		if len(fields) == 8 &&
			strings.HasPrefix(fields[0], "byr") &&
			strings.HasPrefix(fields[1], "cid") &&
			strings.HasPrefix(fields[2], "ecl") &&
			strings.HasPrefix(fields[3], "eyr") &&
			strings.HasPrefix(fields[4], "hcl") &&
			strings.HasPrefix(fields[5], "hgt") &&
			strings.HasPrefix(fields[6], "iyr") &&
			strings.HasPrefix(fields[7], "pid") {
			passport["byr"] = strings.Split(fields[0], ":")[1]
			passport["cid"] = strings.Split(fields[1], ":")[1]
			passport["ecl"] = strings.Split(fields[2], ":")[1]
			passport["eyr"] = strings.Split(fields[3], ":")[1]
			passport["hcl"] = strings.Split(fields[4], ":")[1]
			passport["hgt"] = strings.Split(fields[5], ":")[1]
			passport["iyr"] = strings.Split(fields[6], ":")[1]
			passport["pid"] = strings.Split(fields[7], ":")[1]
		} else if len(fields) == 7 &&
			strings.HasPrefix(fields[0], "byr") &&
			strings.HasPrefix(fields[1], "ecl") &&
			strings.HasPrefix(fields[2], "eyr") &&
			strings.HasPrefix(fields[3], "hcl") &&
			strings.HasPrefix(fields[4], "hgt") &&
			strings.HasPrefix(fields[5], "iyr") &&
			strings.HasPrefix(fields[6], "pid") {
			passport["byr"] = strings.Split(fields[0], ":")[1]
			passport["ecl"] = strings.Split(fields[1], ":")[1]
			passport["eyr"] = strings.Split(fields[2], ":")[1]
			passport["hcl"] = strings.Split(fields[3], ":")[1]
			passport["hgt"] = strings.Split(fields[4], ":")[1]
			passport["iyr"] = strings.Split(fields[5], ":")[1]
			passport["pid"] = strings.Split(fields[6], ":")[1]
		}

		fmt.Println(passport)
		if !valid4Digits.MatchString(passport["byr"]) {
			continue
		}
		byr, err := strconv.Atoi(passport["byr"])
		if err != nil || byr < 1920 || byr > 2002 {
			continue
		}
		if !valid4Digits.MatchString(passport["iyr"]) {
			continue
		}
		iyr, err := strconv.Atoi(passport["iyr"])
		if err != nil || iyr < 2010 || iyr > 2020 {
			continue
		}

		eyr, err := strconv.Atoi(passport["eyr"])
		if err != nil || eyr < 2020 || eyr > 2030 {
			continue
		}

		hgt := passport["hgt"]
		if !validHgt.MatchString(hgt) {
			continue
		}
		hgtNumber, _ := strconv.Atoi(hgt[0 : len(hgt)-2])

		if strings.HasSuffix(hgt, "cm") && (hgtNumber < 150 || hgtNumber > 193) {
			continue
		} else if strings.HasSuffix(hgt, "in") && (hgtNumber < 59 || hgtNumber > 76) {
			continue
		}

		if !validHcl.MatchString(passport["hcl"]) {
			continue
		}

		if !validEcl.MatchString(passport["ecl"]) {
			continue
		}

		if !validPid.MatchString(passport["pid"]) {
			continue
		}

		count++
	}

	return count
}

// Run day3
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(&input))
	fmt.Println("solution task 2:", SolveTask2(&input))

}
