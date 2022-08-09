// Package goaoi implements conventient algorithms for processing iterables.
// It is inspired by the algorithm header from the C++ standard template library (STL for short).
package goaoi

import (
	"golang.org/x/exp/constraints"
    "github.com/JonasMuehlmann/datastructures.go/ds"
)









// FindIfSlice finds the first index i where unary_predicate(haystack[i]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func FindIfSlice[T comparable](haystack []T, unary_predicate func( needle T) bool) (index int, err error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}

    
	for i, value := range haystack {
    
		if unary_predicate(value) {
        
            index = i
        
			return
		}
	}

	err = ElementNotFoundError{}

    return
}




// FindIfMap finds the first index i where unary_predicate(haystack[i]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func FindIfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, unary_predicate func( needle TValue) bool) (index TKey, err error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}

    
	for i, value := range haystack {
    
		if unary_predicate(value) {
        
            index = i
        
			return
		}
	}

	err = ElementNotFoundError{}

    return
}




// FindIfString finds the first index i where unary_predicate(haystack[i]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func FindIfString(haystack string, unary_predicate func( needle rune) bool) (index int, err error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}

    
	for i, value := range haystack {
    
		if unary_predicate(value) {
        
            index = i
        
			return
		}
	}

	err = ElementNotFoundError{}

    return
}




// FindIfIterator finds the first index i where unary_predicate(haystack[i]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func FindIfIterator[T comparable](haystack ds.ReadCompForIndexIterator[T], unary_predicate func( needle T) bool) (index int, err error) {

    
	if haystack.IsEnd() {
    
        err = EmptyIterableError{}
		return
	}

    
    for haystack.Next() {
        value, _ := haystack.Get()
    
		if unary_predicate(value) {
        
            index, _ = haystack.Index()
        
			return
		}
	}

	err = ElementNotFoundError{}

    return
}





// FindEndSlicePred finds the beginning of the last occurrence of sub in super.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func FindEndSlicePred[T comparable](super []T, sub []T, binary_predicate func(T, T) bool) (index int, err error) {

    
	if len(haystack) == 0 || len(needles) == 0 {
    
		err =  EmptyIterableError{}

		return
	}

    
    lenSub := len(sub)
    lenSiper := len(super)
    

    
    var curSub T
    var curSuper T
    


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








// FindEndStringPred finds the beginning of the last occurrence of sub in super.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func FindEndStringPred(super string, sub string, binary_predicate func(string, string) bool) (index int, err error) {

    
	if len(haystack) == 0 || len(needles) == 0 {
    
		err =  EmptyIterableError{}

		return
	}

    
    lenSub := len(sub)
    lenSiper := len(super)
    

    
    var curSub rune
    var curSUper rune
    


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





// FindEndIteratorPred finds the beginning of the last occurrence of sub in super.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func FindEndIteratorPred[T comparable](super ds.ReadCompForIndexIterator[T], sub ds.ReadCompForIndexIterator[T], binary_predicate func(T, T) bool) (index int, err error) {

    
    if haystack.IsEnd() {
    
		err =  EmptyIterableError{}

		return
	}

    
    lenSub := super.Size()
    lenSiper := sub.Size()
    

    
    var curSub T
    var curSuper T
    


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





// FindFirstOfSlicePred finds the first index where an element of haystack is equal to any element in needles.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func FindFirstOfSlicePred[T comparable](haystack []T, needles []T, binary_predicate func(T, T) bool) (index int, err error) {

    
	if len(haystack) == 0 || len(needles) == 0 {
    

		err =  EmptyIterableError{}

		return
	}

    
    var curHaystack T
    

    
	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
    
			if binary_predicate(haystackValue, needleValue) {
				index =  i


				return
			}
		}
	}


	err =  ElementNotFoundError{}

	return
}



// FindFirstOfMapPred finds the first index where an element of haystack is equal to any element in needles.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func FindFirstOfMapPred[TKey comparable, TValue comparable](haystack map[TKey]TValue, needles []TValue, binary_predicate func(TValue, TValue) bool) (index TKey, err error) {

    
	if len(haystack) == 0 || len(needles) == 0 {
    

		err =  EmptyIterableError{}

		return
	}

    
    var curHaystack T
    

    
	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
    
			if binary_predicate(haystackValue, needleValue) {
				index =  i


				return
			}
		}
	}


	err =  ElementNotFoundError{}

	return
}



// FindFirstOfStringPred finds the first index where an element of haystack is equal to any element in needles.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func FindFirstOfStringPred(haystack string, needles string, binary_predicate func(rune, rune) bool) (index int, err error) {

    
	if len(haystack) == 0 || len(needles) == 0 {
    

		err =  EmptyIterableError{}

		return
	}

    
    var curHaystack rune
    

    
	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
    
			if binary_predicate(haystackValue, needleValue) {
				index =  i


				return
			}
		}
	}


	err =  ElementNotFoundError{}

	return
}



// FindFirstOfIteratorPred finds the first index where an element of haystack is equal to any element in needles.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

    // TODO: This should take an iterator, which gets copeid before each run through
func FindFirstOfIteratorPred[T comparable](haystack ds.ReadCompForIndexIterator[T], needles []T, binary_predicate func(T, T) bool) (index int,err  error) {

    
    if haystack.IsEnd() {
    

		err =  EmptyIterableError{}

		return
	}

    
    var curHaystack T
    

    
    for haystack.Next() {
		for _, needleValue := range needles {
            curHaystack, _ = haystack.Get()
    
			if binary_predicate(haystackValue, needleValue) {
				index =  i


				return
			}
		}
	}


	err =  ElementNotFoundError{}

	return
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
        

        
        err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
        

			return
		}
	}

	return nil
}



// AllOfMap checks that unary_predicate(val) == true for ALL val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError

func AllOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {

    
	if len(container) == 0 {
    
		return EmptyIterableError{}
	}

    
	for i, value := range container {
    

		if !unary_predicate(value) {
        

        
        err =  ComparisonError[TKey, TValue]{BadItemIndex: i, BadItem: value}
        

			return
		}
	}

	return nil
}



// AllOfString checks that unary_predicate(val) == true for ALL val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError

func AllOfString(container string, unary_predicate func(rune) bool) error {

    
	if len(container) == 0 {
    
		return EmptyIterableError{}
	}

    
	for i, value := range container {
    

		if !unary_predicate(value) {
        

        
        err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
        

			return
		}
	}

	return nil
}



// AllOfIterator checks that unary_predicate(val) == true for ALL val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError

func AllOfIterator[T any](container ds.ReadCompForIndexIterator[T], unary_predicate func(T) bool) error {

    
    if container.IsEnd() {
    
		return EmptyIterableError{}
	}

    
    var value T

    for container.Next() {
        value, _ = container.Get()
    

		if !unary_predicate(value) {
        
        i, _ := container.Index()
        

        
        err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
        

			return
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

    
	for i, value := range container {
    

		if unary_predicate(value) {
            return nil
		}
	}

    

    
    err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
    

    return
}



// AnyOfMap checks that unary_predicate(val) == true for ANY val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func AnyOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {

    
	if len(container) == 0 {
    
		return EmptyIterableError{}
	}

    
	for i, value := range container {
    

		if unary_predicate(value) {
            return nil
		}
	}

    

    
    err =  ComparisonError[TKey, TValue]{BadItemIndex: i, BadItem: value}
    

    return
}



// AnyOfString checks that unary_predicate(val) == true for ANY val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func AnyOfString(container string, unary_predicate func(rune) bool) error {

    
	if len(container) == 0 {
    
		return EmptyIterableError{}
	}

    
	for i, value := range container {
    

		if unary_predicate(value) {
            return nil
		}
	}

    

    
    err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
    

    return
}



// AnyOfIterator checks that unary_predicate(val) == true for ANY val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func AnyOfIterator[T any](container ds.ReadCompForIndexIterator[T], unary_predicate func(T) bool) error {

    
    if container.IsEnd() {
    
		return EmptyIterableError{}
	}

    
    var value T

    for container.Next() {
        value, _ = container.Get()
    

		if unary_predicate(value) {
            return nil
		}
	}

    
    i, _ := container.Index()
    

    
    err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
    

    return
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
        

        
        err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
        

			return
		}
	}

	return nil
}



// NoneOfMap checks that unary_predicate(val) == true for NO val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError

func NoneOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {

    
	if len(container) == 0 {
    
		return EmptyIterableError{}
	}

    
	for i, value := range container {
    

		if unary_predicate(value) {
        

        
        err =  ComparisonError[TKey, TValue]{BadItemIndex: i, BadItem: value}
        

			return
		}
	}

	return nil
}



// NoneOfString checks that unary_predicate(val) == true for NO val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError

func NoneOfString(container string, unary_predicate func(rune) bool) error {

    
	if len(container) == 0 {
    
		return EmptyIterableError{}
	}

    
	for i, value := range container {
    

		if unary_predicate(value) {
        

        
        err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
        

			return
		}
	}

	return nil
}



// NoneOfIterator checks that unary_predicate(val) == true for NO val in container.
//
// Possible Error values:
//    - EmptyIterableError
//    - ComparisonError

func NoneOfIterator[T any](container ds.ReadCompForIndexIterator[T], unary_predicate func(T) bool) error {

    
    if container.IsEnd() {
    
		return EmptyIterableError{}
	}

    
    var value T

    for container.Next() {
        value, _ = container.Get()
    

		if unary_predicate(value) {
        
        i, _ := container.Index()
        

        
        err =  ComparisonError[int, T]{BadItemIndex: i, BadItem: value}
        

			return
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
		if err != nil{
        

        
        err = ExecutionError[int, T]{BadItemIndex: i, BadItem: value, Inner: err}
        

			return
		}
	}


	return nil
}



// ForeachMap executes unary_func(val) for each val in container.
// Errors returned by unary_func are propagated to the caller of ForeachMap.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError

func ForeachMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_func func(TValue) error) error {

    
	if len(container) == 0 {
    
		return EmptyIterableError{}
	}


 
	for i, value := range container {
    

        err := unary_func(value)
		if err != nil{
        

        
        err = ExecutionError[TKey, TValue]{BadItemIndex: i, BadItem: value, Inner: err}
        

			return
		}
	}


	return nil
}



// ForeachString executes unary_func(val) for each val in container.
// Errors returned by unary_func are propagated to the caller of ForeachString.
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
		if err != nil{
        

        
        err = ExecutionError[int, T]{BadItemIndex: i, BadItem: value, Inner: err}
        

			return
		}
	}


	return nil
}



// ForeachIterator executes unary_func(val) for each val in container.
// Errors returned by unary_func are propagated to the caller of ForeachIterator.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError

func ForeachIterator[T any](container ds.ReadCompForIndexIterator[T], unary_func func(T) error) error {

    
    if container.IsEnd() {
    
		return EmptyIterableError{}
	}


 
    var value T

    for container.Next() {
        value, _ = container.Get()
    

        err := unary_func(value)
		if err != nil{
        
        i, _ := container.Index()
        

        
        err = ExecutionError[int, T]{BadItemIndex: i, BadItem: value, Inner: err}
        

			return
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


 
	for i, value := range container {
    
        unary_func(value)
    }


	return nil
}



// ForeachMapUnsafe executes unary_func(val) for each val in container.
//
// Possible Error values:
//    - EmptyIterableError

func ForeachMapUnsafe[TKey comparable, TValue comparable](container map[TKey]TValue, unary_func func(TValue)) error {

    
	if len(container) == 0 {
    
		return EmptyIterableError{}
	}


 
	for i, value := range container {
    
        unary_func(value)
    }


	return nil
}



// ForeachStringUnsafe executes unary_func(val) for each val in container.
//
// Possible Error values:
//    - EmptyIterableError

func ForeachStringUnsafe(container string, unary_func func(rune)) error {

    
	if len(container) == 0 {
    
		return EmptyIterableError{}
	}


 
	for i, value := range container {
    
        unary_func(value)
    }


	return nil
}



// ForeachIteratorUnsafe executes unary_func(val) for each val in container.
//
// Possible Error values:
//    - EmptyIterableError

func ForeachIteratorUnsafe[T any](container ds.ReadCompForIndexIterator[T], unary_func func(T)) error {

    
    if container.IsEnd() {
    
		return EmptyIterableError{}
	}


 
    var value T

    for container.Next() {
        value, _ = container.Get()
    
        unary_func(value)
    }


	return nil
}




// CountIfSlice counts for how many val of container unary_predicate(val) == true.
//
// Possible Error values:
//    - EmptyIterableError


func CountIfSlice[T comparable](haystack []T, unary_predicate func( needle T) bool) (count int, err error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}

    
	for i, value := range haystack {
    
		if unary_predicate(value) {
            count++
        }
	}

    return
}


// CountIfMap counts for how many val of container unary_predicate(val) == true.
//
// Possible Error values:
//    - EmptyIterableError


func CountIfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, unary_predicate func( needle TValue) bool) (count int, err error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}

    
	for i, value := range haystack {
    
		if unary_predicate(value) {
            count++
        }
	}

    return
}


// CountIfString counts for how many val of container unary_predicate(val) == true.
//
// Possible Error values:
//    - EmptyIterableError


func CountIfString(haystack string, unary_predicate func( needle rune) bool) (count int, err error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}

    
	for i, value := range haystack {
    
		if unary_predicate(value) {
            count++
        }
	}

    return
}


// CountIfIterator counts for how many val of container unary_predicate(val) == true.
//
// Possible Error values:
//    - EmptyIterableError


func CountIfIterator[T comparable](haystack ds.ReadCompForIndexIterator[T], unary_predicate func( needle T) bool) (count int, err error) {

    
	if haystack.IsEnd() {
    
        err = EmptyIterableError{}
		return
	}

    
    for haystack.Next() {
        value, _ := haystack.Get()
    
		if unary_predicate(value) {
            count++
        }
	}

    return
}



// CountIfNotSlice counts for how many val of container unary_predicate(val) == false.
//
// Possible Error values:
//    - EmptyIterableError


func CountIfNotSlice[T comparable](haystack []T, unary_predicate func( needle T) bool) (count int, err error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}

    
	for i, value := range haystack {
    
		if !unary_predicate(value) {
            count++
        }
	}

    return
}


// CountIfNotMap counts for how many val of container unary_predicate(val) == false.
//
// Possible Error values:
//    - EmptyIterableError


func CountIfNotMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, unary_predicate func( needle TValue) bool) (count int, err error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}

    
	for i, value := range haystack {
    
		if !unary_predicate(value) {
            count++
        }
	}

    return
}


// CountIfNotString counts for how many val of container unary_predicate(val) == false.
//
// Possible Error values:
//    - EmptyIterableError


func CountIfNotString(haystack string, unary_predicate func( needle rune) bool) (count int, err error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}

    
	for i, value := range haystack {
    
		if !unary_predicate(value) {
            count++
        }
	}

    return
}


// CountIfNotIterator counts for how many val of container unary_predicate(val) == false.
//
// Possible Error values:
//    - EmptyIterableError


func CountIfNotIterator[T comparable](haystack ds.ReadCompForIndexIterator[T], unary_predicate func( needle T) bool) (count int, err error) {

    
	if haystack.IsEnd() {
    
        err = EmptyIterableError{}
		return
	}

    
    for haystack.Next() {
        value, _ := haystack.Get()
    
		if !unary_predicate(value) {
            count++
        }
	}

    return
}




// MismatchSlicePred finds the first index i where binary_predicate(iterable1[i], iterable2[i] == false).
//
// Possible Error values:
//    - EmptyIterableError
//    - EqualIteratorsError

func MismatchSlicePred[T comparable](iterable1 []T, iterable2 []T, binary_predicate func(T, T) bool) (int, error) {

    
	if len(iterable) == 0 || len(iterable2) == 0 {
    
        err = EmptyIterableError{}
		return
	}


	i := 0
    
    var curIt1 T
    var curIt2 T
    

    
	for ; i < min(len(iterable1), len(iterable2)); i++ {
        curIt1 = iterable1[i]
        curIt2 = iterable2[i]
    
		if !binary_predicate(curIt1, curIt2) {
            
			index =  i
            

			return
		}
	}


	err =  EqualIteratorsError{}

	return
}








// MismatchStringPred finds the first index i where binary_predicate(iterable1[i], iterable2[i] == false).
//
// Possible Error values:
//    - EmptyIterableError
//    - EqualIteratorsError

func MismatchStringPred(iterable1 string, iterable2 string, binary_predicate func(rune, rune) bool) (int, error) {

    
	if len(iterable) == 0 || len(iterable2) == 0 {
    
        err = EmptyIterableError{}
		return
	}


	i := 0
    
    var curIt1 rune
    var curIt2 rune
    

    
	for ; i < min(len(iterable1), len(iterable2)); i++ {
        curIt1 = iterable1[i]
        curIt2 = iterable2[i]
    
		if !binary_predicate(curIt1, curIt2) {
            
			index =  i
            

			return
		}
	}


	err =  EqualIteratorsError{}

	return
}





// MismatchIteratorPred finds the first index i where binary_predicate(iterable1[i], iterable2[i] == false).
//
// Possible Error values:
//    - EmptyIterableError
//    - EqualIteratorsError

func MismatchIteratorPred[T comparable](iterable1 ds.ReadCompForIndexIterator[T], iterable2 ds.ReadCompForIndexIterator[T], binary_predicate func(T, T) bool) (int, error) {

    
	if iterable1.IsEnd() || iterable2.IsEnd() {
    
        err = EmptyIterableError{}
		return
	}


	i := 0
    
    var curIt1 T
    var curIt2 T
    

    
	for ; i < min(len(iterable1), len(iterable2)); i++ {
        curIt1 = iterable1[i]
        curIt2 = iterable2[i]
    
		if !binary_predicate(curIt1, curIt2) {
            
			index =  i
            

			return
		}
	}


	err =  EqualIteratorsError{}

	return
}






// AdjacentFindSlicePred finds the first index i where binary_predicate(container[i], container[i+1]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func AdjacentFindSlicePred[T comparable](haystack []T, binary_predicate func(T, T) bool) (int, error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}


	i := 0
    
    var cur T
    var prev T
    

    
    prev = haystack[0]

    for i := 1; i <= len(haystack); i++ {
        prev = cur
        cur = haystack[i]
    
		if binary_predicate(curIt1, curIt2) {
            
			index =  i
            

			return
		}
	}


	err = ElementNotFoundError{}

	return
}








// AdjacentFindStringPred finds the first index i where binary_predicate(container[i], container[i+1]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func AdjacentFindStringPred(haystack string, binary_predicate func(rune, rune) bool) (int, error) {

    
	if len(haystack) == 0 {
    
        err = EmptyIterableError{}
		return
	}


	i := 0
    
    var cur rune
    var prev rune
    

    
    prev = haystack[0]

    for i := 1; i <= len(haystack); i++ {
        prev = cur
        cur = haystack[i]
    
		if binary_predicate(curIt1, curIt2) {
            
			index =  i
            

			return
		}
	}


	err = ElementNotFoundError{}

	return
}





// AdjacentFindIteratorPred finds the first index i where binary_predicate(container[i], container[i+1]) == true.
//
// Possible Error values:
//    - EmptyIterableError
//    - ElementNotFoundError

func AdjacentFindIteratorPred[T comparable](haystack ds.ReadCompForIndexIterator[T], binary_predicate func(T, T) bool) (int, error) {

    
	if haystack.IsEnd() {
    
        err = EmptyIterableError{}
		return
	}


	i := 0
    
    var cur T
    var prev T
    

    
    prev = haystack[0]

    for i := 1; i <= len(haystack); i++ {
        prev = cur
        cur = haystack[i]
    
		if binary_predicate(curIt1, curIt2) {
            
			index =  i
            

			return
		}
	}


	err = ElementNotFoundError{}

	return
}






// TakeWhileSlice returns a copy of original until the first element not satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError

func TakeWhileSlice[T comparable](haystack []T, unary_predicate func( needle T) bool) (out []T, err error) {

    
	if len(original) == 0 {
    
		err =  EmptyIterableError{}

		return
	}

    
    out = make([]T, 0, len(haystack))
    

    
	for _, value := range original {
    
		if !unary_predicate(value) {
			return
		}

        
		out = append(out, value)
        
	}

	return
}








// TakeWhileString returns a copy of original until the first element not satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError

func TakeWhileString(haystack string, unary_predicate func( needle rune) bool) (out string, err error) {

    
	if len(original) == 0 {
    
		err =  EmptyIterableError{}

		return
	}

    

    
	for _, value := range original {
    
		if !unary_predicate(value) {
			return
		}

        
        out += value
        
	}

	return
}








// DropWhileSlice returns a copy of original starting from first element not satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func DropWhileSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


	var zeroVal []T

	if len(original) == 0 {
		index =  zeroVal
		err =  EmptyIterableError{}

		return
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

	index =  newContainer


	return
}



// DropWhileMap returns a copy of original starting from first element not satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func DropWhileMap[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


	var zeroVal []T

	if len(original) == 0 {
		index =  zeroVal
		err =  EmptyIterableError{}

		return
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

	index =  newContainer


	return
}



// DropWhileString returns a copy of original starting from first element not satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func DropWhileString[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


	var zeroVal []T

	if len(original) == 0 {
		index =  zeroVal
		err =  EmptyIterableError{}

		return
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

	index =  newContainer


	return
}



// DropWhileIterator returns a copy of original starting from first element not satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func DropWhileIterator[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


	var zeroVal []T

	if len(original) == 0 {
		index =  zeroVal
		err =  EmptyIterableError{}

		return
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

	index =  newContainer


	return
}


// 
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
// 
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
// 
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
// 
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
// 


// CopyIfSlice returns a copy of original with all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyIfSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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



// CopyIfMap returns a copy of original with all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyIfMap[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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



// CopyIfString returns a copy of original with all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyIfString[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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



// CopyIfIterator returns a copy of original with all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyIfIterator[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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


// 
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
// 
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
// 
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
// 
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
// 


// CopyReplaceIfSlice returns a copy of original where each element satisfying unary_predicate(element) == true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfSlice[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {


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



// CopyReplaceIfMap returns a copy of original where each element satisfying unary_predicate(element) == true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfMap[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {


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



// CopyReplaceIfString returns a copy of original where each element satisfying unary_predicate(element) == true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfString[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {


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



// CopyReplaceIfIterator returns a copy of original where each element satisfying unary_predicate(element) == true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfIterator[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {


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


// 
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
// 
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
// 
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
// 
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
// 


// CopyReplaceIfNotSlice returns a copy of original where each element satisfying unary_predicate(element) != true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfNotSlice[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {


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



// CopyReplaceIfNotMap returns a copy of original where each element satisfying unary_predicate(element) != true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfNotMap[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {


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



// CopyReplaceIfNotString returns a copy of original where each element satisfying unary_predicate(element) != true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfNotString[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {


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



// CopyReplaceIfNotIterator returns a copy of original where each element satisfying unary_predicate(element) != true is replaced with replacement.
//
// Possible Error values:
//    - EmptyIterableError
func CopyReplaceIfNotIterator[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {


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


// 
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
// 
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
// 
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
// 
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
// 


// CopyExceptIfSlice returns a copy of original without all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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



// CopyExceptIfMap returns a copy of original without all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfMap[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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



// CopyExceptIfString returns a copy of original without all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfString[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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



// CopyExceptIfIterator returns a copy of original without all element satisfying unary_predicate(element) == true).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfIterator[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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


// 
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
// 
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
// 
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
// 
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
// 


// CopyExceptIfNotSlice returns a copy of original without all element satisfying unary_predicate(element) == false).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfNotSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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



// CopyExceptIfNotMap returns a copy of original without all element satisfying unary_predicate(element) == false).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfNotMap[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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



// CopyExceptIfNotString returns a copy of original without all element satisfying unary_predicate(element) == false).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfNotString[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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



// CopyExceptIfNotIterator returns a copy of original without all element satisfying unary_predicate(element) == false).
//
// Possible Error values:
//    - EmptyIterableError
func CopyExceptIfNotIterator[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {


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
// 
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
// 
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
// 
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
// 


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



// FillMap fills the array pointed to by arr with filler.
// all indices in the range [0, cap(*arr)[ are filled regardless of what len(*arr) is.
func FillMap[T any](arr *[]T, filler T) []T {


	for i := range *arr {
		(*arr)[i] = filler
	}

	n_unfilled := cap(*arr) - len(*arr)

	for i := 0; i < n_unfilled; i++ {
		*arr = append(*arr, filler)
	}

	return *arr
}



// FillString fills the array pointed to by arr with filler.
// all indices in the range [0, cap(*arr)[ are filled regardless of what len(*arr) is.
func FillString[T any](arr *[]T, filler T) []T {


	for i := range *arr {
		(*arr)[i] = filler
	}

	n_unfilled := cap(*arr) - len(*arr)

	for i := 0; i < n_unfilled; i++ {
		*arr = append(*arr, filler)
	}

	return *arr
}



// FillIterator fills the array pointed to by arr with filler.
// all indices in the range [0, cap(*arr)[ are filled regardless of what len(*arr) is.
func FillIterator[T any](arr *[]T, filler T) []T {


	for i := range *arr {
		(*arr)[i] = filler
	}

	n_unfilled := cap(*arr) - len(*arr)

	for i := 0; i < n_unfilled; i++ {
		*arr = append(*arr, filler)
	}

	return *arr
}


// 
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
// 
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
// 
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
// 
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
// 


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
		if  {
			index =  ExecutionError[int
			err =  T]{BadItemIndex: i, BadItem: value, Inner: err}

			return
		}
	}

	return nil
}



// TransformMap applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
// Errors returned by transformer are propagated to the caller of TransformMap.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformMap[T any](container []T, transformer func(*T) error) error {


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



// TransformString applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
// Errors returned by transformer are propagated to the caller of TransformString.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformString[T any](container []T, transformer func(*T) error) error {


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



// TransformIterator applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
// Errors returned by transformer are propagated to the caller of TransformIterator.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformIterator[T any](container []T, transformer func(*T) error) error {


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


// 
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
// 
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
// 
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
// 
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
// 


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



// TransformMapUnsafe applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
//
// Possible Error values:
//    - EmptyIterableError
func TransformMapUnsafe[T any](container []T, transformer func(*T)) error {


	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i := range container {
		transformer(&container[i])
	}

	return nil
}



// TransformStringUnsafe applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
//
// Possible Error values:
//    - EmptyIterableError
func TransformStringUnsafe[T any](container []T, transformer func(*T)) error {


	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i := range container {
		transformer(&container[i])
	}

	return nil
}



// TransformIteratorUnsafe applies transformer(&container[i]) for all i in [0, len(container)[ and stores them at container[i].
//
// Possible Error values:
//    - EmptyIterableError
func TransformIteratorUnsafe[T any](container []T, transformer func(*T)) error {


	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i := range container {
		transformer(&container[i])
	}

	return nil
}


// 
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
// 
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
// 
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
// 
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
// 


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



// TransformCopyMap applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
// Errors returned by transformer are propagated to the caller of TransformCopyMap.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformCopyMap[T any, TOut any](container []T, transformer func(T) (TOut, error)) ([]TOut, error) {


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



// TransformCopyString applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
// Errors returned by transformer are propagated to the caller of TransformCopyString.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformCopyString[T any, TOut any](container []T, transformer func(T) (TOut, error)) ([]TOut, error) {


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



// TransformCopyIterator applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
// Errors returned by transformer are propagated to the caller of TransformCopyIterator.
//
// Possible Error values:
//    - EmptyIterableError
//    - ExecutionError
func TransformCopyIterator[T any, TOut any](container []T, transformer func(T) (TOut, error)) ([]TOut, error) {


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


// 
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
// 
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
// 
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
// 
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
// 


// TransformCopySliceUnsafe applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
//
// Possible Error values:
//    - EmptyIterableError
func TransformCopySliceUnsafe[T any, TOut any](container []T, transformer func(T) TOut) ([]TOut, error) {


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



// TransformCopyMapUnsafe applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
//
// Possible Error values:
//    - EmptyIterableError
func TransformCopyMapUnsafe[T any, TOut any](container []T, transformer func(T) TOut) ([]TOut, error) {


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



// TransformCopyStringUnsafe applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
//
// Possible Error values:
//    - EmptyIterableError
func TransformCopyStringUnsafe[T any, TOut any](container []T, transformer func(T) TOut) ([]TOut, error) {


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



// TransformCopyIteratorUnsafe applies transformer(container[i]) for all i in [0, len(container)[  and and returns the newly created container.
// Note that the transformer can return a different type than it's input.
//
// Possible Error values:
//    - EmptyIterableError
func TransformCopyIteratorUnsafe[T any, TOut any](container []T, transformer func(T) TOut) ([]TOut, error) {


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




// MinSlicePred finds the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinSlicePred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {


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



// MinMapPred finds the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMapPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {


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



// MinStringPred finds the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinStringPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {


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



// MinIteratorPred finds the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinIteratorPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {


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


// 
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
// 
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
// 
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
// 
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
// 
// 

// MaxSlicePred finds the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxSlicePred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {


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

// 

// MaxMapPred finds the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxMapPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {


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

// 

// MaxStringPred finds the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxStringPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {


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

// 

// MaxIteratorPred finds the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxIteratorPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (T, error) {


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

// 
// 
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
// 
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
// 
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
// 
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
// 


// MinMaxSlicePred finds the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxSlicePred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (T, T, error) {


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



// MinMaxMapPred finds the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxMapPred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (T, T, error) {


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



// MinMaxStringPred finds the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxStringPred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (T, T, error) {


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



// MinMaxIteratorPred finds the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxIteratorPred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (T, T, error) {


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


// 
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
// 
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
// 
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
// 
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
// 


// MinElementSlicePred finds the index of the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinElementSlicePred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {


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



// MinElementMapPred finds the index of the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinElementMapPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {


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



// MinElementStringPred finds the index of the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinElementStringPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {


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



// MinElementIteratorPred finds the index of the smallest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinElementIteratorPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {


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


// 
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
// 
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
// 
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
// 
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
// 


// MaxElementSlicePred finds the index of the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElementSlicePred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {


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



// MaxElementMapPred finds the index of the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElementMapPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {


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



// MaxElementStringPred finds the index of the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElementStringPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {


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



// MaxElementIteratorPred finds the index of the largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MaxElementIteratorPred[T constraints.Ordered](haystack []T, binary_predicate func(T, T) bool) (int, error) {


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


// 
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
// 
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
// 
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
// 
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
// 


// MinMaxElementSlicePred finds the index of the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElementSlicePred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (int, int, error) {


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



// MinMaxElementMapPred finds the index of the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElementMapPred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (int, int, error) {


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



// MinMaxElementStringPred finds the index of the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElementStringPred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (int, int, error) {


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



// MinMaxElementIteratorPred finds the index of the smallest and largest value in haystack.
// The elements are compared with binary_predicate.
//
// Possible Error values:
//    - EmptyIterableError
func MinMaxElementIteratorPred[T constraints.Ordered](haystack []T, binary_predicate_min func(T, T) bool, binary_predicate_max func(T, T) bool) (int, int, error) {


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


// 
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
// 
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
// 
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
// 
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
// 
