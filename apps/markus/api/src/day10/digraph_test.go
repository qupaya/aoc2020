package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEdge(t *testing.T) {
	digraph := Digraph{}

	AddEdge(&digraph, 1, 2)
	assert.Equal(t, Digraph{1: Edges{2}}, digraph)

	AddEdge(&digraph, 1, 3)
	assert.Equal(t, Digraph{1: Edges{2, 3}}, digraph)

	AddEdge(&digraph, 2, 3)
	assert.Equal(t, Digraph{1: Edges{2, 3}, 2: Edges{3}}, digraph)
}

func TestCountPaths(t *testing.T) {
	digraph := Digraph{1: Edges{2, 3}, 2: Edges{3}}

	assert.Equal(t, 2, CountPaths(&digraph, 1, 3))
}
