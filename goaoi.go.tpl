// Package goaoi implements conventient algorithms for processing iterables.
// It is inspired by the algorithm header from the C++ standard template library (STL for short).
package goaoi

import (
	"golang.org/x/exp/constraints"
    "github.com/JonasMuehlmann/datastructures.go/ds"
)


{{/* For reference
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
*/}}



{{ range $variant := . }}


// FindIf{{ $variant }} finds the first index i where unary_predicate(haystack[i]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
{{ if eq $variant "Slice" }}
func FindIf{{ $variant }}[T comparable](haystack []T, unary_predicate func( needle T) bool) (index int, err error) {
{{ else if eq $variant "Map" }}
func FindIf{{ $variant }}[TKey comparable, TValue comparable](haystack map[TKey]TValue, unary_predicate func( needle TValue) bool) (index TKey, err error) {
{{ else if eq $variant "String" }}
func FindIf{{ $variant }}(haystack string, unary_predicate func( needle rune) bool) (index int, err error) {
{{ else if eq $variant "Iterator" }}
func FindIf{{ $variant }}[T comparable](haystack ds.ReadCompForIndexIterator[T], unary_predicate func( needle T) bool) (index int, err error) {
{{ end }}
    {{ if eq $variant "Iterator" }}
	if haystack.IsEnd() {
    {{ else }}
	if len(haystack) == 0 {
    {{ end }}
        err = EmptyIterableError{}
		return
	}

    {{ if eq $variant "Iterator" }}
    for haystack.Next() {
        value, _ := haystack.Get()
    {{ else }}
	for i, value := range haystack {
    {{ end }}
		if unary_predicate(value) {
        {{ if eq $variant "Iterator" }}
            index, _ = haystack.Index()
        {{ else }}
            index = i
        {{ end }}
			return
		}
	}

	err = ElementNotFoundError{}

    return
}

{{ end }}
{{ range $variant := . }}

{{ if ne $variant "Map" }}
// FindEnd{{ $variant }}Pred finds the beginning of the last occurrence of sub in super.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
{{ if eq $variant "Slice" }}
func FindEnd{{ $variant }}Pred[T comparable](super []T, sub []T, binary_predicate func(T, T) bool) (index int, err error) {
{{ else if eq $variant "String" }}
func FindEnd{{ $variant }}Pred(super string, sub string, binary_predicate func(string, string) bool) (index int, err error) {
{{ else if eq $variant "Iterator" }}
func FindEnd{{ $variant }}Pred[T comparable](super ds.ReadCompForIndexIterator[T], sub ds.ReadCompForIndexIterator[T], binary_predicate func(T, T) bool) (index int, err error) {
{{ end }}
    {{ if eq $variant "Iterator" }}
    if haystack.IsEnd() {
    {{ else }}
	if len(haystack) == 0 || len(needles) == 0 {
    {{ end }}
		err =  EmptyIterableError{}

		return
	}

    {{ if eq $variant "Iterator" }}
    lenSub := super.Size()
    lenSiper := sub.Size()
    {{ else }}
    lenSub := len(sub)
    lenSiper := len(super)
    {{ end }}

    {{ if eq $variant "String" }}
    var curSub rune
    var curSUper rune
    {{ else }}
    var curSub T
    var curSuper T
    {{ end }}


OUTER:
	for i := lenSuper - 1; i >= lenSub-1; i-- {
		for j := 0; j < lenSub; j++ {
            curSub = sub[lenSub-1-j]
            curSuper = super[i-j]

			if !binary_predicate(curSuper, curSub) {
				continue OUTER
			}
		}
		index =  i - lenSub + 1


		return
	}


	err =  ElementNotFoundError{}

	return
}

{{ end }}
{{ end }}
{{ range $variant := . }}

// FindFirstOf{{ $variant }}Pred finds the first index where an element of haystack is equal to any element in needles.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
{{ if eq $variant "Slice" }}
func FindFirstOf{{ $variant }}Pred[T comparable](haystack []T, needles []T, binary_predicate func(T, T) bool) (index int, err error) {
{{ else if eq $variant "Map" }}
func FindFirstOf{{ $variant }}Pred[TKey comparable, TValue comparable](haystack map[TKey]TValue, needles []TValue, binary_predicate func(TValue, TValue) bool) (index TKey, err error) {
{{ else if eq $variant "String" }}
func FindFirstOf{{ $variant }}Pred(haystack string, needles string, binary_predicate func(rune, rune) bool) (index int, err error) {
{{ else if eq $variant "Iterator" }}
    // TODO: This should take an iterator, which gets copeid before each run through
func FindFirstOf{{ $variant }}Pred[T comparable](haystack ds.ReadCompForIndexIterator[T], needles []T, binary_predicate func(T, T) bool) (index int,err  error) {
{{ end }}
    {{ if eq $variant "Iterator" }}
    if haystack.IsEnd() {
    {{ else }}
	if len(haystack) == 0 || len(needles) == 0 {
    {{ end }}

		err =  EmptyIterableError{}

		return
	}

    {{ if eq $variant "String" }}
    var curHaystack rune
    {{ else }}
    var curHaystack T
    {{ end }}

    {{ if eq $variant "Iterator" }}
    for haystack.Next() {
		for _, needleValue := range needles {
            curHaystack, _ = haystack.Get()
    {{ else }}
	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
    {{ end }}
			if binary_predicate(haystackValue, needleValue) {
				index =  i


				return
			}
		}
	}


	err =  ElementNotFoundError{}

	return
}

{{ end }}
{{ range $variant := . }}

// AllOf{{ $variant }} checks that unary_predicate(val) == true for ALL val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
{{ if eq $variant "Slice" }}
func AllOf{{ $variant }}[T any](container []T, unary_predicate func(T) bool) error {
{{ else if eq $variant "Map" }}
func AllOf{{ $variant }}[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {
{{ else if eq $variant "String" }}
func AllOf{{ $variant }}(container string, unary_predicate func(rune) bool) error {
{{ else if eq $variant "Iterator" }}
func AllOf{{ $variant }}[T any](container ds.ReadCompForIndexIterator[T], unary_predicate func(T) bool) error {
{{ end }}
    {{ if eq $variant "Iterator" }}
    if container.IsEnd() {
    {{ else }}
	if len(container) == 0 {
    {{ end }}
		return EmptyIterableError{}
	}

    {{ if eq $variant "Iterator" }}
    var value T

    for container.Next() {
        value, _ = container.Get()
    {{ else }}
	for i, value := range container {
    {{ end }}

		if !unary_predicate(value) {
        {{ if eq $variant "Iterator" }}
        i, _ := container.Index()
        {{ end }}

        {{ if eq $variant "Map" }}
        err =  ComparisonError[TKey, TValue]{BadItemIndex: i, BadItem: value}
        {{ else if eq $variant "string" }}
        err =  ComparisonError[int, rune]{BadItemIndex: i, BadItem: value}
        {{ else }}
        err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
        {{ end }}

			return
		}
	}

	return nil
}

{{ end }}
{{ range $variant := . }}

// AnyOf{{ $variant }} checks that unary_predicate(val) == true for ANY val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
{{ if eq $variant "Slice" }}
func AnyOf{{ $variant }}[T any](container []T, unary_predicate func(T) bool) error {
{{ else if eq $variant "Map" }}
func AnyOf{{ $variant }}[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {
{{ else if eq $variant "String" }}
func AnyOf{{ $variant }}(container string, unary_predicate func(rune) bool) error {
{{ else if eq $variant "Iterator" }}
func AnyOf{{ $variant }}[T any](container ds.ReadCompForIndexIterator[T], unary_predicate func(T) bool) error {
{{ end }}
    {{ if eq $variant "Iterator" }}
    if container.IsEnd() {
    {{ else }}
	if len(container) == 0 {
    {{ end }}
		return EmptyIterableError{}
	}

    {{ if eq $variant "Iterator" }}
    var value T

    for container.Next() {
        value, _ = container.Get()
    {{ else }}
	for i, value := range container {
    {{ end }}

		if unary_predicate(value) {
            return nil
		}
	}

    {{ if eq $variant "Iterator" }}
    i, _ := container.Index()
    {{ end }}

    {{ if eq $variant "Map" }}
    err =  ComparisonError[TKey, TValue]{BadItemIndex: i, BadItem: value}
    {{ else if eq $variant "string" }}
    err =  ComparisonError[int, rune]{BadItemIndex: i, BadItem: value}
    {{ else }}
    err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
    {{ end }}

    return
}

{{ end }}
{{ range $variant := . }}

// NoneOf{{ $variant }} checks that unary_predicate(val) == true for NO val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError
{{ if eq $variant "Slice" }}
func NoneOf{{ $variant }}[T any](container []T, unary_predicate func(T) bool) error {
{{ else if eq $variant "Map" }}
func NoneOf{{ $variant }}[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {
{{ else if eq $variant "String" }}
func NoneOf{{ $variant }}(container string, unary_predicate func(rune) bool) error {
{{ else if eq $variant "Iterator" }}
func NoneOf{{ $variant }}[T any](container ds.ReadCompForIndexIterator[T], unary_predicate func(T) bool) error {
{{ end }}
    {{ if eq $variant "Iterator" }}
    if container.IsEnd() {
    {{ else }}
	if len(container) == 0 {
    {{ end }}
		return EmptyIterableError{}
	}

    {{ if eq $variant "Iterator" }}
    var value T

    for container.Next() {
        value, _ = container.Get()
    {{ else }}
	for i, value := range container {
    {{ end }}

		if unary_predicate(value) {
        {{ if eq $variant "Iterator" }}
        i, _ := container.Index()
        {{ end }}

        {{ if eq $variant "Map" }}
        err =  ComparisonError[TKey, TValue]{BadItemIndex: i, BadItem: value}
        {{ else if eq $variant "string" }}
        err =  ComparisonError[int, rune]{BadItemIndex: i, BadItem: value}
        {{ else }}
        err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
        {{ end }}

			return
		}
	}

	return nil
}

{{ end }}

{{ range $variant := . }}

// Foreach{{ $variant }} executes unary_func(val) for each val in container.
// Errors returned by unary_func are propagated to the caller of Foreach{{ $variant }}.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
{{ if eq $variant "Slice" }}
func Foreach{{ $variant }}[T any](container []T, unary_func func(T) error) error {
{{ else if eq $variant "Map" }}
func Foreach{{ $variant }}[TKey comparable, TValue comparable](container map[TKey]TValue, unary_func func(TValue) error) error {
{{ else if eq $variant "String" }}
func Foreach{{ $variant }}(container string, unary_func func(rune) error) error {
{{ else if eq $variant "Iterator" }}
func Foreach{{ $variant }}[T any](container ds.ReadCompForIndexIterator[T], unary_func func(T) error) error {
{{ end }}
    {{ if eq $variant "Iterator" }}
    if container.IsEnd() {
    {{ else }}
	if len(container) == 0 {
    {{ end }}
		return EmptyIterableError{}
	}


 {{ if eq $variant "Iterator" }}
    var value T

    for container.Next() {
        value, _ = container.Get()
    {{ else }}
	for i, value := range container {
    {{ end }}

        err := unary_func(value)
		if err != nil{
        {{ if eq $variant "Iterator" }}
        i, _ := container.Index()
        {{ end }}

        {{ if eq $variant "Map" }}
        err = ExecutionError[TKey, TValue]{BadItemIndex: i, BadItem: value, Inner: err}
        {{ else if eq $variant "string" }}
        err = ExecutionError[int, rune]{BadItemIndex: i, BadItem: value, Inner: err}
        {{ else }}
        err = ExecutionError[int, T]{BadItemIndex: i, BadItem: value, Inner: err}
        {{ end }}

			return
		}
	}


	return nil
}

{{ end }}
{{ range $variant := . }}

// Foreach{{ $variant }}Unsafe executes unary_func(val) for each val in container.
//
// Possible Error values:
//    - EmptyIterableError
{{ if eq $variant "Slice" }}
func Foreach{{ $variant }}Unsafe[T any](container []T, unary_func func(T)) error {
{{ else if eq $variant "Map" }}
func Foreach{{ $variant }}Unsafe[TKey comparable, TValue comparable](container map[TKey]TValue, unary_func func(TValue)) error {
{{ else if eq $variant "String" }}
func Foreach{{ $variant }}Unsafe(container string, unary_func func(rune)) error {
{{ else if eq $variant "Iterator" }}
func Foreach{{ $variant }}Unsafe[T any](container ds.ReadCompForIndexIterator[T], unary_func func(T)) error {
{{ end }}
    {{ if eq $variant "Iterator" }}
    if container.IsEnd() {
    {{ else }}
	if len(container) == 0 {
    {{ end }}
		return EmptyIterableError{}
	}


 {{ if eq $variant "Iterator" }}
    var value T

    for container.Next() {
        value, _ = container.Get()
    {{ else }}
	for i, value := range container {
    {{ end }}
        unary_func(value)
    }


	return nil
}

{{ end }}
{{ range $variant := . }}

// CountIf{{ $variant }} counts for how many val of container unary_predicate(val) == true.
//
// Possible Error values:
//    - EmptyIterableError

{{ if eq $variant "Slice" }}
func CountIf{{ $variant }}[T comparable](haystack []T, unary_predicate func( needle T) bool) (count int, err error) {
{{ else if eq $variant "Map" }}
func CountIf{{ $variant }}[TKey comparable, TValue comparable](haystack map[TKey]TValue, unary_predicate func( needle TValue) bool) (count int, err error) {
{{ else if eq $variant "String" }}
func CountIf{{ $variant }}(haystack string, unary_predicate func( needle rune) bool) (count int, err error) {
{{ else if eq $variant "Iterator" }}
func CountIf{{ $variant }}[T comparable](haystack ds.ReadCompForIndexIterator[T], unary_predicate func( needle T) bool) (count int, err error) {
{{ end }}
    {{ if eq $variant "Iterator" }}
	if haystack.IsEnd() {
    {{ else }}
	if len(haystack) == 0 {
    {{ end }}
        err = EmptyIterableError{}
		return
	}

    {{ if eq $variant "Iterator" }}
    for haystack.Next() {
        value, _ := haystack.Get()
    {{ else }}
	for i, value := range haystack {
    {{ end }}
		if unary_predicate(value) {
            count++
        }
	}

    return
}
{{ end }}
{{ range $variant := . }}

// CountIfNot{{ $variant }} counts for how many val of container unary_predicate(val) == false.
//
// Possible Error values:
//    - EmptyIterableError

{{ if eq $variant "Slice" }}
func CountIfNot{{ $variant }}[T comparable](haystack []T, unary_predicate func( needle T) bool) (count int, err error) {
{{ else if eq $variant "Map" }}
func CountIfNot{{ $variant }}[TKey comparable, TValue comparable](haystack map[TKey]TValue, unary_predicate func( needle TValue) bool) (count int, err error) {
{{ else if eq $variant "String" }}
func CountIfNot{{ $variant }}(haystack string, unary_predicate func( needle rune) bool) (count int, err error) {
{{ else if eq $variant "Iterator" }}
func CountIfNot{{ $variant }}[T comparable](haystack ds.ReadCompForIndexIterator[T], unary_predicate func( needle T) bool) (count int, err error) {
{{ end }}
    {{ if eq $variant "Iterator" }}
	if haystack.IsEnd() {
    {{ else }}
	if len(haystack) == 0 {
    {{ end }}
        err = EmptyIterableError{}
		return
	}

    {{ if eq $variant "Iterator" }}
    for haystack.Next() {
        value, _ := haystack.Get()
    {{ else }}
	for i, value := range haystack {
    {{ end }}
		if !unary_predicate(value) {
            count++
        }
	}

    return
}
{{ end }}
{{ range $variant := . }}

{{ if ne $variant "Map" }}
// Mismatch{{ $variant }}Pred finds the first index i where binary_predicate(iterable1[i], iterable2[i] == false).
//
// Possible Error values:
//    - EmptyIterableError
//    - EqualIteratorsError
{{ if eq $variant "Slice" }}
func Mismatch{{ $variant }}Pred[T comparable](iterable1 []T, iterable2 []T, binary_predicate func(T, T) bool) (int, error) {
{{ else if eq $variant "String" }}
func Mismatch{{ $variant }}Pred(iterable1 string, iterable2 string, binary_predicate func(rune, rune) bool) (int, error) {
{{ else if eq $variant "Iterator" }}
func Mismatch{{ $variant }}Pred[T comparable](iterable1 ds.ReadCompForIndexIterator[T], iterable2 ds.ReadCompForIndexIterator[T], binary_predicate func(T, T) bool) (int, error) {
{{ end }}
    {{ if eq $variant "Iterator" }}
	if iterable1.IsEnd() || iterable2.IsEnd() {
    {{ else }}
	if len(iterable) == 0 || len(iterable2) == 0 {
    {{ end }}
        err = EmptyIterableError{}
		return
	}


	i := 0
    {{ if eq $variant "String" }}
    var curIt1 rune
    var curIt2 rune
    {{ else }}
    var curIt1 T
    var curIt2 T
    {{ end }}

    {{ if eq $variant "Iterable" }}
    for iterable1.Next() && iterable2.Next() {
        curIt1, _ = iterable1.Get()
        curIt2, _ = iterable2.Get()
    {{ else }}
	for ; i < min(len(iterable1), len(iterable2)); i++ {
        curIt1 = iterable1[i]
        curIt2 = iterable2[i]
    {{ end }}
		if !binary_predicate(curIt1, curIt2) {
            {{ if eq $variant "Iterable" }}
            index, _ := iterable1.Index()
            {{ else }}
			index =  i
            {{ end }}

			return
		}
	}


	err =  EqualIteratorsError{}

	return
}

{{ end }}
{{ end }}
{{ range $variant := . }}

{{ if ne $variant "Map" }}
// AdjacentFind{{ $variant }}Pred finds the first index i where binary_predicate(container[i], container[i+1]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError
{{ if eq $variant "Slice" }}
func AdjacentFind{{ $variant }}Pred[T comparable](haystack []T, binary_predicate func(T, T) bool) (int, error) {
{{ else if eq $variant "String" }}
func AdjacentFind{{ $variant }}Pred(haystack string, binary_predicate func(rune, rune) bool) (int, error) {
{{ else if eq $variant "Iterator" }}
func AdjacentFind{{ $variant }}Pred[T comparable](haystack ds.ReadCompForIndexIterator[T], binary_predicate func(T, T) bool) (int, error) {
{{ end }}
    {{ if eq $variant "Iterator" }}
	if haystack.IsEnd() {
    {{ else }}
	if len(haystack) == 0 {
    {{ end }}
        err = EmptyIterableError{}
		return
	}


	i := 0
    {{ if eq $variant "String" }}
    var cur rune
    var prev rune
    {{ else }}
    var cur T
    var prev T
    {{ end }}

    {{ if eq $variant "Iterable" }}
    prev, _ = haystack.Get()
    haystack.Next()

    for haystack.Next() {
        prev = cur
        cur, _ = haystack.Get()
    {{ else }}
    prev = haystack[0]

    for i := 1; i <= len(haystack); i++ {
        prev = cur
        cur = haystack[i]
    {{ end }}
		if binary_predicate(curIt1, curIt2) {
            {{ if eq $variant "Iterable" }}
            index, _ := haystack.Index()
            {{ else }}
			index =  i
            {{ end }}

			return
		}
	}


	err = ElementNotFoundError{}

	return
}

{{ end }}
{{ end }}
{{ range $variant := . }}

{{ if eq (ne $variant "Iterator") (ne $variant "Map") }}
// TakeWhile{{ $variant }} returns a copy of original until the first element not satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
{{ if eq $variant "Slice" }}
func TakeWhile{{ $variant }}[T comparable](haystack []T, unary_predicate func( needle T) bool) (out []T, err error) {
{{ else if eq $variant "String" }}
func TakeWhile{{ $variant }}(haystack string, unary_predicate func( needle rune) bool) (out string, err error) {
{{ end }}
    {{ if eq $variant "Iterator" }}
    if haystack.IsEnd() {
    {{ else }}
	if len(original) == 0 {
    {{ end }}
		err =  EmptyIterableError{}

		return
	}

    {{ if eq $variant "Slice" }}
    out = make([]T, 0, len(haystack))
    {{ end }}

    {{ if eq $variant "Iterator" }}
    var value T

    for haystack.Next() {
        value, _ = haystack.Get()
    {{ else }}
	for _, value := range original {
    {{ end }}
		if !unary_predicate(value) {
			return
		}

        {{ if eq $variant "Slice" }}
		out = append(out, value)
        {{ else if eq $variant "String" }}
        out += value
        {{ end }}
	}

	return
}

{{ end }}
{{ end }}
{{ range $variant := . }}

{{ if eq (ne $variant "Iterator") (ne $variant "Map") }}
// DropWhile{{ $variant }} returns a copy of original starting from first element not satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
{{ if eq $variant "Slice" }}
func TakeWhile{{ $variant }}[T comparable](haystack []T, unary_predicate func( needle T) bool) (out []T, err error) {
{{ else if eq $variant "String" }}
func TakeWhile{{ $variant }}(haystack string, unary_predicate func( needle rune) bool) (out string, err error) {
{{ end }}
    {{ if eq $variant "Iterator" }}
    if haystack.IsEnd() {
    {{ else }}
	if len(original) == 0 {
    {{ end }}
		err =  EmptyIterableError{}

		return
	}

    {{ if eq $variant "Slice" }}
    out = make([]T, 0, len(haystack))
    {{ end }}

    {{ if eq $variant "Iterator" }}
    var value T

    for haystack.Next() {
        value, _ = haystack.Get()
    {{ else }}
	for _, value := range original {
    {{ end }}
		if !unary_predicate(value) {
			break
		}

        {{ if eq $variant "Slice" }}
		out = append(out, value)
        {{ else if eq $variant "String" }}
        out += value
        {{ end }}
	}

	return
}

{{ end }}

{{ end }}
// {{ range $variant := . }}
//
// // DropWhileMap returns a copy of original starting from the first value not satisfying unary_predicate(value) == true).
// // Note that the iteration order of a map is not stable.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func DropWhileMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
// 	var zeroVal map[TKey]TValue
//
// 	if len(original) == 0 {
// 		index =  zeroVal
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	newContainer := make(map[TKey]TValue, len(original))
//
// 	isDropping := true
// 	for key, value := range original {
// 		if !(unary_predicate(value) && isDropping) {
// 			newContainer[key] = value
// 		}
// 	}
//
// 	index =  newContainer


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// CopyIf{{ $variant }} returns a copy of original with all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyIf{{ $variant }}[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var zeroVal []T

	if len(original) == 0 {
		index =  zeroVal
		err =  EmptyIterableError{}

		return
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if unary_predicate(value) {
			newContainer = append(newContainer, value)
		}
	}

	index =  newContainer


	return
}

{{ end }}
// {{ range $variant := . }}
//
// // CopyIfMap returns a copy of original with all key-value pairs satisfying unary_predicate(value) == true).
// // Note that the iteration order of a map is not stable.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func CopyIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
// 	var zeroVal map[TKey]TValue
//
// 	if len(original) == 0 {
// 		index =  zeroVal
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	newContainer := make(map[TKey]TValue, len(original))
//
// 	for key, value := range original {
// 		if unary_predicate(value) {
// 			newContainer[key] = value
// 		}
// 	}
//
// 	index =  newContainer


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// CopyReplaceIf{{ $variant }} returns a copy of original where each element satisfying unary_predicate(element) == true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIf{{ $variant }}[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var zeroVal []T

	if len(original) == 0 {
		index =  zeroVal
		err =  EmptyIterableError{}

		return
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if unary_predicate(value) {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	index =  newContainer


	return
}

{{ end }}
// {{ range $variant := . }}
//
// // CopyReplaceIfMap returns a copy of original where each value of a key-value pair satisfying unary_predicate(value) == true is replaced with replacement.
// // Note that the iteration order of a map is not stable.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func CopyReplaceIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool, replacement TValue) (map[TKey]TValue, error) {
// 	var zeroVal map[TKey]TValue
//
// 	if len(original) == 0 {
// 		index =  zeroVal
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	newContainer := make(map[TKey]TValue, len(original))
//
// 	for key, value := range original {
// 		if unary_predicate(value) {
// 			newContainer[key] = replacement
// 		} else {
// 			newContainer[key] = value
// 		}
// 	}
//
// 	index =  newContainer


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// CopyReplaceIfNot{{ $variant }} returns a copy of original where each element satisfying unary_predicate(element) != true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfNot{{ $variant }}[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var zeroVal []T

	if len(original) == 0 {
		index =  zeroVal
		err =  EmptyIterableError{}

		return
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if !unary_predicate(value) {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	index =  newContainer


	return
}

{{ end }}
// {{ range $variant := . }}
//
// // CopyReplaceIfNotMap returns a copy of original where each element satisfying unary_predicate(element) != true is replaced with replacement.
// // Note that the iteration order of a map is not stable.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func CopyReplaceIfNotMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool, replacement TValue) (map[TKey]TValue, error) {
// 	var zeroVal map[TKey]TValue
//
// 	if len(original) == 0 {
// 		index =  zeroVal
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	newContainer := make(map[TKey]TValue, len(original))
//
// 	for key, value := range original {
// 		if !unary_predicate(value) {
// 			newContainer[key] = replacement
// 		} else {
// 			newContainer[key] = value
// 		}
// 	}
//
// 	index =  newContainer


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// CopyExceptIf{{ $variant }} returns a copy of original without all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIf{{ $variant }}[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var zeroVal []T

	if len(original) == 0 {
		index =  zeroVal
		err =  EmptyIterableError{}

		return
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if !unary_predicate(value) {
			newContainer = append(newContainer, value)
		}
	}

	index =  newContainer


	return
}

{{ end }}
// {{ range $variant := . }}
//
// // CopyExceptIfMap returns a copy of original without all key-value pairs satisfying unary_predicate(value) == true).
// // Note that the iteration order of a map is not stable.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func CopyExceptIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
// 	var zeroVal map[TKey]TValue
//
// 	if len(original) == 0 {
// 		index =  zeroVal
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	newContainer := make(map[TKey]TValue, len(original))
//
// 	for key, value := range original {
// 		if !unary_predicate(value) {
// 			newContainer[key] = value
// 		}
// 	}
//
// 	index =  newContainer


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// CopyExceptIfNot{{ $variant }} returns a copy of original without all element satisfying unary_predicate(element) == false).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfNot{{ $variant }}[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var zeroVal []T

	if len(original) == 0 {
		index =  zeroVal
		err =  EmptyIterableError{}

		return
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if unary_predicate(value) {
			newContainer = append(newContainer, value)
		}
	}

	index =  newContainer


	return
}

{{ end }}
{{ range $variant := . }}
//
// // CopyExceptIfNotMap returns a copy of original without all key-value pairs satisfying unary_predicate(value) == false).
// // Note that the iteration order of a map is not stable.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func CopyExceptIfNotMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
// 	var zeroVal map[TKey]TValue
//
// 	if len(original) == 0 {
// 		index =  zeroVal
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	newContainer := make(map[TKey]TValue, len(original))
//
// 	for key, value := range original {
// 		if unary_predicate(value) {
// 			newContainer[key] = value
// 		}
// 	}
//
// 	index =  newContainer


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// Fill{{ $variant }} fills the array pointed to by arr with filler.
// all indices in the range [0, cap(*arr)[ are filled regardless of what len(*arr) is.
func Fill{{ $variant }}[T any](arr *[]T, filler T) []T {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	for i := range *arr {
		(*arr)[i] = filler
	}

	n_unfilled := cap(*arr) - len(*arr)

	for i := 0; i < n_unfilled; i++ {
		*arr = append(*arr, filler)
	}

	return *arr
}

{{ end }}
// {{ range $variant := . }}
//
// // TransformMap applies transformer(value) for all key-value pairs in container and stores them at container[key].
// // Note that the iteration order of a map is not stable.
// // Errors returned by transformer are propagated to the caller of TransformMap.
// //
// // Possible Error values:
// //    - EmptyIterableError
// //    - ExecutionError
// func TransformMap[TKey comparable, TValue comparable](container map[TKey]TValue, transformer func(TValue) (TValue, error)) error {
//
// 	if len(container) == 0 {
// 		return EmptyIterableError{}
// 	}
//
// 	for key := range container {
// 		newValue, err := transformer(container[key])
// 		if  {
// 			index =  ExecutionError[TKey
 			err =  TValue]{BadItemIndex: key, BadItem: container[key], Inner: err}

 			return
// 		}
//
// 		container[key] = newValue
// 	}
//
// 	return nil
// }
//
// {{ end }}
{{ range $variant := . }}

// Transform{{ $variant }} applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
// Errors returned by transformer are propagated to the caller of Transform{{ $variant }}.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func Transform{{ $variant }}[T any](container []T, transformer func(*T) error) error {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		err := transformer(&container[i])
		if  {
			index =  ExecutionError[int
			err =  T]{BadItemIndex: i, BadItem: value, Inner: err}

			return
		}
	}

	return nil
}

{{ end }}
// {{ range $variant := . }}
//
// // TransformMapUnsafe applies transformer(value) for all key-value pairs in container and stores them at container[key].
// // Note that the iteration order of a map is not stable.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func TransformMapUnsafe[TKey comparable, TValue comparable](container map[TKey]TValue, transformer func(TValue) TValue) error {
// 	if len(container) == 0 {
// 		return EmptyIterableError{}
// 	}
//
// 	for key := range container {
// 		container[key] = transformer(container[key])
// 	}
//
// 	return nil
// }
//
// {{ end }}
{{ range $variant := . }}

// Transform{{ $variant }}Unsafe applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
//
// Possible Error values:
//    - EmptyIterableError
func Transform{{ $variant }}Unsafe[T any](container []T, transformer func(*T)) error {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i := range container {
		transformer(&container[i])
	}

	return nil
}

{{ end }}
// {{ range $variant := . }}
//
// // TransformCopyMap applies transformer(value) for all key-value pairs in container and and returns the newly created container.
// // Note that the iteration order of a map is not stable.
// // Note that the transformer can return a different type than it's input.
// // Errors returned by transformer are propagated to the caller of TransformCopyMap.
// //
// // Possible Error values:
// //    - EmptyIterableError
// //    - ExecutionError
// func TransformCopyMap[TKey comparable, TValue comparable, TValueOut any](container map[TKey]TValue, transformer func(TValue) (TValueOut, error)) (map[TKey]TValueOut, error) {
// 	res := make(map[TKey]TValueOut)
//
// 	if len(container) == 0 {
// 		index =  res
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	for key := range container {
// 		newValue, err := transformer(container[key])
// 		if  {
// 			index =  res
 			err =  ExecutionError[TKey, TValue]{BadItemIndex: key, BadItem: container[key], Inner: err}

 			return
// 		}
//
// 		res[key] = newValue
// 	}
//
// 	index =  res


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// TransformCopy{{ $variant }} applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
// Errors returned by transformer are propagated to the caller of TransformCopy{{ $variant }}.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformCopy{{ $variant }}[T any, TOut any](container []T, transformer func(T) (TOut, error)) ([]TOut, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	res := make([]TOut, 0, len(container))

	if len(container) == 0 {
		index =  res
		err =  EmptyIterableError{}

		return
	}

	for i, value := range container {
		newVal, err := transformer(container[i])
		if  {
			index =  res
			err =  ExecutionError[int, T]{BadItemIndex: i, BadItem: value, Inner: err}

			return
		}

		res = append(res, newVal)
	}

	index =  res


	return
}

{{ end }}
// {{ range $variant := . }}
//
// // TransformCopyMapUnsafe applies transformer(value) for all key-value pairs in container and and returns the newly created container.
// // Note that the transformer can return a different type than it's input.
// // Note that the iteration order of a map is not stable.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func TransformCopyMapUnsafe[TKey comparable, TValue comparable, TValueOut any](container map[TKey]TValue, transformer func(TValue) TValueOut) (map[TKey]TValueOut, error) {
// 	res := make(map[TKey]TValueOut)
//
// 	if len(container) == 0 {
// 		index =  res
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	for key := range container {
// 		res[key] = transformer(container[key])
// 	}
//
// 	index =  res


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// TransformCopy{{ $variant }}Unsafe applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
//
// Possible Error values:
//    - EmptyIterableError
func TransformCopy{{ $variant }}Unsafe[T any, TOut any](container []T, transformer func(T) TOut) ([]TOut, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	res := make([]TOut, 0, len(container))

	if len(container) == 0 {
		index =  res
		err =  EmptyIterableError{}

		return
	}

	for i := range container {
		res = append(res, transformer(container[i]))
	}

	index =  res


	return
}

{{ end }}
{{ range $variant := . }}

// Min{{ $variant }}Pred finds the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func Min{{ $variant }}Pred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var min T

	if len(haystack) == 0 {
		index =  min
		err =  EmptyIterableError{}

		return
	}

	min = haystack[0]
	for _, val := range haystack {
		if binary_predicate(val, min) {
			min = val
		}
	}

	index =  min


	return
}

{{ end }}
// {{ range $variant := . }}
//
// // MinMapPred finds the smallest value in haystack.
// // The elements are compared with binary_predicate.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func MinMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate func(TValue, TValue) bool) (TValue, error) {
// 	var min TValue
//
// 	for _, val := range haystack {
// 		min = val
// 		break
// 	}
//
// 	if len(haystack) == 0 {
// 		index =  min
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	for _, val := range haystack {
// 		if binary_predicate(val, min) {
// 			min = val
// 		}
// 	}
//
// 	index =  min


 	return
// }
//
// {{ end }}
// {{ range $variant := . }}

// Max{{ $variant }}Pred finds the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func Max{{ $variant }}Pred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var max T

	if len(haystack) == 0 {
		index =  max
		err =  EmptyIterableError{}

		return
	}

	max = haystack[0]
	for _, val := range haystack {
		if binary_predicate(val, max) {
			max = val
		}
	}

	index =  max


	return
}

// {{ end }}
// {{ range $variant := . }}
//
// // MaxMapPred finds the largest value in haystack.
// // The elements are compared with binary_predicate.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func MaxMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate func(TValue, TValue) bool) (TValue, error) {
// 	var max TValue
//
// 	for _, val := range haystack {
// 		max = val
// 		break
// 	}
//
// 	if len(haystack) == 0 {
// 		index =  max
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	for _, val := range haystack {
// 		if binary_predicate(val, max) {
// 			max = val
// 		}
// 	}
//
// 	index =  max


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// MinMax{{ $variant }}Pred finds the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMax{{ $variant }}Pred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (T, T, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var min T
	var max T

	if len(haystack) == 0 {
		index =  min
		err =  max, EmptyIterableError{}

		return
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

	index =  min


	return
}

{{ end }}
// {{ range $variant := . }}
//
// // MinMaxMapPred finds the smallest and largest value in haystack.
// // The elements are compared with binary_predicate.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func MinMaxMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate_min func(TValue, TValue) bool, binary_predicate_max func(TValue, TValue) bool) (TValue, TValue, error) {
// 	var min TValue
// 	var max TValue
//
// 	for _, val := range haystack {
// 		min = val
// 		max = val
// 		break
// 	}
//
// 	if len(haystack) == 0 {
// 		index =  min
 		err =  max, EmptyIterableError{}

 		return
// 	}
//
// 	for _, val := range haystack {
// 		if binary_predicate_min(val, min) {
// 			min = val
// 		}
// 		if binary_predicate_max(val, max) {
// 			max = val
// 		}
// 	}
//
// 	index =  min


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// MinElement{{ $variant }}Pred finds the index of the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinElement{{ $variant }}Pred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var min int

	if len(haystack) == 0 {
		index =  min
		err =  EmptyIterableError{}

		return
	}

	min = 0
	for key, val := range haystack {
		if binary_predicate(val, haystack[min]) {
			min = key
		}
	}

	index =  min


	return
}

{{ end }}
// {{ range $variant := . }}
//
// // MinElementMapPred finds the index of the smallest value in haystack.
// // The elements are compared with binary_predicate.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func MinElementMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate func(TValue, TValue) bool) (TKey, error) {
// 	var min TKey
//
// 	for key := range haystack {
// 		min = key
// 		break
// 	}
//
// 	if len(haystack) == 0 {
// 		index =  min
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	for key, val := range haystack {
// 		if binary_predicate(val, haystack[min]) {
// 			min = key
// 		}
// 	}
//
// 	index =  min


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// MaxElement{{ $variant }}Pred finds the index of the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElement{{ $variant }}Pred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var max int

	if len(haystack) == 0 {
		index =  max
		err =  EmptyIterableError{}

		return
	}

	max = 0
	for key, val := range haystack {
		if binary_predicate(val, haystack[max]) {
			max = key
		}
	}

	index =  max


	return
}

{{ end }}
// {{ range $variant := . }}
//
// // MaxElementMapPred finds the index of the largest value in haystack.
// // The elements are compared with binary_predicate.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func MaxElementMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate func(TValue, TValue) bool) (TKey, error) {
// 	var max TKey
//
// 	for key := range haystack {
// 		max = key
// 		break
// 	}
//
// 	if len(haystack) == 0 {
// 		index =  max
 		err =  EmptyIterableError{}

 		return
// 	}
//
// 	for key, val := range haystack {
// 		if binary_predicate(val, haystack[max]) {
// 			max = key
// 		}
// 	}
//
// 	index =  max


 	return
// }
//
// {{ end }}
{{ range $variant := . }}

// MinMaxElement{{ $variant }}Pred finds the index of the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElement{{ $variant }}Pred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (int, int, error) {
{{ if eq $variant "Slice" }}
{{ else if eq $variant "Map" }}
{{ else if eq $variant "String" }}
{{ else if eq $variant "Iterator" }}
{{ end }}
	var min int
	var max int

	if len(haystack) == 0 {
		index =  min
		err =  max, EmptyIterableError{}

		return
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

	index =  min


	return
}

{{ end }}
// {{ range $variant := . }}
//
// // MinMaxElementMapPred finds the index of the smallest and largest value in haystack.
// // The elements are compared with binary_predicate.
// //
// // Possible Error values:
// //    - EmptyIterableError
// func MinMaxElementMapPred[TKey comparable, TValue constraints.Ordered](haystack map[TKey]TValue, binary_predicate_min func(TValue, TValue) bool, binary_predicate_max func(TValue, TValue) bool) (TKey, TKey, error) {
// 	var min TKey
// 	var max TKey
//
// 	for key := range haystack {
// 		min = key
// 		max = key
// 		break
// 	}
//
// 	if len(haystack) == 0 {
// 		index =  min
 		err =  max, EmptyIterableError{}

 		return
// 	}
//
// 	for i, val := range haystack {
// 		if binary_predicate_min(val, haystack[min]) {
// 			min = i
// 		}
// 		if binary_predicate_max(val, haystack[max]) {
// 			max = i
// 		}
// 	}
//
// 	index =  min


 	return
// }
//
// {{ end }}
