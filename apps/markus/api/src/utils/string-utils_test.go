package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLines(t *testing.T) {
	s := `	line1
			line2
			line3`
	assert.Equal(t, &[]string{"line1", "line2", "line3"}, Lines(&s))
}

func TestLinesGrouped(t *testing.T) {
	s := `	group1 line1
			group1 line2

			group2 line1
			group2 line2`
	assert.Equal(t, &[][]string{{"group1 line1", "group1 line2"}, {"group2 line1", "group2 line2"}}, LinesGrouped(&s))
}

func TestFieldsGrouped(t *testing.T) {
	s := `	group1 line1
			group1 line2

			group2 line1
			group2 line2`
	assert.Equal(t, &[][]string{{"group1", "line1", "group1", "line2"}, {"group2", "line1", "group2", "line2"}}, FieldsGrouped(&s))
}
