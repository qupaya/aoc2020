package day4

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func parsePassport(input *string) (map[string]string, bool) {
	fields := strings.Fields(*input)

	if len(fields) < 7 {
		return nil, false
	}

	passport := make(map[string]string)
	sort.Strings(fields)
	// sort order: 0     1     2     3     4     5     6     7
	// sort order: "byr","cid","ecl","eyr","hcl","hgt","iyr","pid"

	if len(fields) == 8 {
		if !strings.HasPrefix(fields[1], "cid") {
			return nil, false
		}
		passport["cid"] = strings.Split(fields[1], ":")[1]
		fields = append(fields[:1], fields[2:]...)
	}

	fieldNames := []string{"byr", "ecl", "eyr", "hcl", "hgt", "iyr", "pid"}

	for i, fieldName := range fieldNames {
		if !strings.HasPrefix(fields[i], fieldName) {
			return nil, false
		}
		passport[fieldName] = strings.Split(fields[i], ":")[1]
	}

	return passport, true
}

// SolveTask1 count valid passports
func SolveTask1(input *[]string) int {
	count := 0

	for _, maybePassport := range *input {
		if _, valid := parsePassport(&maybePassport); valid {
			count++
		}
	}

	return count
}

// SolveTask2 count valid passports with value evaluation
func SolveTask2(input *[]string) int {
	valid4Digits := regexp.MustCompile(`^[0-9]{4}$`)
	validHgt := regexp.MustCompile(`^[0-9]{3}cm|[0-9]{2}in$`)
	validHcl := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	validEcl := regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
	validPid := regexp.MustCompile(`^[0-9]{9}$`)

	count := 0

	for _, maybePassport := range *input {
		passport, valid := parsePassport(&maybePassport)
		if !valid {
			continue
		}

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

// Run day4
func Run() {
	input := Input()

	fmt.Println("solution task 1:", SolveTask1(&input))
	fmt.Println("solution task 2:", SolveTask2(&input))

}
