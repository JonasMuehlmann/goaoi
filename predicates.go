package goaoi

func AreEqual[T comparable](a T, b T) bool {
	return a == b
}

func AreNotEqual[T comparable](a T, b T) bool {
	return a != b
}

func IsLessThan[T comparable](a T, b T) bool {
	return a < b
}

func IsLessThanEqual[T comparable](a T, b T) bool {
	return a <= b
}

func IsGreaterThan[T comparable](a T, b T) bool {
	return a > b
}

func IsGreaterThanEqual[T comparable](a T, b T) bool {
	return a >= b
}

func IsNull[T comparable](a T) bool {
	return a == nil
}

func IsNotNull[T comparable](a T) bool {
	return a != nil
}
