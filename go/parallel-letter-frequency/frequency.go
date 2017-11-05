package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(strings []string) FreqMap {
	m := FreqMap{}
	c := make(chan FreqMap, len(strings))
	for _, s := range strings {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	for range strings {
		for k, v := range <-c {
			m[k] += v
		}
	}
	return m
}
