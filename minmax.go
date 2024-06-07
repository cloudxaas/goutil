package util

func Min[T comparable](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

