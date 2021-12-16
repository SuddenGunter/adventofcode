package decoder_test

import (
	"aoc-2021-day16/decoder"
	"aoc-2021-day16/packet/pkgtype"
	"aoc-2021-day16/parser"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestGetHeader_ForLiteralValue_DecodesOk(t *testing.T) {
	vector, err := parser.Parse("D2FE28")
	require.NoError(t, err)

	result := decoder.GetHeader(vector)

	assert.Equal(t, pkgtype.ID(4), result.TypeID)
	assert.Equal(t, 6, result.Version)
}

func TestGetHeader_ForOpValue_DecodesOk(t *testing.T) {
	vector, err := parser.Parse("38006F45291200")
	require.NoError(t, err)

	result := decoder.GetHeader(vector)

	assert.Equal(t, pkgtype.ID(6), result.TypeID)
	assert.Equal(t, 1, result.Version)
}
