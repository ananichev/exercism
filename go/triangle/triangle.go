package triangle

import . "math"

const testVersion = 3

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func KindFromSides(a, b, c float64) (k Kind) {
	switch {
	case NotATriangle(a, b, c):
		break
	case EquTriangle(a, b, c):
		k = Equ
		break
	case IsoTriangle(a, b, c):
		k = Iso
		break
	case ScaTriangle(a, b, c):
		k = Sca
		break
	}
	return
}

func EquTriangle(a, b, c float64) bool {
	return a == b && a == c
}

func IsoTriangle(a, b, c float64) bool {
	return a == b || a == c || b == c
}

func ScaTriangle(a, b, c float64) bool {
	return a != b && a != c
}

func NotATriangle(a, b, c float64) bool {
	return IsNaN(a) || IsNaN(b) || IsNaN(c) || IsInf(a, 0) || IsInf(b, 0) || IsInf(c, 0) ||
		a <= 0 || b <= 0 || c <= 0 || (a+b < c) || (a+c < b) || (b+c < a)
}
