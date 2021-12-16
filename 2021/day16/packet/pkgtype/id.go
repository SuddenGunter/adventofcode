package pkgtype

type ID int

func IsLiteral(id ID) bool {
	return id == 4
}

func IsOperator(id ID) bool {
	return !IsLiteral(id)
}
