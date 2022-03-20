package goaoi

import "errors"

func Find[T comparable](haystack []T, needle T) (int, error) {
	for i, value := range haystack {
		if value == needle {
			return i, nil
		}
	}

	return 0, errors.New("Could not find element")
}
