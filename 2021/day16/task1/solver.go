package task1

import (
	"aoc-2021-day16/decoder"
	"aoc-2021-day16/packet"
	"aoc-2021-day16/parser"
)

func Solve(data string) (int, error) {
	vector, err := parser.Parse(data)
	if err != nil {
		return 0, err
	}

	result, _, err := decoder.Decode(vector)
	if err != nil {
		return 0, err
	}

	return SumVersions(result), nil
}

func SumVersions(p packet.Packet) int {
	sum := 0
	notChecked := stack{data: make([]packet.Packet, 0)}
	notChecked.push(p)

	for !notChecked.isEmpty() {
		current := notChecked.pop()
		asOp, ok := current.(packet.OpPacket)
		if ok {
			for _, x := range asOp.Subpackets {
				notChecked.push(x)
			}
		}

		sum += current.Version()
	}

	return sum
}
