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

	h := GetHeader(vector)
	if pkgtype.IsLiteral(h.TypeID) {
		return parseLiteral(h, vector)
	} else {
		return parseOp(h, vector)
	}
}

func parseOp(h packet.Header, v *bitvector.BitVector) packet.Packet {
	return nil
}

func parseLiteral(h packet.Header, v *bitvector.BitVector) packet.Packet {
	return nil
}

func GetHeader(vector *bitvector.BitVector) packet.Header {
	version := getFirstBits(vector, 3)
	deleteFirstBits(vector, 3)

	typeID := getFirstBits(vector, 3)
	deleteFirstBits(vector, 3)

	return packet.Header{
		Version: version,
		TypeID:  pkgtype.ID(typeID),
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
