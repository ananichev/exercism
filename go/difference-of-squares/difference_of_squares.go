package diffsquares

const testVersion = 1

func SquareOfSums(n int) (r int) {
	for i := 1; i <= n; i++ {
		r += i
	}
	return r * r
}

func SumOfSquares(n int) (r int) {
	return (2*n*n*n + 3*n*n + n) / 6
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
