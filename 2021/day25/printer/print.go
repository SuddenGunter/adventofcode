package printer

import "fmt"

func PrettyPrint(runes [][]rune) {
	for i := range runes {
		fmt.Println(string(runes[i]))
	}
}
