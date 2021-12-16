package parser

import (
	"encoding/hex"

	"github.com/dropbox/godropbox/container/bitvector"
)

func Parse(data string) (*bitvector.BitVector, error) {
	hexdec, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}

	vector := make([]byte, 0)
	for _, char := range hexdec {
		bits := make([]byte, 8)
		for i := 0; i <= 7; i++ {
			if ((char >> (7 - i)) & 1) == 1 {
				bits[i] = 1
			} else {
				bits[i] = 0
			}
		}

		vector = append(vector, bits...)
	}

	v := bitvector.NewBitVector(make([]byte, 0), 0)
	for i := range vector {
		v.Append(vector[i])
	}

	return v, nil
}

//
