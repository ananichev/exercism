package pangram

import (
	"regexp"
	"strings"
)

const (
	testVersion = 2
	alphabetLen = 26
)

func IsPangram(s string) bool {
	lettersMap := map[rune]bool{}
	for _, l := range regexp.MustCompile(`[^a-z|A-Z]+`).ReplaceAllString(strings.ToLower(s), "") {
		lettersMap[l] = true
	}
	return len(lettersMap) == alphabetLen
}
