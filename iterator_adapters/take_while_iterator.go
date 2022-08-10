package iteratoradapters

import (
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type TakeWhile[TKey any, TValue any] struct {
	compounditerators.ReadForIndexIterator[TKey, TValue]
	unaryPredicate func(value TValue) bool
	index          int
	size           int
}

func NewTakeWhile[TKey any, TValue any](inner compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool) compounditerators.ReadForIndexIterator[TKey, TValue] {
	return &TakeWhile[TKey, TValue]{
		ReadForIndexIterator: inner,
		unaryPredicate:       unaryPredicate,
		index:                -1,
		// Will be set later
		size: -1,
	}
}

func (it *TakeWhile[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *TakeWhile[TKey, TValue]) IsEnd() bool {
	return it.size == 0 || it.index == it.size
}

func (it *TakeWhile[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *TakeWhile[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

func (it *TakeWhile[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *TakeWhile[TKey, TValue]) Get() (value TValue, found bool) {
	return it.ReadForIndexIterator.Get()
}

func (it *TakeWhile[TKey, TValue]) Next() bool {
	found := it.ReadForIndexIterator.Next()

	value, _ := it.Get()

	if !it.unaryPredicate(value) {
		it.size = it.index

		return false
	}

	return found
}

func (it *TakeWhile[TKey, TValue]) NextN(n int) bool {
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

func (it *TakeWhile[TKey, TValue]) Size() int {
	return it.size
}

func (it *TakeWhile[TKey, TValue]) Index() (int, bool) {
	return it.index, it.IsValid()
}
