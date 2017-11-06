package accumulate

const testVersion = 1

func Accumulate(strings []string, f func(string) string) (result []string) {
	for _, s := range strings {
		result = append(result, f(s))
	}
	return
}
