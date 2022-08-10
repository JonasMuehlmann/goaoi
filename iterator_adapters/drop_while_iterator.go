package iteratoradapters

import (
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type DropWhile[TKey any, TValue any] struct {
	compounditerators.ReadForIndexIterator[TKey, TValue]
	unaryPredicate func(value TValue) bool
	index          int
	size           int
}

func NewDropWhile[TKey any, TValue any](inner compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) compounditerators.ReadForIndexIterator[TKey, TValue] {
	return &DropWhile[TKey, TValue]{
		ReadForIndexIterator: inner,
		unaryPredicate:       unaryPredicate,
		index:                -1,
		size:                 0,
	}
}

func (it *DropWhile[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *DropWhile[TKey, TValue]) IsEnd() bool {
	return it.size == -1
}

func (it *DropWhile[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *DropWhile[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

func (it *DropWhile[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *DropWhile[TKey, TValue]) Get() (value TValue, found bool) {
	return it.ReadForIndexIterator.Get()
}

func (it *DropWhile[TKey, TValue]) Next() bool {
	// Dropping
	for it.IsBegin() && !it.ReadForIndexIterator.IsEnd() {
		it.ReadForIndexIterator.Next()
		value, found := it.Get()

		if !it.unaryPredicate(value) {
			it.index++

			return found
		}
	}

	// Taking
	it.ReadForIndexIterator.Next()

	if it.ReadForIndexIterator.IsEnd() {
		it.size = -1

		return false
	}

	it.index++

	return true
}

func (it *DropWhile[TKey, TValue]) NextN(n int) bool {
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

func (it *DropWhile[TKey, TValue]) Size() int {
	return -1
}

func (it *DropWhile[TKey, TValue]) Index() (int, bool) {
	return it.index, it.IsValid()
}
