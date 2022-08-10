package iteratoradapters

import (
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type TakeIf[TKey any, TValue any] struct {
	compounditerators.ReadForIndexIterator[TKey, TValue]
	unaryPredicate func(value TValue) bool
	index          int
	size           int
	done           bool
}

func NewTakeIf[TKey any, TValue any](inner compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) compounditerators.ReadForIndexIterator[TKey, TValue] {
	return &TakeIf[TKey, TValue]{
		ReadForIndexIterator: inner,
		unaryPredicate:       unaryPredicate,
		index:                -1,
		size:                 0,
	}
}

func (it *TakeIf[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *TakeIf[TKey, TValue]) IsEnd() bool {
	return it.done || it.ReadForIndexIterator.IsEnd()
}

func (it *TakeIf[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *TakeIf[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

func (it *TakeIf[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *TakeIf[TKey, TValue]) Get() (value TValue, found bool) {
	return it.ReadForIndexIterator.Get()
}

func (it *TakeIf[TKey, TValue]) Next() bool {
	if it.ReadForIndexIterator.IsEnd() {
		it.done = true

		return false
	}

	it.ReadForIndexIterator.Next()

	value, found := it.Get()

	for found && !it.unaryPredicate(value) {
		it.ReadForIndexIterator.Next()
		value, found = it.Get()
	}

	it.index++
	it.size++

	return found
}

func (it *TakeIf[TKey, TValue]) NextN(n int) bool {
	if !it.IsValid() {
		return false
	}

	for i := 0; i < n; i++ {
		found := it.Next()

		if !found {
			return false
		}
	}

	return true
}

func (it *TakeIf[TKey, TValue]) Size() int {
	return it.size
}

func (it *TakeIf[TKey, TValue]) Index() (int, bool) {
	return it.index, it.IsValid()
}
