package iteratoradapters

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
)

type Join[TKey any, TValue any] struct {
	originals     []ds.ReadForIndexIterator[TKey, TValue]
	index         int
	size          int
	sizeOriginals int
	iOriginals    int
}

func NewJoin[TKey any, TValue any](originals ...ds.ReadForIndexIterator[TKey, TValue]) ds.ReadForIndexIterator[TKey, TValue] {
	return &Join[TKey, TValue]{
		originals:     originals,
		index:         -1,
		size:          -1,
		sizeOriginals: len(originals),
	}
}

func (it *Join[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *Join[TKey, TValue]) IsEnd() bool {
	return it.iOriginals == it.sizeOriginals
}

func (it *Join[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *Join[TKey, TValue]) IsLast() bool {
	return false
}

func (it *Join[TKey, TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *Join[TKey, TValue]) Get() (value TValue, found bool) {
	if it.IsValid() {
		value, found = it.originals[it.iOriginals].Get()

		return
	}

	return
}

func (it *Join[TKey, TValue]) Next() bool {
	if it.sizeOriginals == 0 || it.IsEnd() {
		return false
	}

	it.originals[it.iOriginals].Next()

	for !it.IsEnd() && it.originals[it.iOriginals].IsEnd() {
		it.iOriginals++
		if !it.IsEnd() {
			it.originals[it.iOriginals].Next()
		} else {
			return false
		}
	}

	it.index++

	return it.IsValid()
}

func (it *Join[TKey, TValue]) NextN(n int) bool {
	for i := 0; i < n; i++ {
		found := it.Next()

		if !found {
			return false
		}
	}

	return true
}

func (it *Join[TKey, TValue]) Size() int {
	return it.size
}

func (it *Join[TKey, TValue]) Index() (int, bool) {
	return it.index, it.IsValid()
}

func (it *Join[TKey, TValue]) GetKey() (TKey, bool) {
	return it.originals[0].GetKey()
}
