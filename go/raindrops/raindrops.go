package raindrops

import "strconv"

const testVersion = 3

var (
	factors    = [...]int{3, 5, 7}
	factorsMap = map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
)

func Convert(i int) (c string) {
	for _, f := range factors {
		if i%f == 0 {
			c += factorsMap[f]
		}
	}
	if len(c) == 0 {
		c = strconv.Itoa(i)
	}
	return
}
