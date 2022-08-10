package intermediatemap

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/barweiss/go-tuple"
)

type Iterator[TKey comparable, TValue any] struct {
	ds.ReadWriteOrdCompBidRandCollIterator[int, tuple.T2[TKey, TValue]]
}

func (m *Map[TKey, TValue]) NewIterator(index int, size int) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return &Iterator[TKey, TValue]{m.List.Begin()}
}

func (it *Iterator[TKey, TValue]) MoveToKey(key TKey) (found bool) {
	forwardBack := *it
	backwardBack := *it

	for forwardBack.Next() {
		key, _ := forwardBack.GetKey()
		if key == key {
			i, _ := forwardBack.Index()
			it.MoveTo(i)

			return true
		}
	}
	for backwardBack.Previous() {
		key, _ := forwardBack.GetKey()
		if key == key {
			i, _ := forwardBack.Index()
			it.MoveTo(i)

			return true
		}
	}

	return false
}

func (it *Iterator[TKey, TValue]) Get() (value TValue, found bool) {
	pair, found := it.ReadWriteOrdCompBidRandCollIterator.Get()

	return pair.V2, found
}

func (it *Iterator[TKey, TValue]) GetKey() (value TKey, found bool) {
	pair, found := it.ReadWriteOrdCompBidRandCollIterator.Get()

	return pair.V1, found
}
func (it *Iterator[TKey, TValue]) GetAt(i int) (value TValue, found bool) {
	pair, found := it.ReadWriteOrdCompBidRandCollIterator.GetAt(i)

	return pair.V2, found
}

func (it *Iterator[TKey, TValue]) GetAtKey(key TKey) (value TValue, found bool) {
	curKey, _ := it.GetKey()
	if curKey == key {
		return it.Get()
	}

	iBack, _ := it.Index()

	for it.Next() {
		value, _ := it.ReadWriteOrdCompBidRandCollIterator.Get()
		if value.V1 == key {
			i, _ := it.Index()
			value, _ := it.GetAt(i)
			it.MoveTo(iBack)

			return value, true
		}
	}

	return
}

func (it *Iterator[TKey, TValue]) Set(value TValue) (found bool) {
	new, _ := it.ReadWriteOrdCompBidRandCollIterator.Get()
	new.V2 = value

	return it.ReadWriteOrdCompBidRandCollIterator.Set(new)
}

func (it *Iterator[TKey, TValue]) SetAt(i int, value TValue) (found bool) {
	new, _ := it.ReadWriteOrdCompBidRandCollIterator.Get()
	new.V2 = value

	return it.ReadWriteOrdCompBidRandCollIterator.SetAt(i, new)
}

func (it *Iterator[TKey, TValue]) SetAtKey(key TKey, value TValue) bool {
	curKey, _ := it.GetKey()
	if curKey == key {
		return it.Set(value)
	}

	iBack, _ := it.Index()

	for it.Next() {
		value, _ := it.ReadWriteOrdCompBidRandCollIterator.Get()
		if value.V1 == key {
			i, _ := it.Index()

			found := it.SetAt(i, value.V2)
			it.MoveTo(iBack)

			return found
		}
	}

	return false
}
