package task2

import (
	"aoc-2021-day16/decoder"
	"aoc-2021-day16/packet"
	"aoc-2021-day16/packet/pkgtype"
	"aoc-2021-day16/parser"
	"errors"
	"fmt"
	"math"
)

func Solve(data string) (int, error) {
	vector, err := parser.Parse(data)
	if err != nil {
		return 0, err
	}

	packet, _, err := decoder.Decode(vector)
	if err != nil {
		return 0, err
	}

	result, err := calculateResult(packet)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func calculateResult(p packet.Packet) (int, error) {
	asOp, ok := p.(packet.OpPacket)
	if !ok {
		asVal, ok := p.(packet.LVPacket)
		if !ok {
			return 0, errors.New("unknown packet type")
		}

		return asVal.Value, nil
	}

	switch asOp.Header.TypeID {
	case pkgtype.Sum:
		return sum(asOp.Subpackets)
	case pkgtype.Product:
		return product(asOp.Subpackets)
	case pkgtype.Min:
		return min(asOp.Subpackets)
	case pkgtype.Max:
		return max(asOp.Subpackets)
	case pkgtype.Gt:
		return gt(asOp.Subpackets)
	case pkgtype.Lt:
		return lt(asOp.Subpackets)
	case pkgtype.Eq:
		return eq(asOp.Subpackets)

	default:
		return 0, fmt.Errorf("unsupported Op: %v", asOp.Header.TypeID)
	}
}

func product(packets []packet.Packet) (int, error) {
	prod := 1

	if len(packets) == 0 {
		return 0, errors.New("unsupported operation: product of 0 packets")
	}

	for _, p := range packets {
		val, err := calculateResult(p)
		if err != nil {
			return 0, err
		}

		prod *= val
	}

	return prod, nil
}

func sum(packets []packet.Packet) (int, error) {
	sum := 0

	if len(packets) == 0 {
		return 0, errors.New("unsupported operation: sum of 0 packets")
	}

	for _, p := range packets {
		val, err := calculateResult(p)
		if err != nil {
			return 0, err
		}

		sum += val
	}

	return sum, nil
}

func min(packets []packet.Packet) (int, error) {
	min := math.MaxInt

	if len(packets) == 0 {
		return 0, errors.New("unsupported operation: min of 0 packets")
	}

	for _, p := range packets {
		val, err := calculateResult(p)
		if err != nil {
			return 0, err
		}

		if min > val {
			min = val
		}
	}

	return min, nil
}

func max(packets []packet.Packet) (int, error) {
	max := math.MinInt

	if len(packets) == 0 {
		return 0, errors.New("unsupported operation: max of 0 packets")
	}

	for _, p := range packets {
		val, err := calculateResult(p)
		if err != nil {
			return 0, err
		}

		if max < val {
			max = val
		}
	}

	return max, nil
}

func gt(packets []packet.Packet) (int, error) {
	if len(packets) != 2 {
		return 0, fmt.Errorf("unsupported operation: gt of %v packets", len(packets))
	}

	first, err := calculateResult(packets[0])
	if err != nil {
		return 0, err
	}

	second, err := calculateResult(packets[1])
	if err != nil {
		return 0, err
	}

	if first > second {
		return 1, nil
	}

	return 0, nil
}

func lt(packets []packet.Packet) (int, error) {
	if len(packets) != 2 {
		return 0, fmt.Errorf("unsupported operation: lt of %v packets", len(packets))
	}

	first, err := calculateResult(packets[0])
	if err != nil {
		return 0, err
	}

	second, err := calculateResult(packets[1])
	if err != nil {
		return 0, err
	}

	if first < second {
		return 1, nil
	}

	return 0, nil
}

func eq(packets []packet.Packet) (int, error) {
	if len(packets) != 2 {
		return 0, fmt.Errorf("unsupported operation: eq of %v packets", len(packets))
	}

	first, err := calculateResult(packets[0])
	if err != nil {
		return 0, err
	}

	second, err := calculateResult(packets[1])
	if err != nil {
		return 0, err
	}

	if first == second {
		return 1, nil
	}

	return 0, nil
}
