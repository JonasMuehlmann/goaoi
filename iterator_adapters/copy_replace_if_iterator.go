package iteratoradapters

import (
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type CopyIf[TKey any, TValue any] struct {
	compounditerators.ReadForIndexIterator[TKey, TValue]
	unaryPredicate func(value TValue) bool
	index          int
	size           int
	done           bool
}

func NewCopyIf[TKey any, TValue any](inner compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) compounditerators.ReadForIndexIterator[TKey, TValue] {
	return &CopyIf[TKey, TValue]{
		ReadForIndexIterator: inner,
		unaryPredicate:       unaryPredicate,
		index:                -1,
		size:                 0,
	}
}

func (it *CopyIf[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *CopyIf[TKey, TValue]) IsEnd() bool {
	return it.done || it.ReadForIndexIterator.IsEnd()
}

func (it *CopyIf[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *CopyIf[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

func (it *CopyIf[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *CopyIf[TKey, TValue]) Get() (value TValue, found bool) {
	return it.ReadForIndexIterator.Get()
}

func (it *CopyIf[TKey, TValue]) Next() bool {
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

func (it *CopyIf[TKey, TValue]) NextN(n int) bool {
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

func (it *CopyIf[TKey, TValue]) Size() int {
	return it.size
}

func (it *CopyIf[TKey, TValue]) Index() (int, bool) {
	return it.index, it.IsValid()
}
