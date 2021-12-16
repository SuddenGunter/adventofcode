package lentype

type ID int

func IsLenInBits(id ID) bool {
	return id == 0
}

func IsNumOfSubpackets(id ID) bool {
	return id == 1
}

const (
	BitsForLenInBits    = 15
	BitsNumOfSubpackets = 11
)
