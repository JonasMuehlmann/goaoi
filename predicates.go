package goaoi

import "golang.org/x/exp/constraints"

func AreEqual[T comparable](a T, b T) bool {
	return a == b
}

func AreNotEqual[T comparable](a T, b T) bool {
	return a != b
}

func IsLessThan[T constraints.Ordered](a T, b T) bool {
	return a < b
}

func IsLessThanEqual[T constraints.Ordered](a T, b T) bool {
	return a <= b
}

func IsGreaterThan[T constraints.Ordered](a T, b T) bool {
	return a > b
}

func IsGreaterThanEqual[T constraints.Ordered](a T, b T) bool {
	return a >= b
}
