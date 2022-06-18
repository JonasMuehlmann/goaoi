package goaoi

import "fmt"

type ComparisonError[TIndex comparable, TItem any] struct {
	BadItemIndex TIndex
	BadItem      TItem
}

func (error ComparisonError[TIndex, TItem]) Error() string {
	return fmt.Sprintf("Item at index %v did not satisfy comparison", error.BadItemIndex)
}

type EmptyIterableError struct{}

func (error EmptyIterableError) Error() string {
	return "Iterable is empty"
}

type ExecutionError[TIndex, TItem any] struct {
	BadItemIndex TIndex
	BadItem      TItem
	Inner        error
}

func (error ExecutionError[TIndex, TItem]) Error() string {
	return fmt.Sprintf("Item at index %v returned error after application of function: %v", error.BadItemIndex, error.Inner)
}

type EqualIteratorsError struct{}

func (error EqualIteratorsError) Error() string {
	return "Iterables are equal"
}

type ElementNotFoundError struct{}

func (error ElementNotFoundError) Error() string {
	return "Could not find element"
}
