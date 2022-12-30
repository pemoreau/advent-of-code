package utils

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	res := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		res = LCM(res, integers[i])
	}
	return res
}
