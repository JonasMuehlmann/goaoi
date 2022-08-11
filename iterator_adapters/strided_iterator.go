package iteratoradapters

import (
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type Strided[TKey any, TValue any] struct {
	compounditerators.ReadForIndexIterator[TKey, TValue]
	stride int
	index  int
	size   int
}

// TODO: Split the Size() method into separate interface
func NewStrided[TKey any, TValue any](inner compounditerators.ReadForIndexIterator[TKey, TValue], n int) compounditerators.ReadForIndexIterator[TKey, TValue] {
	it := &Strided[TKey, TValue]{
		ReadForIndexIterator: inner,
		stride:               n,
		index:                -1,
		size:                 0,
	}

	if n == 0 {
		it.size = -1
	}

	return it
}

func (it *Strided[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *Strided[TKey, TValue]) IsEnd() bool {
	return it.size == -1
}

func (it *Strided[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *Strided[TKey, TValue]) IsLast() bool {
	return false
}

func (it *Strided[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *Strided[TKey, TValue]) Get() (value TValue, found bool) {
	return it.ReadForIndexIterator.Get()
}

func (it *Strided[TKey, TValue]) Next() bool {
	if it.IsEnd() {
		return false
	}

	if it.IsBegin() {
		it.index++
		it.ReadForIndexIterator.Next()

		return it.IsValid()
	}

	valid := it.ReadForIndexIterator.NextN(it.stride)

	if !valid {
		it.size = -1

		return false
	}

	it.index++

	return true
}

func (it *Strided[TKey, TValue]) NextN(n int) bool {
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

func (it *Strided[TKey, TValue]) Size() int {
	return -1
}

func (it *Strided[TKey, TValue]) Index() (int, bool) {
	return it.index, it.IsValid()
}
