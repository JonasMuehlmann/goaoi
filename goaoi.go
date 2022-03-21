package goaoi

import "errors"

func FindSlice[T comparable](haystack []T, needle T) (int, error) {
	for i, value := range haystack {
		if value == needle {
			return i, nil
		}
	}

	return 0, errors.New("Could not find element")
}

func FindIfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) (TKey, error) {
	for key, value := range haystack {
		if comparator(value) {
			return key, nil
		}
	}

	var zeroVal TKey
	return zeroVal, errors.New("Could not find element")
}

func FindIfSlice[T comparable](haystack []T, comparator func(T) bool) (int, error) {
	for i, value := range haystack {
		if comparator(value) {
			return i, nil
		}
	}

	return 0, errors.New("Could not find element")
}
