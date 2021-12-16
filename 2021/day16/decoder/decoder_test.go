package decoder

import (
	"aoc-2021-day16/packet"
	"aoc-2021-day16/packet/lentype"
	"aoc-2021-day16/packet/pkgtype"
	"aoc-2021-day16/parser"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetHeader_ForLiteralValue_DecodesOk(t *testing.T) {
	vector, err := parser.Parse("D2FE28")
	require.NoError(t, err)

	result, err := GetHeader(vector)
	require.NoError(t, err)

	assert.Equal(t, pkgtype.ID(4), result.TypeID)
	assert.Equal(t, 6, result.Version)
}

func TestGetHeader_ForOpValue_DecodesOk(t *testing.T) {
	vector, err := parser.Parse("38006F45291200")
	require.NoError(t, err)

	result, err := GetHeader(vector)
	require.NoError(t, err)

	assert.Equal(t, pkgtype.ID(6), result.TypeID)
	assert.Equal(t, 1, result.Version)
}

func TestDecode_ForLiteralValue_DecodesOk(t *testing.T) {
	vector, err := parser.Parse("D2FE28")
	require.NoError(t, err)

	result, v, err := Decode(vector)
	require.NoError(t, err)

	p, ok := result.(packet.LVPacket)

	assert.True(t, ok)
	assert.Equal(t, pkgtype.ID(4), p.Header.TypeID)
	assert.Equal(t, 6, p.Header.Version)
	assert.Equal(t, 3, v.Length())
}

func TestDecode_ForOpValueSample1_DecodesOk(t *testing.T) {
	vector, err := parser.Parse("38006F45291200")
	require.NoError(t, err)

	result, v, err := Decode(vector)
	require.NoError(t, err)

	p, ok := result.(packet.OpPacket)

	assert.True(t, ok)

	assert.Equal(t, pkgtype.ID(6), p.Header.TypeID)
	assert.Equal(t, 1, p.Header.Version)

	assert.Equal(t, 27, p.Len.Value)
	assert.Equal(t, lentype.ID(0), p.Len.ID)

	assert.Equal(t, 7, v.Length())
	assert.Equal(t, 2, len(p.Subpackets))

	s1, ok := p.Subpackets[0].(packet.LVPacket)

	assert.True(t, ok)

	assert.Equal(t, pkgtype.ID(4), s1.Header.TypeID)
	assert.Equal(t, 6, s1.Header.Version)
	assert.Equal(t, 10, s1.Value)

	s2, ok := p.Subpackets[1].(packet.LVPacket)

	assert.True(t, ok)

	assert.Equal(t, pkgtype.ID(4), s2.Header.TypeID)
	assert.Equal(t, 2, s2.Header.Version)
	assert.Equal(t, 20, s2.Value)
}
