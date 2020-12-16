package utils

import (
	"strconv"
	"strings"
)

// Lines get a list of all lines, each trimmed
func Lines(s *string) *[]string {
	lines := strings.Split(*s, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	return &lines
}

// LinesGrouped get a list of all lines, grouped by empty lines
func LinesGrouped(s *string) *[][]string {
	splittedByEmptyLine := strings.Split(*s, "\n\n")
	groups := make([][]string, len(splittedByEmptyLine))

	for i, group := range splittedByEmptyLine {
		lines := Lines(&group)
		groups[i] = *lines
	}
	return &groups
}

// FieldsGrouped like LinesGrouped, but lines as Fields
func FieldsGrouped(s *string) *[][]string {
	splittedByEmptyLine := strings.Split(*s, "\n\n")
	groups := make([][]string, len(splittedByEmptyLine))

	for i, group := range splittedByEmptyLine {
		groups[i] = strings.Fields(group)
	}
	return &groups
}

// ConvertToInts ConvertToInts
func ConvertToInts(input *[]string) *[]int {
	numbers := make([]int, len(*input))
	for i, n := range *input {
		numbers[i], _ = strconv.Atoi(n)
	}
	return &numbers
}
