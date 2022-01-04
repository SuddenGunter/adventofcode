package amphipod

import (
	"math"
	"strings"
)

const (
	HallSize = 11
	Rooms    = 4
	RoomSize = 2
)

func EnterableHallCell(i int) bool {
	return !(i%2 == 0 && i != 0 && i != HallSize-1)
}

func RoomByName(name rune) int {
	return int(name - 'A')
}

func NameByRoom(room int) rune {
	return 'A' + rune(room)
}

type Burrow struct {
	Hall  [HallSize]rune
	Rooms [Rooms][RoomSize]rune
}

func (b Burrow) HallEmpty() bool {
	for _, r := range b.Hall {
		if r != '.' {
			return false
		}
	}

	return true
}

func (b Burrow) Clone() Burrow {
	var hall [HallSize]rune
	for i := range b.Hall {
		hall[i] = b.Hall[i]
	}

	var rooms [Rooms][RoomSize]rune
	for i := range b.Rooms {
		for j := range b.Rooms[i] {
			rooms[i][j] = b.Rooms[i][j]
		}
	}

	return Burrow{
		Hall:  hall,
		Rooms: rooms,
	}
}

func (b Burrow) String() string {
	sb := strings.Builder{}
	for i := 0; i < HallSize+2; i++ {
		sb.WriteRune('#')
	}
	sb.WriteRune('\n')

	sb.WriteRune('#')
	for _, r := range b.Hall {
		sb.WriteRune(r)
	}
	sb.WriteRune('#')
	sb.WriteRune('\n')

	for i := 0; i < RoomSize; i++ {
		sb.WriteRune(' ')
		sb.WriteRune(' ')
		sb.WriteRune('#')

		for j := 0; j < Rooms; j++ {
			sb.WriteRune(b.Rooms[j][i])
			sb.WriteRune('#')
		}

		sb.WriteRune(' ')
		sb.WriteRune(' ')
		sb.WriteRune('\n')
	}

	sb.WriteRune(' ')
	sb.WriteRune(' ')
	for i := 0; i < HallSize-2; i++ {
		sb.WriteRune('#')
	}

	sb.WriteRune('\n')

	return sb.String()
}

func Cost(name rune) int {
	return int(math.Pow10(RoomByName(name)))
}
