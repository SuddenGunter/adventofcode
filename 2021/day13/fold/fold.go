package fold

type Direction string

const (
	X Direction = "x"
	Y Direction = "y"
)

type Fold struct {
	Direction Direction
	Value     int
}
