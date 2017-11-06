package bob

import (
	"regexp"
)

const testVersion = 3

func Hey(s string) string {
	rQuestion := regexp.MustCompile(`\?\s*$`)
	rYell := regexp.MustCompile(`^[A-Z]+$`)
	rFine := regexp.MustCompile(`^\W*$`)
	switch {
	case rYell.MatchString(regexp.MustCompile(`[^a-z|A-Z]+`).ReplaceAllString(s, "")):
		return "Whoa, chill out!"
	case rQuestion.MatchString(s):
		return "Sure."
	case rFine.MatchString(s):
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
