package task2

import (
	"aoc-2021-day18/input"
	"aoc-2021-day18/tree"
	"fmt"
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

func TestMagnitude_Example1WorksCorrectly(t *testing.T) {
	node, err := tree.ParseNode("[9,1]")
	require.NoError(t, err)

	val := magnitude(node)

	assert.Equal(t, 29, val)
}

func TestMagnitude_Example2WorksCorrectly(t *testing.T) {
	node, err := tree.ParseNode("[1,9]")
	require.NoError(t, err)

	val := magnitude(node)

	assert.Equal(t, 21, val)
}

func TestMagnitude_Example3WorksCorrectly(t *testing.T) {
	node, err := tree.ParseNode("[[1,2],[[3,4],5]]")
	require.NoError(t, err)

	val := magnitude(node)

	assert.Equal(t, 143, val)
}

func TestMagnitude_Example4WorksCorrectly(t *testing.T) {
	node, err := tree.ParseNode("[[9,1],[1,9]]")
	require.NoError(t, err)

	val := magnitude(node)

	assert.Equal(t, 129, val)
}

func TestSolver_Example1WorksCorrectly(t *testing.T) {
	nodes, err := tree.ParseNodes(strings.Split("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]\n[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]\n[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]\n[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]\n[7,[5,[[3,8],[1,4]]]]\n[[2,[2,2]],[8,[8,1]]]\n[2,9]\n[1,[[[9,3],9],[[9,0],[0,7]]]]\n[[[5,[7,4]],7],1]\n[[[[4,2],2],6],[8,7]]", "\n"))
	require.NoError(t, err)

	_, root, err := Solve(input.Data{Numbers: nodes})

	fmt.Println(root.String())

	assert.Equal(t, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", root.String())
}

func TestSolver_Example2WorksCorrectly(t *testing.T) {
	nodes, err := tree.ParseNodes(strings.Split("[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]\n[[[5,[2,8]],4],[5,[[9,9],0]]]\n[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]\n[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]\n[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]\n[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]\n[[[[5,4],[7,7]],8],[[8,3],8]]\n[[9,3],[[9,9],[6,[4,9]]]]\n[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]\n[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]", "\n"))
	require.NoError(t, err)

	magnitude, root, err := Solve(input.Data{Numbers: nodes})

	fmt.Println(root.String())

	assert.Equal(t, "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", root.String())
	assert.Equal(t, 4140, magnitude)
}
