package goaoi

import (
	"fmt"
	"reflect"
)

//******************************************************************//
//                          ComparisonError                         //
//******************************************************************//

type ComparisonError[TIndex comparable, TItem any] struct {
	BadItemIndex TIndex
	BadItem      TItem
}

func (error ComparisonError[TIndex, TItem]) Error() string {
	return fmt.Sprintf("Item at index %v did not satisfy comparison", error.BadItemIndex)
}

func (err ComparisonError[TIndex, TItem]) Is(other error) bool {
	switch other.(type) {
	case ComparisonError[TIndex, TItem]:
		return true
	default:
		return false
	}
}

func (err ComparisonError[TIndex, TItem]) As(target any) bool {
	switch target.(type) {
	case ComparisonError[TIndex, TItem]:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
		return true
	default:
		return false
	}
}

//******************************************************************//
//                        EmptyIterableError                        //
//******************************************************************//

type EmptyIterableError struct{}

func (error EmptyIterableError) Error() string {
	return "Iterable is empty"
}
func (err EmptyIterableError) Is(other error) bool {
	switch other.(type) {
	case EmptyIterableError:
		return true
	default:
		return false
	}
}

func (err EmptyIterableError) As(target any) bool {
	switch target.(type) {
	case EmptyIterableError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
		return true
	default:
		return false
	}
}

//******************************************************************//
//                          ExecutionError                          //
//******************************************************************//

type ExecutionError[TIndex, TItem any] struct {
	BadItemIndex TIndex
	BadItem      TItem
	Inner        error
}

func (error ExecutionError[TIndex, TItem]) Error() string {
	return fmt.Sprintf("Item at index %v returned error after application of function: %v", error.BadItemIndex, error.Inner)
}

func (err ExecutionError[TIndex, TItem]) Unwrap() error {
	return err.Inner
}

func (err ExecutionError[TIndex, TItem]) Is(other error) bool {
	switch other.(type) {
	case ExecutionError[TIndex, TItem]:
		return true
	default:
		return false
	}
}

func (err ExecutionError[TIndex, TItem]) As(target any) bool {
	switch target.(type) {
	case ExecutionError[TIndex, TItem]:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
		return true
	default:
		return false
	}
}

//******************************************************************//
//                        EqualIteratorsError                       //
//******************************************************************//

type EqualIteratorsError struct{}

func (error EqualIteratorsError) Error() string {
	return "Iterables are equal"
}

func (err EqualIteratorsError) Is(other error) bool {
	switch other.(type) {
	case EqualIteratorsError:
		return true
	default:
		return false
	}
}

func (err EqualIteratorsError) As(target any) bool {
	switch target.(type) {
	case EqualIteratorsError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
		return true
	default:
		return false
	}
}

//******************************************************************//
//                       ElementNotFoundError                       //
//******************************************************************//

type ElementNotFoundError struct{}

func (error ElementNotFoundError) Error() string {
	return "Could not find element"
}

func (err ElementNotFoundError) Is(other error) bool {
	switch other.(type) {
	case ElementNotFoundError:
		return true
	default:
		return false
	}
}

func (err ElementNotFoundError) As(target any) bool {
	switch target.(type) {
	case ElementNotFoundError:
		reflect.Indirect(reflect.ValueOf(target)).Set(reflect.ValueOf(err))
		return true
	default:
		return false
	}
}
