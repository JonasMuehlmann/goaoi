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

func FindEndSlice[T comparable](super []T, sub []T, comparator func(T, T) bool) (int, error) {
	if len(super) == 0 || len(sub) == 0 {
		return 0, EmptyIterableError{}
	}
OUTER:
	for i := len(super) - 1; i >= len(sub)-1; i-- {
		for j := 0; j < len(sub); j++ {
			if !comparator(super[i-j], sub[len(sub)-1-j]) {
				continue OUTER
			}
		}
		return i - len(sub) + 1, nil
	}

	return 0, errors.New("Could not find element")
}

func FindFirstOfSlice[T comparable](haystack []T, needles []T, comparator func(T, T) bool) (int, error) {
	if len(haystack) == 0 || len(needles) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
			if comparator(haystackValue, needleValue) {
				return i, nil
			}
		}
	}

	return 0, errors.New("Could not find element")
}

func FindFirstOfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, needles []TValue, comparator func(TValue, TValue) bool) (TKey, error) {
	var zeroVal TKey
	if len(haystack) == 0 || len(needles) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
			if comparator(haystackValue, needleValue) {
				return i, nil
			}
		}
	}

	return zeroVal, errors.New("Could not find element")
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

func NoneOfSlice[T comparable](haystack []T, comparator func(T) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range haystack {
		if comparator(value) {
			return ComparisonError[int]{len(haystack) - 1}
		}
	}

	return nil
}

func NoneOfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range haystack {
		if comparator(value) {
			return ComparisonError[TKey]{}
		}
	}

	return nil
}

func ForeachSlice[T comparable](haystack []T, comparator func(T) error) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range haystack {
		err := comparator(value)
		if err != nil {
			return ExecutionError[int]{i, err}
		}
	}

	return nil
}

func ForeachMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) error) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range haystack {
		err := comparator(value)
		if err != nil {
			return ExecutionError[TKey]{key, err}
		}
	}

	return nil
}

func CountSlice[T comparable](haystack []T, wanted T) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range haystack {
		if value == wanted {
			counter++
		}
	}

	return counter, nil
}

func CountMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, wanted TValue) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range haystack {
		if value == wanted {
			counter++
		}
	}

	return counter, nil
}

func CountIfSlice[T comparable](haystack []T, comparator func(T) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range haystack {
		if comparator(value) {
			counter++
		}
	}

	return counter, nil
}

func CountIfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range haystack {
		if comparator(value) {
			counter++
		}
	}

	return counter, nil
}

func MismatchSlice[T comparable](iterable1 []T, iterable2 []T) (int, error) {
	if len(iterable1) == 0 || len(iterable2) == 0 {
		return 0, EmptyIterableError{}
	}

	i := 0
	for ; i < min(len(iterable1), len(iterable2)); i++ {
		if iterable1[i] != iterable2[i] {
			return i, nil
		}
	}

	return 0, EqualIteratorsError{}
}
