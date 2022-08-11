package iteratoradapters

import (
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type DropN[TKey any, TValue any] struct {
	compounditerators.ReadForIndexIterator[TKey, TValue]
	nToDrop int
	index   int
	size    int
}

func NewDropN[TKey any, TValue any](inner compounditerators.ReadForIndexIterator[TKey, TValue], n int) compounditerators.ReadForIndexIterator[TKey, TValue] {
	return &DropN[TKey, TValue]{
		ReadForIndexIterator: inner,
		nToDrop:              n,
		index:                -1,
		size:                 0,
	}
}

func (it *DropN[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *DropN[TKey, TValue]) IsEnd() bool {
	return it.size == -1
}

func (it *DropN[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *DropN[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

func (it *DropN[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *DropN[TKey, TValue]) Get() (value TValue, found bool) {
	return it.ReadForIndexIterator.Get()
}

func (it *DropN[TKey, TValue]) Next() bool {
	// Dropping
	for it.IsBegin() && !it.ReadForIndexIterator.IsEnd() {
		found := it.ReadForIndexIterator.Next()

		if it.nToDrop == 0 {
			it.index++

			return found
		}

		it.nToDrop--
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

func (it *DropN[TKey, TValue]) NextN(n int) bool {
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

func (it *DropN[TKey, TValue]) Size() int {
	return -1
}

func (it *DropN[TKey, TValue]) Index() (int, bool) {
	return it.index, it.IsValid()
}
