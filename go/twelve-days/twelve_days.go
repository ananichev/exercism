package twelve

import (
	"fmt"
	"strings"
)

const testVersion = 1

var (
	ordinal   = [12]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
	songParts = [12]string{
		"a Partridge in a Pear Tree",
		"two Turtle Doves",
		"three French Hens",
		"four Calling Birds",
		"five Gold Rings",
		"six Geese-a-Laying",
		"seven Swans-a-Swimming",
		"eight Maids-a-Milking",
		"nine Ladies Dancing",
		"ten Lords-a-Leaping",
		"eleven Pipers Piping",
		"twelve Drummers Drumming",
	}
)

func Song() string {
	var parts []string
	for i := 1; i <= len(songParts); i++ {
		parts = append(parts, Verse(i))
	}
	return strings.Join(parts, "\n") + "\n"
}

func Verse(r int) (output string) {
	str := []string{fmt.Sprintf("On the %s day of Christmas my true love gave to me", ordinal[r-1])}
	for i := r - 1; i >= 0; i-- {
		str = append(str, songParts[i])
	}

	return toSentence(str)
}

func toSentence(s []string) string {
	if len(s) > 2 {
		return strings.Join(s[:len(s)-1], ", ") + ", and " + s[len(s)-1] + "."
	}
	return strings.Join(s, ", ") + "."
}
