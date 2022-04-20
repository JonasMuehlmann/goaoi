package goaoi

import "golang.org/x/exp/constraints"

func AreEqual[T comparable](a T, b T) bool {
	return a == b
}

func AreEqualPartial[T comparable](fixedComparee T) func(T) bool {
	return func(comparee T) bool { return comparee == fixedComparee }
}

func AreNotEqual[T comparable](a T, b T) bool {
	return a != b
}

func AreNotEqualPartial[T comparable](fixedComparee T) func(T) bool {
	return func(comparee T) bool { return comparee != fixedComparee }
}

func IsLessThan[T constraints.Ordered](a T, b T) bool {
	return a < b
}

func IsLessThanPartial[T constraints.Ordered](fixedComparee T) func(T) bool {
	return func(comparee T) bool { return comparee < fixedComparee }
}

func IsLessThanEqual[T constraints.Ordered](a T, b T) bool {
	return a <= b
}

func IsLessThanEqualPartial[T constraints.Ordered](fixedComparee T) func(T) bool {
	return func(comparee T) bool { return comparee <= fixedComparee }
}

func IsGreaterThan[T constraints.Ordered](a T, b T) bool {
	return a > b
}

func IsGreaterThanPartial[T constraints.Ordered](fixedComparee T) func(T) bool {
	return func(comparee T) bool { return comparee < fixedComparee }
}

func IsGreaterThanEqual[T constraints.Ordered](a T, b T) bool {
	return a >= b
}

func IsGreaterThanEqualPartial[T constraints.Ordered](fixedComparee T) func(T) bool {
	return func(comparee T) bool { return comparee >= fixedComparee }
}
