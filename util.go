package goaoi

import "golang.org/x/exp/constraints"

func min[T constraints.Ordered](x T, y T) T {
	if x < y {
		return x
	}
	return y
}
