package decoder

import (
	"aoc-2021-day16/packet"
	"aoc-2021-day16/packet/pkgtype"

	"github.com/dropbox/godropbox/container/bitvector"
)

func Decode(vector *bitvector.BitVector) packet.Packet {
	// todo:
	// 1. check if op -> go to decodeOp
	// 1.1 foreach subpacket -> check if op -> recursive decodeOp
	// 2. check if lv -> go to decodeLv

	packet := getPacket(vector)
	return packet
}

func getPacket(vector *bitvector.BitVector) packet.Packet {
	version := getFirstBits(vector, 3)
	deleteFirstBits(vector, 3)

	typeID := getFirstBits(vector, 3)
	deleteFirstBits(vector, 3)

	return packet.Packet{
		Version: version,
		TypeID:  pkgtype.ID(typeID),
		Body:    vector,
	}
}

func getFirstBits(vector *bitvector.BitVector, count int) int {
	// vector.Element(0)<<2 + vector.Element(1)<<1 + vector.Element(2)<<0
	num := 0
	for i, shift := 0, count-1; i < count; i++ {
		num += int(vector.Element(i) << shift)
		shift--
	}

	return num
}
func deleteFirstBits(vector *bitvector.BitVector, count int) {
	for i := 0; i < count; i++ {
		vector.Delete(0)
	}
}
