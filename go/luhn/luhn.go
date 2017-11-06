package luhn

import (
	"regexp"
	"strconv"
)

func Valid(str string) bool {
	var acc int
	str = regexp.MustCompile(`(^0?|\s*)`).ReplaceAllString(str, "")

	if len(str) == 0 {
		return false
	}

	for i, n := range str {
		d, err := strconv.Atoi(string(n))
		if err != nil {
			return false
		}
		if i%2 == 0 {
			d = d * 2
			if d > 9 {
				d -= 9
			}
		}
		acc += d
	}
	return acc%10 == 0
}
