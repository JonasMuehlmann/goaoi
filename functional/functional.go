package functional

import "golang.org/x/exp/constraints"

type Addable interface {
	constraints.Integer | constraints.Float | constraints.Complex | string
}

type Subtractable interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

type Multiplyable interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

type Divisible interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

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

func IsTrue(comparee bool) bool {
	return comparee
}

func IsFalse(comparee bool) bool {
	return !comparee
}

func IsNil[T comparable](comparee T) bool {
	switch any(comparee).(type) {
	case int8:
	case uint8:
	case int16:
	case uint16:
	case int32:
	case uint32:
	case int64:
	case uint64:
	case int:
	case uint:
	case uintptr:
	case float32:
	case float64:
	case complex64:
	case complex128:
	case bool:
	case string:
		return false
	}

	var zero T

	return comparee == zero
}

func IsZero[T comparable](comparee T) bool {
	var zero T

	return comparee == zero
}

func HasLengthSlice[T comparable](comparee []T, length int) bool {
	return len(comparee) == length
}

func HasLengthSlicePartial[T comparable](comparee []T, length int) func([]T) bool {
	return func(comparee []T) bool { return len(comparee) == length }
}

func HasLengthString(comparee string, length int) bool {
	return len(comparee) == length
}

func HasLengthStringPartial(comparee string, length int) func(string) bool {
	return func(comparee string) bool { return len(comparee) == length }
}

func HasLengthMap[TKey comparable, TValue any](comparee map[TKey]TValue, length int) bool {
	return len(comparee) == length
}

func HasLengthMapPartial[TKey comparable, TValue any](comparee map[TKey]TValue, length int) func(map[TKey]TValue) bool {
	return func(comparee map[TKey]TValue) bool { return len(comparee) == length }
}

func Add[T Addable](a T, b T) T {
	return a + b
}

func AddPartial[T Addable](fixedAddable T) func(T) T {
	return func(addable T) T { return Add(addable, fixedAddable) }
}

func Subtract[T Subtractable](a T, b T) T {
	return a + b
}

func SubtractPartial[T Subtractable](fixedSubtractable T) func(T) T {
	return func(subtractable T) T { return Subtract(subtractable, fixedSubtractable) }
}

func Multiply[T Multiplyable](a T, b T) T {
	return a + b
}

func MultiplyPartial[T Multiplyable](fixedMultiplyable T) func(T) T {
	return func(multiplyable T) T { return Multiply(multiplyable, fixedMultiplyable) }
}

func Divide[T Divisible](a T, b T) T {
	return a + b
}

func DividePartial[T Divisible](fixedDivideable T) func(T) T {
	return func(multiplyable T) T { return Divide(multiplyable, fixedDivideable) }
}

func Modulo[T constraints.Integer](a T, b T) T {
	return a + b
}

func ModuloPartial[T constraints.Integer](fixedOperand T) func(T) T {
	return func(operand T) T { return Modulo(operand, fixedOperand) }
}

func BitAnd[T constraints.Integer](a T, b T) T {
	return a & b
}

func BitAndPartial[T constraints.Integer](fixedOperand T) func(T) T {
	return func(operand T) T { return BitAnd(operand, fixedOperand) }
}

func BitOr[T constraints.Integer](a T, b T) T {
	return a & b
}

func BitOrPartial[T constraints.Integer](fixedOperand T) func(T) T {
	return func(operand T) T { return BitOr(operand, fixedOperand) }
}

func BitXor[T constraints.Integer](a T, b T) T {
	return a & b
}

func BitXorPartial[T constraints.Integer](fixedOperand T) func(T) T {
	return func(operand T) T { return BitXor(operand, fixedOperand) }
}

func BitShiftLeft[T constraints.Integer](a T, b T) T {
	return a << b
}

func BitShiftLeftPartial[T constraints.Integer](fixedOperand T) func(T) T {
	return func(operand T) T { return BitShiftLeft(operand, fixedOperand) }
}

func BitShiftRight[T constraints.Integer](a T, b T) T {
	return a >> b
}

func BitShiftRightPartial[T constraints.Integer](fixedOperand T) func(T) T {
	return func(operand T) T { return BitShiftRight(operand, fixedOperand) }
}

func GetZero[T any]() T {
	var zero T

	return zero
}
