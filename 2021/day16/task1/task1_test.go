package task1

import (
	"aoc-2021-day16/decoder"
	"aoc-2021-day16/parser"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestSolve_Example1(t *testing.T) {
	vector, err := parser.Parse("8A004A801A8002F478")
	require.NoError(t, err)

	result, _, err := decoder.Decode(vector)
	require.NoError(t, err)

	sum := SumVersions(result)

	assert.Equal(t, 16, sum)
}

func TestSolve_Example2(t *testing.T) {
	vector, err := parser.Parse("620080001611562C8802118E34")
	require.NoError(t, err)

	result, _, err := decoder.Decode(vector)
	require.NoError(t, err)

	sum := SumVersions(result)

	assert.Equal(t, 12, sum)
}

func TestSolve_Example3(t *testing.T) {
	vector, err := parser.Parse("C0015000016115A2E0802F182340")
	require.NoError(t, err)

	result, _, err := decoder.Decode(vector)
	require.NoError(t, err)

	sum := SumVersions(result)

	assert.Equal(t, 23, sum)
}

func TestSolve_Example4(t *testing.T) {
	vector, err := parser.Parse("A0016C880162017C3686B18A3D4780")
	require.NoError(t, err)

	result, _, err := decoder.Decode(vector)
	require.NoError(t, err)

	sum := SumVersions(result)

	assert.Equal(t, 31, sum)
}
