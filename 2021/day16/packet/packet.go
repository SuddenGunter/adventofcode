package packet

import (
	"aoc-2021-day16/packet/pkgtype"

	"github.com/dropbox/godropbox/container/bitvector"
)

type Packet struct {
	Version int
	TypeID  pkgtype.ID

	// without version and type
	Body *bitvector.BitVector
}
