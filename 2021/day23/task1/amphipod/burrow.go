package amphipod

import "strings"

const (
	HallSize = 11
	Rooms    = 4
	RoomSize = 2
)

func EnterableHallCell(i int) bool {
	return !(i%2 == 0 && i != 0 && i != HallSize-1)
}

func RoomByName(name rune) int {
	switch name {
	case 'A':
		return 0
	case 'B':
		return 1
	case 'C':
		return 2
	case 'D':
		return 3
	default:
		return -1
	}
}

func NameByRoom(room int) rune {
	return 'A' + rune(room)
}

type Burrow struct {
	Hall  [HallSize]rune
	Rooms [Rooms][RoomSize]rune
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

var Cost = map[rune]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}
