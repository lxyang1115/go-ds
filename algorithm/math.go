package algorithm

func max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func abs[T Numeric](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
