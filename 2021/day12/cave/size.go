package cave

import "strings"

type Size byte

const (
	Small Size = iota
	Big
)

func SizeByName(name string) Size {
	if strings.ToLower(name) == name {
		return Small
	}

	return Big
}
