package packet

import (
	"aoc-2021-day16/packet/lentype"
	"aoc-2021-day16/packet/pkgtype"
)

type LVPacket struct {
	Value  int
	Header Header
}

type Header struct {
	Version int
	TypeID  pkgtype.ID
}

type OpPacket struct {
	Header     Header
	Len        Len
	Subpackets []Packet
}

type Len struct {
	ID    lentype.ID
	Value int
}

type Packet interface{}
