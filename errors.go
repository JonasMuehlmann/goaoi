package goaoi

import "fmt"

type ComparisonError[T comparable] struct {
	BadItem T
}

func (error ComparisonError[T]) Error() string {
	return fmt.Sprintf("Item at index %v did not satisfy comparison", error.BadItem)
}

type EmptyIterableError struct {
}

func (error EmptyIterableError) Error() string {
	return "Iterable is empty"
}

type ExecutionError[T any] struct {
	BadItem T
	Inner   error
}

func (error ExecutionError[T]) Error() string {
	return fmt.Sprintf("Item at index %v returned error after application of function: %v", error.BadItem, error.Inner)
}

type EqualIteratorsError struct {
}

func (error EqualIteratorsError) Error() string {
	return "Iterables are equal"
}

type ElementNotFoundError struct {
}

func (error ElementNotFoundError) Error() string {
	return "Could not find element"
}
