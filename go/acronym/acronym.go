package acronym

import (
	"bytes"
	"regexp"
	"strings"
)

const testVersion = 3

func Abbreviate(s string) string {
	var b bytes.Buffer
	for _, w := range regexp.MustCompile(`\W+`).Split(s, -1) {
		b.WriteString(w[:1])
	}
	return strings.ToUpper(b.String())
}
