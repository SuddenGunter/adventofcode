package decoder

import (
	"aoc-2021-day16/packet"
	"aoc-2021-day16/packet/lentype"
	"aoc-2021-day16/packet/pkgtype"
	"errors"
	"fmt"

	"github.com/dropbox/godropbox/container/bitvector"
)

func Decode(vector *bitvector.BitVector) (packet.Packet, *bitvector.BitVector, error) {
	h, err := GetHeader(vector)
	if err != nil {
		return nil, nil, err
	}

	if pkgtype.IsLiteral(h.TypeID) {
		packet := parseLiteral(h, vector)
		return packet, vector, nil
	} else {
		packet, err := parseOp(h, vector)
		if err != nil {
			return nil, nil, err
		}

		return packet, vector, nil
	}
}

func parseOp(h packet.Header, v *bitvector.BitVector) (packet.Packet, error) {
	lenType, err := GetLen(v)
	if err != nil {
		return nil, err
	}

	root := packet.OpPacket{
		Header:     h,
		Len:        lenType,
		Subpackets: nil,
	}

	var subpackets []packet.Packet

	switch {
	case lentype.IsLenInBits(lenType.ID):
		subpackets, err = parseWithLenInBits(lenType.Value, v)
		if err != nil {
			return nil, err
		}
	case lentype.IsLenInBits(lenType.ID):
		subpackets, err = parseWithLenInPackets(lenType.Value, v)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown lentype: %v", lenType.ID)
	}

	// todo: do we need to trim zeroes at the end?

	root.Subpackets = subpackets

	return root, nil
}

func parseWithLenInPackets(maxReadLen int, v *bitvector.BitVector) ([]packet.Packet, error) {
	subpackets := make([]packet.Packet, 0)
	readPackets := 0
	// todo: what if we read more than allowed
	for maxReadLen > readPackets {
		packet, vector, err := Decode(v)
		if err != nil {
			return nil, err
		} else {
			v = vector
			subpackets = append(subpackets, packet)
		}

		readPackets++
	}

	if maxReadLen < readPackets {
		return nil, errors.New("read more packets than required, something went wrong")
	}

	return subpackets, nil
}

func parseWithLenInBits(maxReadLen int, v *bitvector.BitVector) ([]packet.Packet, error) {
	subpackets := make([]packet.Packet, 0)
	beforeReadLen := v.Length()
	// todo: what if we read more than allowed
	for maxReadLen > beforeReadLen-v.Length() {
		packet, vector, err := Decode(v)
		if err != nil {
			return nil, err
		} else {
			v = vector
			subpackets = append(subpackets, packet)
		}
	}

	if maxReadLen < beforeReadLen-v.Length() {
		return nil, errors.New("read more bits than required, something went wrong")
	}

	return subpackets, nil
}

func GetLen(v *bitvector.BitVector) (packet.Len, error) {
	id := getFirstBits(v, 1)
	deleteFirstBits(v, 1)

	if lentype.IsLenInBits(lentype.ID(id)) {
		bits := getFirstBits(v, lentype.BitsForLenInBits)
		deleteFirstBits(v, lentype.BitsForLenInBits)

		return packet.Len{
			ID:    lentype.ID(id),
			Value: bits,
		}, nil
	}

	if lentype.IsNumOfSubpackets(lentype.ID(id)) {
		bits := getFirstBits(v, lentype.BitsNumOfSubpackets)
		deleteFirstBits(v, lentype.BitsNumOfSubpackets)

		return packet.Len{
			ID:    lentype.ID(id),
			Value: bits,
		}, nil
	}

	return packet.Len{}, fmt.Errorf("invalid lentype.ID: %v", id)
}

func parseLiteral(h packet.Header, v *bitvector.BitVector) packet.Packet {
	value := 0

	for bit := 1; bit != 0; {
		bit = getFirstBits(v, 1)
		deleteFirstBits(v, 1)

		value = value << 4

		num := getFirstBits(v, 4)
		deleteFirstBits(v, 4)

		value += num
	}

	return packet.LVPacket{
		Header: h,
		Value:  value,
	}
}

func GetHeader(vector *bitvector.BitVector) (packet.Header, error) {
	if vector.Length() < 6 {
		return packet.Header{}, ErrCantParseHeaderEOF
	}

	version := getFirstBits(vector, 3)
	deleteFirstBits(vector, 3)

	typeID := getFirstBits(vector, 3)
	deleteFirstBits(vector, 3)

	return packet.Header{
		Version: version,
		TypeID:  pkgtype.ID(typeID),
	}, nil
}

func getFirstBits(vector *bitvector.BitVector, count int) int {
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

var ErrCantParseHeaderEOF = errors.New("cant parse header: EOF")
