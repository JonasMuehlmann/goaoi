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
func AllOfSlice[T comparable](container []T, unary_predicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		if !unary_predicate(value) {
			return ComparisonError[int]{i}
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
			return ComparisonError[TKey]{key}
		}
	}

	return nil
}

// AnyOfSlice checks that unary_predicate(val) == true for ANY val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func AnyOfSlice[T comparable](container []T, unary_predicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unary_predicate(value) {
			return nil
		}
	}

	return ComparisonError[int]{len(container) - 1}
}

// AnyOfMap checks that unary_predicate(val) == true for ANY val in container.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func AnyOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unary_predicate(value) {
			return nil
		}
	}

	return ComparisonError[TKey]{}
}

// NoneOfSlice checks that unary_predicate(val) == true for NO val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func NoneOfSlice[T comparable](container []T, unary_predicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unary_predicate(value) {
			return ComparisonError[int]{len(container) - 1}
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

	for _, value := range container {
		if unary_predicate(value) {
			return ComparisonError[TKey]{}
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
func ForeachSlice[T comparable](container []T, unary_func func(T) error) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		err := unary_func(value)
		if err != nil {
			return ExecutionError[int]{i, err}
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
			return ExecutionError[TKey]{key, err}
		}
	}

	return nil
}

// ForeachSliceUnsafe executes unary_func(val) for each val in container.
//
// Possible Error values:
//    - EmptyIterableError
func ForeachSliceUnsafe[T comparable](container []T, unary_func func(T)) error {
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
func FillSlice[T comparable](arr *[]T, filler T) []T {
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
			return ExecutionError[TValue]{container[key], err}
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
func TransformSlice[T comparable](container []T, transformer func(*T) error) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		err := transformer(&container[i])
		if err != nil {
			return ExecutionError[T]{value, err}
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
func TransformSliceUnsafe[T comparable](container []T, transformer func(*T)) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i := range container {
		transformer(&container[i])
	}

	return nil
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
func MinMapInt[THaystack comparable, TValue constraints.Integer](haystack map[THaystack]TValue) (TValue, error) {
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
func MinMapFloat[THaystack comparable, TValue constraints.Float](haystack map[THaystack]TValue) (TValue, error) {
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
func MinMapPred[THaystack comparable, TValue constraints.Ordered](haystack map[THaystack]TValue, binary_predicate func(TValue, TValue) bool) (TValue, error) {
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
