package goaoi

func min[T comparable](x T, y T) T {
	if x < y {
		return x
	}
	return y
}
