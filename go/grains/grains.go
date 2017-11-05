package grains

import (
	"errors"
)

const squaresCount = 64

func Square(n int) (uint64, error) {
	if n < 1 || n > squaresCount {
		return 0, errors.New("Square number is not valid")
	}
	return 1 << uint(n-1), nil
}

func Total() uint64 {
	return (1 << squaresCount) - 1
}
