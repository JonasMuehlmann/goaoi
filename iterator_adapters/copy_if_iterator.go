package iteratoradapters

import (
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type CopyReplaceIf[TKey any, TValue any] struct {
	compounditerators.ReadForIndexIterator[TKey, TValue]
	unaryPredicate func(value TValue) bool
	index          int
	size           int
	replacement    TValue
}

func NewCopyReplaceIf[TKey any, TValue any](inner compounditerators.ReadForIndexIterator[TKey, TValue], unaryPredicate func(TValue) bool, replacement TValue) compounditerators.ReadForIndexIterator[TKey, TValue] {
	return &CopyReplaceIf[TKey, TValue]{
		ReadForIndexIterator: inner,
		unaryPredicate:       unaryPredicate,
		index:                -1,
		size:                 0,
	}
}

func (it *CopyReplaceIf[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *CopyReplaceIf[TKey, TValue]) IsEnd() bool {
	return it.ReadForIndexIterator.IsEnd()
}

func (it *CopyReplaceIf[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *CopyReplaceIf[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

func (it *CopyReplaceIf[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *CopyReplaceIf[TKey, TValue]) Get() (value TValue, found bool) {
	val, found := it.ReadForIndexIterator.Get()
	if it.unaryPredicate(val) {
		val = it.replacement
	}

	return val, found
}

func (it *CopyReplaceIf[TKey, TValue]) Next() bool {
	if it.ReadForIndexIterator.IsEnd() {
		return false
	}

	found := it.ReadForIndexIterator.Next()
	if found {
		it.index++
		it.size++

		return true
	}

	return false
}

func (it *CopyReplaceIf[TKey, TValue]) NextN(n int) bool {
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

func (it *CopyReplaceIf[TKey, TValue]) Size() int {
	return it.size
}

func (it *CopyReplaceIf[TKey, TValue]) Index() (int, bool) {
	return it.index, it.IsValid()
}
