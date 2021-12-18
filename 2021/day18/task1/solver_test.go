package task1

import (
	"aoc-2021-day18/input"
	"aoc-2021-day18/tree"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestSolve_SumsCorrectly(t *testing.T) {
	nodes, err := tree.ParseNodes(strings.Split("[1,1]\n[2,2]\n[3,3]\n[4,4]", "\n"))
	require.NoError(t, err)

	_, root, err := Solve(input.Data{Numbers: nodes})

	assert.Equal(t, "[[[[1,1],[2,2]],[3,3]],[4,4]]", root.String())
}
