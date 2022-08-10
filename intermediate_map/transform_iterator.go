package intermediatemap

// import (
// 	"github.com/JonasMuehlmann/datastructures.go/ds"
// )

// type TransformIterator[TKey comparable, TValue any] struct {
// 	*Iterator[TKey, TValue]

// 	// GetValueTransformer func(inner TValueInner) TValue
// 	// GetKeyTransformer   func(inner TKeyInner) TKey

// 	// SetValueTransformer func(inner TValue) TValueInner
// 	// SetKeyTransformer   func(inner TKey) TKeyInner
// }

// func NewTransformIterator[TKey comparable, TValue any](inner *Iterator[TKey, TValue]) ds.ReadWriteOrdCompBidRandCollMapIterator[TKey, TValue] {

// 	return &TransformIterator[TKey, TValue]{inner}
// }

// func (it *TransformIterator[TKey, TValue]) MoveToKey(key TKey) (found bool) {
// 	forwardBack := *it
// 	backwardBack := *it

// 	for forwardBack.Next() {
// 		key, _ := forwardBack.GetKey()
// 		if key == key {
// 			i, _ := forwardBack.Index()
// 			it.MoveTo(i)

// 			return true
// 		}
// 	}
// 	for backwardBack.Previous() {
// 		key, _ := forwardBack.GetKey()
// 		if key == key {
// 			i, _ := forwardBack.Index()
// 			it.MoveTo(i)

// 			return true
// 		}
// 	}

// 	return false
// }

// func (it *TransformIterator[TKey, TValue]) Get() (value TValue, found bool) {
// 	pair, found := it.Iterator.Get()

// 	return pair.V2, found
// }

// func (it *TransformIterator[TKey, TValue]) GetKey() (value TKey, found bool) {
// 	pair, found := it.Iterator.Get()

// 	return pair.V1, found
// }

// func (it *TransformIterator[TKey, TValue]) GetAtKey(key TKey) (value TValue, found bool) {
// 	itBack := *it.Iterator

// 	for itBack.Next() {
// 		value, _ := itBack.Get()
// 		if value.V1 == key {
// 			i, _ := itBack.Index()
// 			pair, _ := it.GetAt(i)

// 			return pair.V2, true
// 		}
// 	}

// 	return
// }

// func (it *TransformIterator[TKey, TValue]) Set(value TValue) (found bool) {
// 	new, _ := it.Iterator.Get()
// 	new.V2 = value

// 	return it.Iterator.Set(new)
// }

// func (it *TransformIterator[TKey, TValue]) SetAtKey(key TKey, value TValue) bool {
// 	itBack := *it.Iterator

// 	for itBack.Next() {
// 		value, _ := itBack.Get()
// 		if value.V1 == key {
// 			i, _ := itBack.Index()

// 			return it.SetAt(i, value)
// 		}
// 	}

// 	return false
// }
