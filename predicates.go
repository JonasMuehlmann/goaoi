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

func NegateUnaryPredicate[T any](f func(T) bool) func(T) bool {
	return func(t T) bool { return !f(t) }
}

func NegateBinaryPredicate[T any](f func(T, T) bool) func(T, T) bool {
	return func(t1 T, t2 T) bool { return !f(t1, t2) }
}
