package day16

import (
	"fmt"
	"strings"
)

func isValidForAnyField(v int, fields *[]Field) bool {
	for _, f := range *fields {
		if v >= f.rule1[0] && v <= f.rule1[1] || v >= f.rule2[0] && v <= f.rule2[1] {
			return true
		}
	}
	return false
}

// SolveTask1 solve aoc task 1
func SolveTask1(fields *[]Field, tickets *[][]int) int {
	count := 0

	for _, ticket := range (*tickets)[1:] {
		for _, v := range ticket {
			if !isValidForAnyField(v, fields) {
				count += v
			}
		}
	}

	return count
}

func isValid(ticket *[]int, fields *[]Field) bool {
	for _, v := range *ticket {
		if !isValidForAnyField(v, fields) {
			return false
		}
	}
	return true
}

func isValidForField(tickets *[][]int, f *Field, column int) bool {
	for _, i := range *tickets {
		v := i[column]
		if v < (*f).rule1[0] || v > (*f).rule1[1] && v < (*f).rule2[0] || v > (*f).rule2[1] {
			return false
		}
	}
	return true
}

func contains(list *[]int, v int) bool {
	for _, i := range *list {
		if i == v {
			return true
		}
	}
	return false
}

func findMatchingField(tickets *[][]int, fields *[]Field, column int, alreadyMappedColumns *[]int) []int {
	fieldsFound := []int{}
	for i, f := range *fields {
		if !contains(alreadyMappedColumns, i) && isValidForField(tickets, &f, column) {
			fieldsFound = append(fieldsFound, i)
		}
	}
	if len(fieldsFound) == 0 {
		panic("no matching field found")
	}

	return fieldsFound
}

// SolveTask2 solve aoc task 2
func SolveTask2(fields *[]Field, tickets *[][]int) int {
	validTickets := [][]int{}

	for _, ticket := range (*tickets)[1:] {
		if isValid(&ticket, fields) {
			validTickets = append(validTickets, ticket)
		}
	}

	columnsToFieldMapping := make([]int, len(*fields))
	for i := range columnsToFieldMapping {
		columnsToFieldMapping[i] = -1
	}

	for column := 0; contains(&columnsToFieldMapping, -1); {
		if columnsToFieldMapping[column] != -1 {
			column++
			continue
		}
		foundFields := findMatchingField(&validTickets, fields, column, &columnsToFieldMapping)
		if len(foundFields) == 1 {
			columnsToFieldMapping[column] = foundFields[0]
			column = 0
		} else {
			column++
		}
	}

	yourTicket := (*tickets)[0]
	count := 1
	for i, f := range *fields {
		if strings.HasPrefix(f.name, "departure") {
			column := getColumnForFieldIndex(&columnsToFieldMapping, i)
			count *= yourTicket[column]
		}
	}
	return count
}

func getColumnForFieldIndex(columnsToFieldMapping *[]int, fieldIndex int) int {
	for i, c := range *columnsToFieldMapping {
		if c == fieldIndex {
			return i
		}
	}
	panic("did not find fieldIndex")
}

// Run solve and print solutions
func Run() {
	fields, tickets := Input()

	fmt.Println("solution task 1:", SolveTask1(fields, tickets))
	fmt.Println("solution task 2:", SolveTask2(fields, tickets))
}
