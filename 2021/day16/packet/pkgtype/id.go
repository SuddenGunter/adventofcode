package pkgtype

type ID int

func IsLiteral(id ID) bool {
	return id == 4
}

func IsOperator(id ID) bool {
	return !IsLiteral(id)
}

const (
	Sum     ID = 0
	Product ID = 1
	Min     ID = 2
	Max     ID = 3
	Val     ID = 4
	Gt      ID = 5
	Lt      ID = 6
	Eq      ID = 7
)
