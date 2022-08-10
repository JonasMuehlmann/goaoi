package iteratoradapters

import (
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type Transform[TKey any, TValue any] struct {
	compounditerators.ReadForIndexIterator[TKey, TValue]
	transformer func(value TValue) (TValue, error)
	index       int
	size        int
	done        bool
}

func NewTransformIterator[TKey any, TValue any](inner compounditerators.ReadForIndexIterator[TKey, TValue], transformer func(TValue) (TValue, error)) compounditerators.ReadForIndexIterator[TKey, TValue] {
	return &Transform[TKey, TValue]{
		ReadForIndexIterator: inner,
		transformer:          transformer,
		index:                -1,
		size:                 0,
	}
}

func (it *Transform[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *Transform[TKey, TValue]) IsEnd() bool {
	return it.done || it.ReadForIndexIterator.IsEnd()
}

func (it *Transform[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *Transform[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

func (it *Transform[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *Transform[TKey, TValue]) Get() (value TValue, found bool) {
	var err error
	val, found := it.ReadForIndexIterator.Get()

	val, err = it.transformer(val)
	if err != nil {
		it.done = true
	}

	return val, found
}

func (it *Transform[TKey, TValue]) Next() bool {
	if it.ReadForIndexIterator.IsEnd() {
		return false
	}

	found := it.ReadForIndexIterator.Next()

	if !found {
		return false
	}

	it.index++
	it.size++

	return true
}

func (it *Transform[TKey, TValue]) NextN(n int) bool {
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

func (it *Transform[TKey, TValue]) Size() int {
	return it.size
}

func (it *Transform[TKey, TValue]) Index() (int, bool) {
	return it.ReadForIndexIterator.Index()
}
