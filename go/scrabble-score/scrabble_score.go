package scrabble

import (
	"strings"
	"unicode"
)

var matcher = map[string]int{"aeioulnrst": 1, "dg": 2, "bcmp": 3, "fhvwy": 4, "k": 5, "jx": 8, "qz": 10}

func Score(str string) (score int) {
	for _, l := range str {
		lower := unicode.ToLower(l)
		for s, v := range matcher {
			if strings.ContainsRune(s, lower) {
				score += v
				break
			}
		}
	}
	return
}
