package amphipod

import "strings"

const (
	HallSize = 11
	Rooms    = 4
	RoomSize = 2
)

type Burrow struct {
	Hall  [HallSize]rune
	Rooms [Rooms][RoomSize]rune
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
