package algorithm

import "errors"

type Algorithm []byte

func FromString(str string) (Algorithm, error) {
	result := make([]byte, 0, len(str))

	for _, char := range []rune(str) {
		switch char {
		case '#':
			result = append(result, 1)
		case '.':
			result = append(result, 0)
		default:
			return nil, errors.New("unknown pixel type")
		}
	}

	return result, nil
}
