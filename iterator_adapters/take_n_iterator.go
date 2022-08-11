package iteratoradapters

import (
	"github.com/JonasMuehlmann/datastructures.go/utils"
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type TakeN[TKey any, TValue any] struct {
	compounditerators.ReadForIndexIterator[TKey, TValue]
	index int
	size  int
}

func NewTakeN[TKey any, TValue any](inner compounditerators.ReadForIndexIterator[TKey, TValue], n int) compounditerators.ReadForIndexIterator[TKey, TValue] {
	return &TakeN[TKey, TValue]{
		ReadForIndexIterator: inner,
		index:                -1,
		size:                 n,
	}
}

func (it *TakeN[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *TakeN[TKey, TValue]) IsEnd() bool {
	return it.size == 0 || it.index == it.size
}

func (it *TakeN[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *TakeN[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

func (it *TakeN[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *TakeN[TKey, TValue]) Get() (value TValue, found bool) {
	return it.ReadForIndexIterator.Get()
}

func (it *TakeN[TKey, TValue]) Next() bool {
	it.index = utils.Min(it.index+1, it.size)
	if !it.IsValid() {
		return false
	}

	found := it.ReadForIndexIterator.Next()
	if !found {
		it.size = it.index

		return false
	}

	return true
}

func (it *TakeN[TKey, TValue]) NextN(n int) bool {
	for i := 0; i < n; i++ {
		found := it.Next()

		if !found {
			return false
		}
	}

	return true
}

func (it *TakeN[TKey, TValue]) Size() int {
	return it.size
}

func (it *TakeN[TKey, TValue]) Index() (int, bool) {
	return it.index, it.IsValid()
}
