package algorithm

import (
	"go-ds/utils/comparator"
)

func Max[T comparator.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T comparator.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Abs[T comparator.Numeric](a T) T {
	if a < T(0) {
		return -a
	}
	return a
}
