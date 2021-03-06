// Package goaoi implements conventient algorithms for processing iterables.
// It is inspired by the algorithm header from the C++ standard template library (STL for short).
package goaoi

import (
	"math"

	"golang.org/x/exp/constraints"
)

// FindSlice finds the first index i where haystack[i] == needle.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindSlice[T comparable](haystack []T, needle T) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, value := range haystack {
		if value == needle {
			return i, nil
		}
	}

	return 0, ElementNotFoundError{}
}

// FindIfMap finds the first key where unary_predicate(haystack[key]) == true.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindIfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, unary_predicate func(TValue) bool) (TKey, error) {
	var zeroVal TKey

	if len(haystack) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	for key, value := range haystack {
		if unary_predicate(value) {
			return key, nil
		}
	}

	return zeroVal, ElementNotFoundError{}
}

// FindIfSlice finds the first index i where unary_predicate(haystack[i]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindIfSlice[T comparable](haystack []T, unary_predicate func(T) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, value := range haystack {
		if unary_predicate(value) {
			return i, nil
		}
	}

	return 0, ElementNotFoundError{}
}

// FindEndSlicePred finds the beginning of the last occurrence of sub in super.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindEndSlicePred[T comparable](super []T, sub []T, binary_predicate func(T, T) bool) (int, error) {
	if len(super) == 0 || len(sub) == 0 {
		return 0, EmptyIterableError{}
	}
OUTER:
	for i := len(super) - 1; i >= len(sub)-1; i-- {
		for j := 0; j < len(sub); j++ {
			if !binary_predicate(super[i-j], sub[len(sub)-1-j]) {
				continue OUTER
			}
		}
		return i - len(sub) + 1, nil
	}

	return 0, ElementNotFoundError{}
}

// FindEndSlice finds the beginning of the last occurrence of sub in super.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindEndSlice[T comparable](super []T, sub []T) (int, error) {
	if len(super) == 0 || len(sub) == 0 {
		return 0, EmptyIterableError{}
	}
OUTER:
	for i := len(super) - 1; i >= len(sub)-1; i-- {
		for j := 0; j < len(sub); j++ {
			if super[i-j] != sub[len(sub)-1-j] {
				continue OUTER
			}
		}
		return i - len(sub) + 1, nil
	}

	return 0, ElementNotFoundError{}
}

// FindFirstOfSlicePred finds the first index where an element of haystack is equal to any element in needles.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindFirstOfSlicePred[T comparable](haystack []T, needles []T, binary_predicate func(T, T) bool) (int, error) {
	if len(haystack) == 0 || len(needles) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
			if binary_predicate(haystackValue, needleValue) {
				return i, nil
			}
		}
	}

	return 0, ElementNotFoundError{}
}

// FindFirstOfMapPred finds the first key where an element of haystack is equal to any element in needles.
// Note that the iteration order of a map is not stable.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindFirstOfMapPred[TKey comparable, TValue comparable](haystack map[TKey]TValue, needles []TValue, binary_predicate func(TValue, TValue) bool) (TKey, error) {
	var zeroVal TKey
	if len(haystack) == 0 || len(needles) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
			if binary_predicate(haystackValue, needleValue) {
				return i, nil
			}
		}
	}

	return zeroVal, ElementNotFoundError{}
}

// FindFirstOfSlice finds the first index where an element of haystack is equal to any element in needles.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindFirstOfSlice[T comparable](haystack []T, needles []T) (int, error) {
	if len(haystack) == 0 || len(needles) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
			if haystackValue == needleValue {
				return i, nil
			}
		}
	}

	return 0, ElementNotFoundError{}
}

// FindFirstOfMap finds the first key where an element of haystack is equal to any element in needles.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindFirstOfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, needles []TValue) (TKey, error) {
	var zeroVal TKey
	if len(haystack) == 0 || len(needles) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
			if haystackValue == needleValue {
				return i, nil
			}
		}
	}

	return zeroVal, ElementNotFoundError{}
}

// AllOfSlice checks that unary_predicate(val) == true for ALL val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func AllOfSlice[T any](container []T, unary_predicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		if !unary_predicate(value) {
			return ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
		}
	}

	return nil
}

// AllOfMap checks that unary_predicate(val) == true for ALL val in container.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func AllOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range container {
		if !unary_predicate(value) {
			return ComparisonError[TKey, TValue]{BadItemIndex: key, BadItem: value}
		}
	}

	return nil
}

// AnyOfSlice checks that unary_predicate(val) == true for ANY val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func AnyOfSlice[T any](container []T, unary_predicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unary_predicate(value) {
			return nil
		}
	}

	return ElementNotFoundError{}
}

// AnyOfMap checks that unary_predicate(val) == true for ANY val in container.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func AnyOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unary_predicate(value) {
			return nil
		}
	}

	return ElementNotFoundError{}
}

// NoneOfSlice checks that unary_predicate(val) == true for NO val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func NoneOfSlice[T any](container []T, unary_predicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		if unary_predicate(value) {
			return ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
		}
	}

	return nil
}

// NoneOfMap checks that unary_predicate(val) == true for NO val in container.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func NoneOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range container {
		if unary_predicate(value) {
			return ComparisonError[TKey, TValue]{BadItemIndex: key, BadItem: value}
		}
	}

	return nil
}

// ForeachSlice executes unary_func(val) for each val in container.
// Errors returned by unary_func are propagated to the caller of ForeachSlice.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func ForeachSlice[T any](container []T, unary_func func(T) error) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		err := unary_func(value)
		if err != nil {
			return ExecutionError[int, T]{BadItemIndex: i, BadItem: container[i], Inner: err}
		}
	}

	return nil
}

// ForeachMap executes unary_func(val) for each val in container.
// Note that the iteration order of a map is not stable.
// Errors returned by unary_func are propagated to the caller of ForeachMap.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func ForeachMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_func func(TValue) error) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range container {
		err := unary_func(value)
		if err != nil {
			return ExecutionError[TKey, TValue]{BadItemIndex: key, BadItem: value, Inner: err}
		}
	}

	return nil
}

// ForeachSliceUnsafe executes unary_func(val) for each val in container.
//
// Possible Error values:
//    - EmptyIterableError
func ForeachSliceUnsafe[T any](container []T, unary_func func(T)) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		unary_func(value)
	}

	return nil
}

// ForeachMapUnsafe executes unary_func(val) for each val in container.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func ForeachMapUnsafe[TKey comparable, TValue comparable](container map[TKey]TValue, unary_func func(TValue)) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		unary_func(value)
	}

	return nil
}

// CountSlice counts how many elements of container are equal to wanted.
//
// Possible Error values:
//    - EmptyIterableError
func CountSlice[T comparable](container []T, wanted T) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if value == wanted {
			counter++
		}
	}

	return counter, nil
}

// CountMap counts how many elements of container are equal to wanted.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CountMap[TKey comparable, TValue comparable](container map[TKey]TValue, wanted TValue) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if value == wanted {
			counter++
		}
	}

	return counter, nil
}

// CountIfSlice counts for how many val of container unary_predicate(val) == true.
//
// Possible Error values:
//    - EmptyIterableError
func CountIfSlice[T comparable](container []T, unary_predicate func(T) bool) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if unary_predicate(value) {
			counter++
		}
	}

	return counter, nil
}

// CountIfMap counts for how many val of container unary_predicate(val) == true.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CountIfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if unary_predicate(value) {
			counter++
		}
	}

	return counter, nil
}

// MismatchSlicePred finds the first index i where binary_predicate(iterable1[i], iterable2[i] == false).
//
// Possible Error values:
//    - EmptyIterableError
//    - EqualIteratorsError
func MismatchSlicePred[T comparable](iterable1 []T, iterable2 []T, binary_predicate func(T, T) bool) (int, error) {
	if len(iterable1) == 0 || len(iterable2) == 0 {
		return 0, EmptyIterableError{}
	}

	i := 0
	for ; i < min(len(iterable1), len(iterable2)); i++ {
		if !binary_predicate(iterable1[i], iterable2[i]) {
			return i, nil
		}
	}

	return 0, EqualIteratorsError{}
}

// MismatchSlice finds the first index i where iterable1[i] != iterable2[i].
//
// Possible Error values:
//    - EmptyIterableError
//    - EqualIteratorsError
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

// AdjacentFindSlice finds the first index i where container[i] == container[i+1]).
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func AdjacentFindSlice[T comparable](container []T) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	for i := 0; i < len(container)-1; i++ {
		if container[i] == container[i+1] {
			return i, nil
		}
	}

	return 0, ElementNotFoundError{}
}

// AdjacentFindSlicePred finds the first index i where binary_predicate(container[i], container[i+1]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func AdjacentFindSlicePred[T comparable](container []T, binary_predicate func(T, T) bool) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	for i := 0; i < len(container)-1; i++ {
		if binary_predicate(container[i], container[i+1]) {
			return i, nil
		}
	}

	return 0, ElementNotFoundError{}
}

// TakeWhileSlice returns a copy of original until the first element not satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func TakeWhileSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if !unary_predicate(value) {
			return newContainer, nil
		}

		newContainer = append(newContainer, value)
	}

	return newContainer, nil
}

// TakeWhileMap returns a copy of original until the first value not satisfying unary_predicate(value) == true).
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func TakeWhileMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if !unary_predicate(value) {
			return newContainer, nil
		}

		newContainer[key] = value
	}

	return newContainer, nil
}

// DropWhileSlice returns a copy of original starting from first element not satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func DropWhileSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	i := 0
	for _, value := range original {
		if !unary_predicate(value) {
			break
		}

		i++
	}

	for _, value := range original[i:] {
		newContainer = append(newContainer, value)
	}

	return newContainer, nil
}

// DropWhileMap returns a copy of original starting from the first value not satisfying unary_predicate(value) == true).
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func DropWhileMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	isDropping := true
	for key, value := range original {
		if !(unary_predicate(value) && isDropping) {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

// CopyIfSlice returns a copy of original with all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyIfSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if unary_predicate(value) {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

// CopyIfMap returns a copy of original with all key-value pairs satisfying unary_predicate(value) == true).
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CopyIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if unary_predicate(value) {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

// CopyReplaceSlice returns a copy of original where each element equal to toReplace is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceSlice[T comparable](original []T, toReplace T, replacement T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if value == toReplace {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

// CopyReplaceMap returns a copy of original where each value of a key-value pair equal to toReplace is replaced with replacement.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceMap[TKey comparable, TValue comparable](original map[TKey]TValue, toReplace TValue, replacement TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if value == toReplace {
			newContainer[key] = replacement
		} else {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

// CopyReplaceIfSlice returns a copy of original where each element satisfying unary_predicate(element) == true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfSlice[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if unary_predicate(value) {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

// CopyReplaceIfMap returns a copy of original where each value of a key-value pair satisfying unary_predicate(value) == true is replaced with replacement.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool, replacement TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if unary_predicate(value) {
			newContainer[key] = replacement
		} else {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

// CopyReplaceIfNotSlice returns a copy of original where each element satisfying unary_predicate(element) != true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfNotSlice[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if !unary_predicate(value) {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

// CopyReplaceIfNotMap returns a copy of original where each element satisfying unary_predicate(element) != true is replaced with replacement.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfNotMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool, replacement TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if !unary_predicate(value) {
			newContainer[key] = replacement
		} else {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

// CopyExceptSlice returns a copy of original without all elements equal to toExclude.
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptSlice[T comparable](original []T, toExclude T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if value != toExclude {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

// CopyExceptMap returns a copy of original without all key-value pairs equal to toExclude.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptMap[TKey comparable, TValue comparable](original map[TKey]TValue, toExclude TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if value != toExclude {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

// CopyExceptIfSlice returns a copy of original without all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if !unary_predicate(value) {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

// CopyExceptIfMap returns a copy of original without all key-value pairs satisfying unary_predicate(value) == true).
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if !unary_predicate(value) {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

// CopyExceptIfNotSlice returns a copy of original without all element satisfying unary_predicate(element) == false).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfNotSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if unary_predicate(value) {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

// CopyExceptIfNotMap returns a copy of original without all key-value pairs satisfying unary_predicate(value) == false).
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfNotMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if unary_predicate(value) {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

// FillSlice fills the array pointed to by arr with filler.
// all indices in the range [0, cap(*arr)[ are filled regardless of what len(*arr) is.
func FillSlice[T any](arr *[]T, filler T) []T {
	for i := range *arr {
		(*arr)[i] = filler
	}

	n_unfilled := cap(*arr) - len(*arr)

	for i := 0; i < n_unfilled; i++ {
		*arr = append(*arr, filler)
	}

	return *arr
}

// TransformMap applies transformer(value) for all key-value pairs in container and stores them at container[key].
// Note that the iteration order of a map is not stable.
// Errors returned by transformer are propagated to the caller of TransformMap.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformMap[TKey comparable, TValue comparable](container map[TKey]TValue, transformer func(TValue) (TValue, error)) error {

	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key := range container {
		newValue, err := transformer(container[key])
		if err != nil {
			return ExecutionError[TKey, TValue]{BadItemIndex: key, BadItem: container[key], Inner: err}
		}

		container[key] = newValue
	}

	return nil
}

// TransformSlice applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
// Errors returned by transformer are propagated to the caller of TransformSlice.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformSlice[T any](container []T, transformer func(*T) error) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		err := transformer(&container[i])
		if err != nil {
			return ExecutionError[int, T]{BadItemIndex: i, BadItem: value, Inner: err}
		}
	}

	return nil
}

// TransformMapUnsafe applies transformer(value) for all key-value pairs in container and stores them at container[key].
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func TransformMapUnsafe[TKey comparable, TValue comparable](container map[TKey]TValue, transformer func(TValue) TValue) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key := range container {
		container[key] = transformer(container[key])
	}

	return nil
}

// TransformSliceUnsafe applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
//
// Possible Error values:
//    - EmptyIterableError
func TransformSliceUnsafe[T any](container []T, transformer func(*T)) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i := range container {
		transformer(&container[i])
	}

	return nil
}

// TransformCopyMap applies transformer(value) for all key-value pairs in container and and returns the newly created container.
// Note that the iteration order of a map is not stable.
// Note that the transformer can return a different type than it's input.
// Errors returned by transformer are propagated to the caller of TransformCopyMap.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformCopyMap[TKey comparable, TValue comparable, TValueOut any](container map[TKey]TValue, transformer func(TValue) (TValueOut, error)) (map[TKey]TValueOut, error) {
	res := make(map[TKey]TValueOut)

	if len(container) == 0 {
		return res, EmptyIterableError{}
	}

	for key := range container {
		newValue, err := transformer(container[key])
		if err != nil {
			return res, ExecutionError[TKey, TValue]{BadItemIndex: key, BadItem: container[key], Inner: err}
		}

		res[key] = newValue
	}

	return res, nil
}

// TransformCopySlice applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
// Errors returned by transformer are propagated to the caller of TransformCopySlice.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformCopySlice[T any, TOut any](container []T, transformer func(T) (TOut, error)) ([]TOut, error) {
	res := make([]TOut, 0, len(container))

	if len(container) == 0 {
		return res, EmptyIterableError{}
	}

	for i, value := range container {
		newVal, err := transformer(container[i])
		if err != nil {
			return res, ExecutionError[int, T]{BadItemIndex: i, BadItem: value, Inner: err}
		}

		res = append(res, newVal)
	}

	return res, nil
}

// TransformCopyMapUnsafe applies transformer(value) for all key-value pairs in container and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func TransformCopyMapUnsafe[TKey comparable, TValue comparable, TValueOut any](container map[TKey]TValue, transformer func(TValue) TValueOut) (map[TKey]TValueOut, error) {
	res := make(map[TKey]TValueOut)

	if len(container) == 0 {
		return res, EmptyIterableError{}
	}

	for key := range container {
		res[key] = transformer(container[key])
	}

	return res, nil
}

// TransformCopySliceUnsafe applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
//
// Possible Error values:
//    - EmptyIterableError
func TransformCopySliceUnsafe[T any, TOut any](container []T, transformer func(T) TOut) ([]TOut, error) {
	res := make([]TOut, 0, len(container))

	if len(container) == 0 {
		return res, EmptyIterableError{}
	}

	for i := range container {
		res = append(res, transformer(container[i]))
	}

	return res, nil
}

// MinSliceInt finds the smallest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MinSliceInt[T constraints.Integer](haystack []T) (T, error) {
	var min T

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	min = haystack[0]
	for _, val := range haystack {
		if val < min {
			min = val
		}
	}

	return min, nil
}

// MinSliceFloat finds the smallest value in haystack.
// This funnction uses math.Min for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MinSliceFloat[T constraints.Float](haystack []T) (T, error) {
	var min T

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	min = haystack[0]
	for _, val := range haystack {
		min = T(math.Min(float64(min), float64(val)))
	}

	return min, nil
}

// MinSlicePred finds the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinSlicePred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {
	var min T

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	min = haystack[0]
	for _, val := range haystack {
		if binary_predicate(val, min) {
			min = val
		}
	}

	return min, nil
}

// MinMapInt finds the smallest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MinMapInt[TKey comparable, TValue constraints.Integer](haystack map[TKey]TValue) (TValue, error) {
	var min TValue

	for _, val := range haystack {
		min = val
		break
	}

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	for _, val := range haystack {
		if val < min {
			min = val
		}
	}

	return min, nil
}

// MinMapFloat finds the smallest value in haystack.
// This funnction uses math.Min for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MinMapFloat[TKey comparable, TValue constraints.Float](haystack map[TKey]TValue) (TValue, error) {
	var min TValue

	for _, val := range haystack {
		min = val
		break
	}

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	for _, val := range haystack {
		min = TValue(math.Min(float64(min), float64(val)))
	}

	return min, nil
}

// MinMapPred finds the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate func(TValue, TValue) bool) (TValue, error) {
	var min TValue

	for _, val := range haystack {
		min = val
		break
	}

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	for _, val := range haystack {
		if binary_predicate(val, min) {
			min = val
		}
	}

	return min, nil
}

// MaxSliceInt finds the largest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MaxSliceInt[T constraints.Integer](haystack []T) (T, error) {
	var max T

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	max = haystack[0]
	for _, val := range haystack {
		if val > max {
			max = val
		}
	}

	return max, nil
}

// MaxSliceFloat finds the largest value in haystack.
// This funnction uses math.Max for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MaxSliceFloat[T constraints.Float](haystack []T) (T, error) {
	var max T

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	max = haystack[0]
	for _, val := range haystack {
		max = T(math.Max(float64(max), float64(val)))
	}

	return max, nil
}

// MaxSlicePred finds the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxSlicePred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {
	var max T

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	max = haystack[0]
	for _, val := range haystack {
		if binary_predicate(val, max) {
			max = val
		}
	}

	return max, nil
}

// MaxMapInt finds the largest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MaxMapInt[TKey comparable, TValue constraints.Integer](haystack map[TKey]TValue) (TValue, error) {
	var max TValue

	for _, val := range haystack {
		max = val
		break
	}

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	for _, val := range haystack {
		if val > max {
			max = val
		}
	}

	return max, nil
}

// MaxMapFloat finds the largest value in haystack.
// This funnction uses math.Max for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MaxMapFloat[TKey comparable, TValue constraints.Float](haystack map[TKey]TValue) (TValue, error) {
	var max TValue

	for _, val := range haystack {
		max = val
		break
	}

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	for _, val := range haystack {
		max = TValue(math.Max(float64(max), float64(val)))
	}

	return max, nil
}

// MaxMapPred finds the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate func(TValue, TValue) bool) (TValue, error) {
	var max TValue

	for _, val := range haystack {
		max = val
		break
	}

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	for _, val := range haystack {
		if binary_predicate(val, max) {
			max = val
		}
	}

	return max, nil
}

// MinMaxSliceInt finds the smallest and largest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxSliceInt[T constraints.Integer](haystack []T) (T, T, error) {
	var min T
	var max T

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	min = haystack[0]
	for _, val := range haystack {
		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}

	return min, max, nil
}

// MinMaxSliceFloat finds the smallest and largest value in haystack.
// This funnction uses math.MinMax for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxSliceFloat[T constraints.Float](haystack []T) (T, T, error) {
	var min T
	var max T

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	min = haystack[0]
	for _, val := range haystack {
		min = T(math.Min(float64(min), float64(val)))
		max = T(math.Max(float64(max), float64(val)))
	}

	return min, max, nil
}

// MinMaxSlicePred finds the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxSlicePred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (T, T, error) {
	var min T
	var max T

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	min = haystack[0]
	for _, val := range haystack {
		if binary_predicate_min(val, min) {
			min = val
		}
		if binary_predicate_max(val, min) {
			max = val
		}
	}

	return min, max, nil
}

// MinMaxMapInt finds the smallest and largest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxMapInt[TKey comparable, TValue constraints.Integer](haystack map[TKey]TValue) (TValue, TValue, error) {
	var min TValue
	var max TValue

	for _, val := range haystack {
		min = val
		max = val
		break
	}

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	for _, val := range haystack {
		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}

	return min, max, nil
}

// MinMaxMapFloat finds the smallest and largest value in haystack.
// This funnction uses math.MinMax for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxMapFloat[TKey comparable, TValue constraints.Float](haystack map[TKey]TValue) (TValue, TValue, error) {
	var min TValue
	var max TValue

	for _, val := range haystack {
		min = val
		max = val
		break
	}

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	for _, val := range haystack {
		min = TValue(math.Min(float64(min), float64(val)))
		max = TValue(math.Max(float64(max), float64(val)))
	}

	return min, max, nil
}

// MinMaxMapPred finds the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate_min func(TValue, TValue) bool, binary_predicate_max func(TValue, TValue) bool) (TValue, TValue, error) {
	var min TValue
	var max TValue

	for _, val := range haystack {
		min = val
		max = val
		break
	}

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	for _, val := range haystack {
		if binary_predicate_min(val, min) {
			min = val
		}
		if binary_predicate_max(val, max) {
			max = val
		}
	}

	return min, max, nil
}

// MinElementSliceInt finds the index of the smallest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MinElementSliceInt[T constraints.Integer](haystack []T) (int, error) {
	var min int

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	min = 0
	for i, val := range haystack {
		if val < haystack[min] {
			min = i
		}
	}

	return min, nil
}

// MinElementSliceFloat finds the index of the smallest value in haystack.
// This funnction uses math.Min for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MinElementSliceFloat[T constraints.Float](haystack []T) (int, error) {
	var min int

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	min = 0
	for i, val := range haystack {
		if float64(val) == math.Min(float64(haystack[min]), float64(val)) {
			min = i
		}
	}

	return min, nil
}

// MinElementSlicePred finds the index of the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinElementSlicePred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {
	var min int

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	min = 0
	for key, val := range haystack {
		if binary_predicate(val, haystack[min]) {
			min = key
		}
	}

	return min, nil
}

// MinElementMapInt finds the index of the smallest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MinElementMapInt[TKey comparable, TValue constraints.Integer](haystack map[TKey]TValue) (TKey, error) {
	var min TKey

	for key := range haystack {
		min = key
		break
	}

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	for key, val := range haystack {
		if val < haystack[min] {
			min = key
		}
	}

	return min, nil
}

// MinElementMapFloat finds the index of the smallest value in haystack.
// This funnction uses math.Min for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MinElementMapFloat[TKey comparable, TValue constraints.Float](haystack map[TKey]TValue) (TKey, error) {
	var min TKey

	for key := range haystack {
		min = key
		break
	}

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	for key, val := range haystack {

		if float64(val) == math.Min(float64(haystack[min]), float64(val)) {
			min = key
		}
	}

	return min, nil
}

// MinElementMapPred finds the index of the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinElementMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate func(TValue, TValue) bool) (TKey, error) {
	var min TKey

	for key := range haystack {
		min = key
		break
	}

	if len(haystack) == 0 {
		return min, EmptyIterableError{}
	}

	for key, val := range haystack {
		if binary_predicate(val, haystack[min]) {
			min = key
		}
	}

	return min, nil
}

// MaxElementSliceInt finds the index of the largest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElementSliceInt[T constraints.Integer](haystack []T) (int, error) {
	var max int

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	max = 0
	for i, val := range haystack {
		if val > haystack[max] {
			max = i
		}
	}

	return max, nil
}

// MaxElementSliceFloat finds the index of the largest value in haystack.
// This funnction uses math.Max for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElementSliceFloat[T constraints.Float](haystack []T) (int, error) {
	var max int

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	max = 0
	for key, val := range haystack {
		if float64(val) == math.Max(float64(haystack[max]), float64(val)) {
			max = key
		}
	}

	return max, nil
}

// MaxElementSlicePred finds the index of the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElementSlicePred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {
	var max int

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	max = 0
	for key, val := range haystack {
		if binary_predicate(val, haystack[max]) {
			max = key
		}
	}

	return max, nil
}

// MaxElementMapInt finds the index of the largest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElementMapInt[TKey comparable, TValue constraints.Integer](haystack map[TKey]TValue) (TKey, error) {
	var max TKey

	for key := range haystack {
		max = key
		break
	}

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	for key, val := range haystack {
		if val > haystack[max] {
			max = key
		}
	}

	return max, nil
}

// MaxElementMapFloat finds the index of the largest value in haystack.
// This funnction uses math.Max for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElementMapFloat[TKey comparable, TValue constraints.Float](haystack map[TKey]TValue) (TKey, error) {
	var max TKey

	for key := range haystack {
		max = key
		break
	}

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	for key, val := range haystack {
		if float64(val) == math.Max(float64(haystack[max]), float64(val)) {
			max = key
		}
	}

	return max, nil
}

// MaxElementMapPred finds the index of the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElementMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate func(TValue, TValue) bool) (TKey, error) {
	var max TKey

	for key := range haystack {
		max = key
		break
	}

	if len(haystack) == 0 {
		return max, EmptyIterableError{}
	}

	for key, val := range haystack {
		if binary_predicate(val, haystack[max]) {
			max = key
		}
	}

	return max, nil
}

// MinMaxElementSliceInt finds the index of the smallest and largest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElementSliceInt[T constraints.Integer](haystack []T) (int, int, error) {
	var min int
	var max int

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	min = 0
	for i, val := range haystack {
		if val < haystack[min] {
			min = i
		}

		if val > haystack[max] {
			max = i
		}
	}

	return min, max, nil
}

// MinMaxElementSliceFloat finds the index of the smallest and largest value in haystack.
// This funnction uses math.MinMax for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElementSliceFloat[T constraints.Float](haystack []T) (int, int, error) {
	var min int
	var max int

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	min = 0
	for i, val := range haystack {
		if float64(val) == math.Min(float64(haystack[min]), float64(val)) {
			min = i
		}

		if float64(val) == math.Max(float64(haystack[max]), float64(val)) {
			max = i
		}
	}

	return min, max, nil
}

// MinMaxElementSlicePred finds the index of the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElementSlicePred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (int, int, error) {
	var min int
	var max int

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	min = 0
	for i, val := range haystack {
		if binary_predicate_min(val, haystack[min]) {
			min = i
		}
		if binary_predicate_max(val, haystack[min]) {
			max = i
		}
	}

	return min, max, nil
}

// MinMaxElementMapInt finds the index of the smallest and largest value in haystack.
// This funnction is optimized for integers.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElementMapInt[TKey comparable, TValue constraints.Integer](haystack map[TKey]TValue) (TKey, TKey, error) {
	var min TKey
	var max TKey

	for key := range haystack {
		min = key
		max = key
		break
	}

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	for key, val := range haystack {
		if val < haystack[min] {
			min = key
		}

		if val > haystack[max] {
			max = key
		}
	}

	return min, max, nil
}

// MinMaxElementMapFloat finds the index of the smallest and largest value in haystack.
// This funnction uses math.MinMax for robustness.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElementMapFloat[TKey comparable, TValue constraints.Float](haystack map[TKey]TValue) (TKey, TKey, error) {
	var min TKey
	var max TKey

	for key := range haystack {
		min = key
		max = key
		break
	}

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	for key, val := range haystack {

		if float64(val) == math.Min(float64(haystack[min]), float64(val)) {
			min = key
		}

		if float64(val) == math.Max(float64(haystack[max]), float64(val)) {
			max = key
		}
	}

	return min, max, nil
}

// MinMaxElementMapPred finds the index of the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElementMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate_min func(TValue, TValue) bool, binary_predicate_max func(TValue, TValue) bool) (TKey, TKey, error) {
	var min TKey
	var max TKey

	for key := range haystack {
		min = key
		max = key
		break
	}

	if len(haystack) == 0 {
		return min, max, EmptyIterableError{}
	}

	for i, val := range haystack {
		if binary_predicate_min(val, haystack[min]) {
			min = i
		}
		if binary_predicate_max(val, haystack[max]) {
			max = i
		}
	}

	return min, max, nil
}
