// Package goaoi implements conventient algorithms for processing iterables.
// It is inspired by the algorithm header from the C++ standard template library (STL for short).
package goaoi

import (
	"bytes"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
	iteratoradapters "github.com/JonasMuehlmann/goaoi/iterator_adapters"
	"golang.org/x/exp/constraints"
)

// FindIfMap finds the first key where unaryPredicate(haystack[key]) == true.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindIfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, unaryPredicate func(TValue) bool) (TKey, error) {
	var zeroVal TKey

	if len(haystack) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	for key, value := range haystack {
		if unaryPredicate(value) {
			return key, nil
		}
	}

	return zeroVal, ElementNotFoundError{}
}

// FindIfSlice finds the first index i where unaryPredicate(haystack[i]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindIfSlice[T comparable](haystack []T, unaryPredicate func(T) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, value := range haystack {
		if unaryPredicate(value) {
			return i, nil
		}
	}

	return 0, ElementNotFoundError{}
}

// FindIfString finds the first index i where unaryPredicate(haystack[i]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindIfString(haystack string, unaryPredicate func(rune) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, value := range haystack {
		if unaryPredicate(value) {
			return i, nil
		}
	}

	return 0, ElementNotFoundError{}
}

// FindIfIterator finds the first index i where unaryPredicate(haystack[i]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindIfIterator[TKey any, TValue comparable](haystack compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) (int, error) {
	if haystack.IsEnd() {
		return 0, EmptyIterableError{}
	}

	for haystack.Next() {
		value, _ := haystack.Get()
		if unaryPredicate(value) {
			i, _ := haystack.Index()

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

// FindEndStringPred finds the beginning of the last occurrence of sub in super.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindEndStringPred(super string, sub string, binary_predicate func(byte, byte) bool) (int, error) {
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

// TODO: implement FindEndIteratorPred

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

// FindFirstOfStringPred finds the first index where an element of haystack is equal to any element in needles.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func FindFirstOfStringPred(haystack string, needles string, binary_predicate func(rune, rune) bool) (int, error) {
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

// AllOfSlice checks that unaryPredicate(val) == true for ALL val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func AllOfSlice[T any](container []T, unaryPredicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		if !unaryPredicate(value) {
			return ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
		}
	}

	return nil
}

// AllOfMap checks that unaryPredicate(val) == true for ALL val in container.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func AllOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unaryPredicate func(TValue) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range container {
		if !unaryPredicate(value) {
			return ComparisonError[TKey, TValue]{BadItemIndex: key, BadItem: value}
		}
	}

	return nil
}

// AllOfString checks that unaryPredicate(val) == true for ALL val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func AllOfString(container string, unaryPredicate func(rune) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		if !unaryPredicate(value) {
			return ComparisonError[int, rune]{BadItemIndex: i, BadItem: value}
		}
	}

	return nil
}

// AllOfIterator checks that unaryPredicate(val) == true for ALL val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func AllOfIterator[TKey any, TValue any](container compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) error {
	if container.IsEnd() {
		return EmptyIterableError{}
	}

	for container.Next() {
		value, _ := container.Get()
		if !unaryPredicate(value) {
			i, _ := container.Index()
			return ComparisonError[int, TValue]{BadItemIndex: i, BadItem: value}
		}
	}

	return nil
}

// AnyOfSlice checks that unaryPredicate(val) == true for ANY val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func AnyOfSlice[T any](container []T, unaryPredicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unaryPredicate(value) {
			return nil
		}
	}

	return ElementNotFoundError{}
}

// AnyOfMap checks that unaryPredicate(val) == true for ANY val in container.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func AnyOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unaryPredicate func(TValue) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unaryPredicate(value) {
			return nil
		}
	}

	return ElementNotFoundError{}
}

// AnyOfString checks that unaryPredicate(val) == true for ANY val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func AnyOfString(container string, unaryPredicate func(rune) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unaryPredicate(value) {
			return nil
		}
	}

	return ElementNotFoundError{}
}

// AnyOfIterator checks that unaryPredicate(val) == true for ANY val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func AnyOfIterator[TKey any, TValue any](container compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) error {
	if container.IsEnd() {
		return EmptyIterableError{}
	}

	for container.Next() {
		value, _ := container.Get()
		if unaryPredicate(value) {
			return nil
		}
	}

	return ElementNotFoundError{}
}

// NoneOfSlice checks that unaryPredicate(val) == true for NO val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func NoneOfSlice[T any](container []T, unaryPredicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		if unaryPredicate(value) {
			return ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
		}
	}

	return nil
}

// NoneOfMap checks that unaryPredicate(val) == true for NO val in container.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func NoneOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unaryPredicate func(TValue) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range container {
		if unaryPredicate(value) {
			return ComparisonError[TKey, TValue]{BadItemIndex: key, BadItem: value}
		}
	}

	return nil
}

// NoneOfString checks that unaryPredicate(val) == true for NO val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func NoneOfString(container string, unaryPredicate func(rune) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		if unaryPredicate(value) {
			return ComparisonError[int, rune]{BadItemIndex: i, BadItem: value}
		}
	}

	return nil
}

// NoneOfIterator checks that unaryPredicate(val) == true for ALL val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
func NoneOfIterator[TKey any, TValue any](container compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) error {
	if container.IsEnd() {
		return EmptyIterableError{}
	}

	for container.Next() {
		value, _ := container.Get()
		if unaryPredicate(value) {
			i, _ := container.Index()
			return ComparisonError[int, TValue]{BadItemIndex: i, BadItem: value}
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

// ForeachString executes unary_func(val) for each val in container.
// Errors returned by unary_func are propagated to the caller of ForeachSlice.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func ForeachString(container string, unary_func func(rune) error) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		err := unary_func(value)
		if err != nil {
			return ExecutionError[int, byte]{BadItemIndex: i, BadItem: container[i], Inner: err}
		}
	}

	return nil
}

// ForeachIterator executes unary_func(val) for each val in container.
// Errors returned by unary_func are propagated to the caller of ForeachSlice.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func ForeachIterator[TKey any, TValue any](container compounditerators.ReadForIndexIterator[TKey, TValue], unary_func func(TValue) error) error {
	if container.IsEnd() {
		return EmptyIterableError{}
	}

	for container.Next() {
		value, _ := container.Get()
		err := unary_func(value)
		if err != nil {
			i, _ := container.Index()
			return ExecutionError[int, TValue]{BadItemIndex: i, BadItem: value, Inner: err}
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

// ForeachStringUnsafe executes unary_func(val) for each val in container.
// Errors returned by unary_func are propagated to the caller of ForeachSlice.
//
// Possible Error values:
//    - EmptyIterableError
func ForeachStringUnsafe(container string, unary_func func(rune)) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		unary_func(value)
	}

	return nil
}

// ForeachIteratorUnsafe executes unary_func(val) for each val in container.
// Errors returned by unary_func are propagated to the caller of ForeachSlice.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func ForeachIteratorUnsafe[TKey any, TValue any](container compounditerators.ReadForIndexIterator[TKey, TValue], unary_func func(TValue)) error {
	if container.IsEnd() {
		return EmptyIterableError{}
	}

	for container.Next() {
		value, _ := container.Get()
		unary_func(value)
	}

	return nil
}

// CountIfSlice counts for how many val of container unaryPredicate(val) == true.
//
// Possible Error values:
//    - EmptyIterableError
func CountIfSlice[T comparable](container []T, unaryPredicate func(T) bool) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if unaryPredicate(value) {
			counter++
		}
	}

	return counter, nil
}

// CountIfMap counts for how many val of container unaryPredicate(val) == true.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CountIfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unaryPredicate func(TValue) bool) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if unaryPredicate(value) {
			counter++
		}
	}

	return counter, nil
}

// CountIfString counts for how many val of container unaryPredicate(val) == true.
//
// Possible Error values:
//    - EmptyIterableError
func CountIfString(container string, unaryPredicate func(rune) bool) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if unaryPredicate(value) {
			counter++
		}
	}

	return counter, nil
}

// CountIfIterator counts for how many val of container unaryPredicate(val) == true.
//
// Possible Error values:
//    - EmptyIterableError
func CountIfIterator[TKey any, TValue comparable](container compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) (int, error) {
	if container.IsEnd() {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for container.Next() {
		value, _ := container.Get()
		if unaryPredicate(value) {
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

// MismatchStringPred finds the first index i where binary_predicate(iterable1[i], iterable2[i] == false).
//
// Possible Error values:
//    - EmptyIterableError
//    - EqualIteratorsError
func MismatchStringPred(iterable1 string, iterable2 string, binary_predicate func(byte, byte) bool) (int, error) {
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

// TODO: implement MismatchIteratorPred

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

// AdjacentFindStringPred finds the first index i where binary_predicate(container[i], container[i+1]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func AdjacentFindStringPred(container string, binary_predicate func(byte, byte) bool) (int, error) {
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

// AdjacentFindIteratorPred finds the first index i where binary_predicate(container[i], container[i+1]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
func AdjacentFindIteratorPred[TKey any, TValue comparable](container compounditerators.ReadForIndexIterator[TKey, TValue], binary_predicate func(TValue, TValue) bool) (int, error) {
	if container.IsEnd() {
		return 0, EmptyIterableError{}
	}

	var cur TValue
	prev, _ := container.Get()
	container.Next()

	for container.Next() {
		cur, _ = container.Get()
		if binary_predicate(cur, prev) {
			i, _ := container.Index()

			return i, nil
		}

		prev = cur
	}

	return 0, ElementNotFoundError{}
}

// TakeWhileSlice returns a copy of original until the first element not satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func TakeWhileSlice[T comparable](original []T, unaryPredicate func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if !unaryPredicate(value) {
			return newContainer, nil
		}

		newContainer = append(newContainer, value)
	}

	return newContainer, nil
}

// TakeWhileString returns a copy of original until the first element not satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func TakeWhileString(original string, unaryPredicate func(rune) bool) (string, error) {
	var out bytes.Buffer

	if len(original) == 0 {
		return "", EmptyIterableError{}
	}

	var i int
	var value rune
	for i, value = range original {
		if !unaryPredicate(value) {
			return original[:i], nil
		}

		out.WriteRune(value)
	}

	return out.String(), nil
}

// TakeWhileIterator returns a copy of original until the first element not satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func TakeWhileIterator[TKey any, TValue any](original compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) (compounditerators.ReadForIndexIterator[TKey, TValue], error) {
	if original.IsEnd() {
		return original, EmptyIterableError{}
	}

	return iteratoradapters.NewTakeWhile[TKey, TValue](original, unaryPredicate), nil
}

// DropWhileSlice returns a copy of original starting from first element not satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func DropWhileSlice[T comparable](original []T, unaryPredicate func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	i := 0
	for _, value := range original {
		if !unaryPredicate(value) {
			break
		}

		i++
	}

	newContainer = append(newContainer, original[i:]...)

	return newContainer, nil
}

// DropWhileString returns a copy of original starting from first element not satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func DropWhileString(original string, unaryPredicate func(rune) bool) (string, error) {
	if len(original) == 0 {
		return "", EmptyIterableError{}
	}

	var i int
	var value rune

	for i, value = range original {
		if unaryPredicate(value) {
			return original[i-1:], nil
		}

	}

	return "", nil
}

// DropWhileIterator returns a copy of original until the first element satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func DropWhileIterator[TKey any, TValue any](original compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) (compounditerators.ReadForIndexIterator[TKey, TValue], error) {
	if original.IsEnd() {
		return original, EmptyIterableError{}
	}

	return iteratoradapters.NewDropWhile[TKey, TValue](original, unaryPredicate), nil
}

// CopyIfSlice returns a copy of original with all element satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyIfSlice[T comparable](original []T, unaryPredicate func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if unaryPredicate(value) {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

// CopyIfMap returns a copy of original with all key-value pairs satisfying unaryPredicate(value) == true).
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CopyIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, unaryPredicate func(TValue) bool) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if unaryPredicate(value) {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

// CopyIfString returns a copy of original with all element satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyIfString(original string, unaryPredicate func(rune) bool) (string, error) {
	var out bytes.Buffer

	if len(original) == 0 {
		return "", EmptyIterableError{}
	}

	for _, value := range original {
		if unaryPredicate(value) {
			out.WriteRune(value)
		}
	}

	return out.String(), nil
}

// CopyIfString returns a copy of original with all element satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyIfIterator[TKey any, TValue any](original ds.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) (ds.ReadForIndexIterator[TKey, TValue], error) {

	if original.IsEnd() {
		return original, EmptyIterableError{}
	}

	return iteratoradapters.NewCopyIf[TKey, TValue](original, unaryPredicate), nil
}

// CopyReplaceIfSlice returns a copy of original where each element satisfying unaryPredicate(element) == true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfSlice[T comparable](original []T, unaryPredicate func(T) bool, replacement T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if unaryPredicate(value) {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

// CopyReplaceIfMap returns a copy of original where each value of a key-value pair satisfying unaryPredicate(value) == true is replaced with replacement.
// Note that the iteration order of a map is not stable.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, unaryPredicate func(TValue) bool, replacement TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if unaryPredicate(value) {
			newContainer[key] = replacement
		} else {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

// CopyReplaceIfString returns a copy of original with all element satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfString(original string, unaryPredicate func(rune) bool, replacement rune) (string, error) {
	var out bytes.Buffer

	if len(original) == 0 {
		return "", EmptyIterableError{}
	}

	for _, value := range original {
		if unaryPredicate(value) {
			out.WriteRune(replacement)
		} else {
			out.WriteRune(value)
		}
	}

	return out.String(), nil
}

// CopyReplaceIfIterator returns a copy of original with all element satisfying unaryPredicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfIterator[TKey any, TValue any](original ds.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool, replacement TValue) (ds.ReadForIndexIterator[TKey, TValue], error) {

	if original.IsEnd() {
		return original, EmptyIterableError{}
	}

	return iteratoradapters.NewCopyReplaceIf[TKey, TValue](original, unaryPredicate, replacement), nil
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

// TransformIterator applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
// Errors returned by transformer are propagated to the caller of TransformSlice.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformIterator[TKey any, TValue any](container ds.ReadForIndexIterator[TKey, TValue], transformer func(TValue) (TValue, error)) (ds.ReadForIndexIterator[TKey, TValue], error) {
	if container.IsEnd() {
		return container, EmptyIterableError{}
	}

	return iteratoradapters.NewTransformIterator[TKey, TValue](container, transformer), nil
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

// TODO: implement TransformIteratorUnsafe

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

// TransformCopySlice applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
// Errors returned by transformer are propagated to the caller of TransformCopySlice.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformCopyString(container string, transformer func(rune) (rune, error)) (string, error) {
	var res bytes.Buffer

	if len(container) == 0 {
		return "", EmptyIterableError{}
	}

	for i, value := range container {
		newVal, err := transformer(value)
		if err != nil {
			return "", ExecutionError[int, rune]{BadItemIndex: i, BadItem: value, Inner: err}
		}

		res.WriteRune(newVal)
	}

	return res.String(), nil
}

// TODO: implement TransformCopyIterator

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

// TransformCopySlice applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
// Errors returned by transformer are propagated to the caller of TransformCopySlice.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformCopyStringUnsafe(container string, transformer func(rune) rune) (string, error) {
	var res bytes.Buffer

	if len(container) == 0 {
		return "", EmptyIterableError{}
	}

	for _, value := range container {
		newVal := transformer(value)

		res.WriteRune(newVal)
	}

	return res.String(), nil
}

// TODO: implement TransformCopyIteratorUnsafe

// MinSlicePred finds the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinSlicePred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, T, error) {
	var min T
	var iMin int

	if len(haystack) == 0 {
		return iMin, min, EmptyIterableError{}
	}

	min = haystack[0]
	for i, val := range haystack {
		if binary_predicate(val, min) {
			min = val
			iMin = i
		}
	}

	return iMin, min, nil
}

// MinMapPred finds the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate func(TValue, TValue) bool) (TKey, TValue, error) {
	var min TValue
	var val TValue
	var keyMin TKey

	for keyMin, val = range haystack {
		min = val
		break
	}

	if len(haystack) == 0 {
		return keyMin, min, EmptyIterableError{}
	}

	for key, val := range haystack {
		if binary_predicate(val, min) {
			min = val
			keyMin = key
		}
	}

	return keyMin, min, nil
}

// MinStringPred finds the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinStringPred[T constraints.Ordered](haystack string, binary_predicate func(rune, rune) bool) (int, rune, error) {
	var min rune
	var iMin int

	if len(haystack) == 0 {
		return iMin, min, EmptyIterableError{}
	}

	for i, val := range haystack {
		if binary_predicate(val, min) {
			min = val
			iMin = i
		}
	}

	return iMin, min, nil
}

// TODO: implement MinIteratorPred

// MaxSlicePred finds the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxSlicePred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, T, error) {
	var max T
	var IMax int

	if len(haystack) == 0 {
		return IMax, max, EmptyIterableError{}
	}

	max = haystack[0]
	for i, val := range haystack {
		if binary_predicate(val, max) {
			max = val
			IMax = i
		}
	}

	return IMax, max, nil
}

// MaxMapPred finds the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate func(TValue, TValue) bool) (TKey, TValue, error) {
	var max TValue
	var keyMax TKey

	for key, val := range haystack {
		max = val
		keyMax = key
		break
	}

	if len(haystack) == 0 {
		return keyMax, max, EmptyIterableError{}
	}

	for key, val := range haystack {
		if binary_predicate(val, max) {
			max = val
			keyMax = key
		}
	}

	return keyMax, max, nil
}

// MaxStringPred finds the highest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxStringPred[T constraints.Ordered](haystack string, binary_predicate func(rune, rune) bool) (int, rune, error) {
	var max rune
	var iMax int

	if len(haystack) == 0 {
		return iMax, max, EmptyIterableError{}
	}

	for i, val := range haystack {
		if binary_predicate(val, max) {
			max = val
			iMax = i
		}
	}

	return iMax, max, nil
}

// TODO: implement MinIteratorPred

// MinMaxSlicePred finds the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxSlicePred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (int, int, T, T, error) {
	var min T
	var max T
	var iMin int
	var iMax int

	if len(haystack) == 0 {
		return iMin, iMax, min, max, EmptyIterableError{}
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

	return iMin, iMax, min, max, nil
}

// MinMaxMapPred finds the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate_min func(TValue, TValue) bool, binary_predicate_max func(TValue, TValue) bool) (TKey, TKey, TValue, TValue, error) {
	var min TValue
	var max TValue
	var keyMin TKey
	var keyMax TKey

	for key, val := range haystack {
		min = val
		max = val

		keyMin = key
		keyMax = key
		break
	}

	if len(haystack) == 0 {
		return keyMin, keyMax, min, max, EmptyIterableError{}
	}

	for key, val := range haystack {
		if binary_predicate_min(val, min) {
			min = val
			keyMin = key
		}
		if binary_predicate_max(val, max) {
			max = val
			keyMax = key
		}
	}

	return keyMin, keyMax, min, max, nil
}

// MinMaxStringPred finds the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxStringPred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (int, int, T, T, error) {
	var min T
	var max T
	var iMin int
	var iMax int

	if len(haystack) == 0 {
		return iMin, iMax, min, max, EmptyIterableError{}
	}

	min = haystack[0]
	for i, val := range haystack {
		if binary_predicate_min(val, min) {
			min = val
			iMin = i
		}
		if binary_predicate_max(val, min) {
			max = val
			iMax = i
		}
	}

	return iMin, iMax, min, max, nil
}

// TODO: implement MinMaxIteratorPred
