package goaoi

import "errors"

func FindSlice[T comparable](haystack []T, needle T) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, value := range haystack {
		if value == needle {
			return i, nil
		}
	}

	return 0, errors.New("Could not find element")
}

func FindIfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) (TKey, error) {
	var zeroVal TKey

	if len(haystack) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	for key, value := range haystack {
		if comparator(value) {
			return key, nil
		}
	}

	return zeroVal, errors.New("Could not find element")
}

func FindIfSlice[T comparable](haystack []T, comparator func(T) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, value := range haystack {
		if comparator(value) {
			return i, nil
		}
	}

	return 0, errors.New("Could not find element")
}

func AllOfSlice[T comparable](haystack []T, comparator func(T) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range haystack {
		if !comparator(value) {
			return ComparisonError[int]{i}
		}
	}

	return nil
}

func AllOfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range haystack {
		if !comparator(value) {
			return ComparisonError[TKey]{key}
		}
	}

	return nil
}

func AnyOfSlice[T comparable](haystack []T, comparator func(T) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range haystack {
		if comparator(value) {
			return nil
		}
	}

	return ComparisonError[int]{len(haystack) - 1}
}

func AnyOfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range haystack {
		if comparator(value) {
			return nil
		}
	}

	return ComparisonError[TKey]{}
}
