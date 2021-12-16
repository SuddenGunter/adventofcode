package packet

import (
	"aoc-2021-day16/packet/lentype"
	"aoc-2021-day16/packet/pkgtype"

	"github.com/dropbox/godropbox/container/bitvector"
)

type ValPacket struct {
	Header Header
	Body   *bitvector.BitVector
}

type Header struct {
	Version int
	TypeID  pkgtype.ID
}

type OpPacket struct {
	Header     Header
	Len        lentype.ID
	Subpackets []Packet
}

type Packet interface{}
