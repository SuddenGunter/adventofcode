package parser_test

import (
	"aoc-2021-day16/parser"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestParse_LiteralPacket_InputReturnedAsBitVector(t *testing.T) {
	expected := "110100101111111000101000"

	vector, err := parser.Parse("D2FE28")

	require.NoError(t, err)

	for i := 0; i < vector.Length(); i++ {
		assert.Equal(t, expected[i]-'0', vector.Element(i))
	}
}

func TestParse_OpPacket_InputReturnedAsBitVector(t *testing.T) {
	expected := "00111000000000000110111101000101001010010001001000000000"

	vector, err := parser.Parse("38006F45291200")

	require.NoError(t, err)

	for i := 0; i < vector.Length(); i++ {
		assert.Equal(t, expected[i]-'0', vector.Element(i))
	}
}
