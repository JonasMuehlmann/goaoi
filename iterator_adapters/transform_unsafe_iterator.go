package iteratoradapters

import (
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type TransformUnsafe[TKey any, TValue any] struct {
	compounditerators.ReadForIndexIterator[TKey, TValue]
	transformer func(value TValue) TValue
	index       int
	size        int
}

func NewTransformUnsafeIterator[TKey any, TValue any](inner compounditerators.ReadForIndexIterator[TKey, TValue], transformer func(TValue) TValue) compounditerators.ReadForIndexIterator[TKey, TValue] {
	return &TransformUnsafe[TKey, TValue]{
		ReadForIndexIterator: inner,
		transformer:          transformer,
		index:                -1,
		size:                 0,
	}
}

func (it *TransformUnsafe[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *TransformUnsafe[TKey, TValue]) IsEnd() bool {
	return it.ReadForIndexIterator.IsEnd()
}

func (it *TransformUnsafe[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *TransformUnsafe[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

func (it *TransformUnsafe[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *TransformUnsafe[TKey, TValue]) Get() (value TValue, found bool) {
	val, found := it.ReadForIndexIterator.Get()

	val = it.transformer(val)

	return val, found
}

func (it *TransformUnsafe[TKey, TValue]) Next() bool {
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

func (it *TransformUnsafe[TKey, TValue]) NextN(n int) bool {
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

func (it *TransformUnsafe[TKey, TValue]) Size() int {
	return it.size
}

func (it *TransformUnsafe[TKey, TValue]) Index() (int, bool) {
	return it.ReadForIndexIterator.Index()
}
