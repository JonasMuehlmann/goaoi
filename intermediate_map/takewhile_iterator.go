package intermediatemap

// import "github.com/JonasMuehlmann/datastructures.go/ds"

// type TakeWhileIterator[TKey comparable, TValue any, TInner ds.ReadableIterator[TValue]] struct {
// 	inner              TInner
// 	unaryPredicate     func(value TValue) bool
// 	orderedIterator    ds.OrderedIterator
// 	comparableIterator ds.ComparableIterator
// }

// type InvalidComparableIterator struct {
// 	ds.Iterator
// }

// func (it *InvalidComparableIterator) IsEqual(other ds.ComparableIterator) bool {
// 	panic("Inner iterator is not comparable")
// }

// type InvalidOrderedIterator struct {
// 	ds.Iterator
// }

// func (it *InvalidOrderedIterator) DistanceTo(other ds.OrderedIterator) int {
// 	panic("Inner iterator is not ordered")
// }

// func (it *InvalidOrderedIterator) IsAfter(other ds.OrderedIterator) bool {
// 	panic("Inner iterator is not ordered")
// }

// func (it *InvalidOrderedIterator) IsBefore(other ds.OrderedIterator) bool {
// 	panic("Inner iterator is not ordered")
// }

// type InvalidSizedIterator struct{}

// func (it *InvalidComparableIterator) Size() int {
// 	panic("Inner iterator is not sized")
// }

// func NewTakeWhileIterator[TKey comparable, TValue any, TInner ds.ReadableIterator[TValue]](inner TInner, unaryPredicate func(TValue) bool) ds.ReadWriteOrdCompBidRandCollMapIterator[TKey, TValue] {
// 	it := &TakeWhileIterator[TKey, TValue, TInner]{
// 		inner:          inner,
// 		unaryPredicate: unaryPredicate,
// 	}

// 	orderedIterator, ok := any(inner).(ds.OrderedIterator)
// 	if !ok {
// 		it.orderedIterator = &InvalidOrderedIterator{}
// 	} else {
// 		it.orderedIterator = orderedIterator
// 	}

// 	comparableIterator, ok := any(inner).(ds.ComparableIterator)
// 	if !ok {
// 		it.comparableIterator = &InvalidComparableIterator{}
// 	} else {
// 		it.comparableIterator = comparableIterator
// 	}

// 	return it
// }

// // IsBegin implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) IsBegin() bool {
// 	return it.IsBegin()
// }

// // IsEnd implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) IsEnd() bool {
// 	return it.IsEnd()
// }

// // IsFirst implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) IsFirst() bool {
// 	return it.IsFirst()
// }

// // IsLast implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) IsLast() bool {
// 	return it.IsLast()
// }

// // IsValid implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) IsValid() bool {
// 	return it.IsValid()
// }

// // IsEqual implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) IsEqual(other ds.ComparableIterator) bool {
// 	return it.comparableIterator.IsEqual(other)
// }

// // DistanceTo implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) DistanceTo(other ds.OrderedIterator) int {
// 	return it.orderedIterator.DistanceTo(other)
// }

// // IsAfter implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) IsAfter(other ds.OrderedIterator) bool {
// 	return it.orderedIterator.IsAfter(other)
// }

// // IsBefore implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) IsBefore(other ds.OrderedIterator) bool {
// 	return it.orderedIterator.IsAfter(other)
// }

// // Size implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) Size() int {
// 	return it.sizedIterator.Size()
// }

// // Index implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) Index() (int, bool) {
// 	return it.indexedIterator.Size()
// }

// // GetKey implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) GetKey() (TKey, bool) {
// 	panic("unimplemented")
// }

// // MoveToKey implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) MoveToKey(i TKey) bool {
// 	panic("unimplemented")
// }

// // Get implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) Get() (value TValue, found bool) {
// 	panic("unimplemented")
// }

// // GetAtKey implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) GetAtKey(i TKey) (value TValue, found bool) {
// 	panic("unimplemented")
// }

// // Set implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) Set(value TValue) bool {
// 	panic("unimplemented")
// }

// // SetAtKey implements ds.ReadWriteOrdCompBidRandCollMapIterator
// func (it *TakeWhileIterator[TKey, TValue, TInner]) SetAtKey(i TKey, value TValue) bool {
// 	panic("unimplemented")
// }

// func (it *TakeWhileIterator[TKey, TValue, TInner]) Next() bool {
// 	found := it.ReadWriteOrdCompBidRandCollMapIterator.Next()

// 	value, _ := it.Get()

// 	if !it.unaryPredicate(value) {
// 		end := it.Size()
// 		it.MoveTo(end)

// 		return false
// 	}

// 	return found
// }

// func (it *TakeWhileIterator[TKey, TValue, TInner]) NextN(n int) bool {
// 	if !it.IsValid() {
// 		return false
// 	}

// 	for i := 0; i < n; i++ {
// 		found := it.Next()

// 		if !found {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (it *TakeWhileIterator[TKey, TValue, TInner]) Previous() bool {
// 	found := it.ReadWriteOrdCompBidRandCollMapIterator.Previous()

// 	value, _ := it.Get()

// 	if !it.unaryPredicate(value) {
// 		it.MoveTo(-1)

// 		return false
// 	}

// 	return found
// }

// func (it *TakeWhileIterator[TKey, TValue, TInner]) PreviousN(n int) bool {
// 	if !it.IsValid() {
// 		return false
// 	}

// 	for i := 0; i < n; i++ {
// 		found := it.Previous()

// 		if !found {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (it *TakeWhileIterator[TKey, TValue, TInner]) MoveBy(n int) bool {
// 	if n > 0 {
// 		return it.NextN(n)
// 	} else if n < 0 {
// 		return it.PreviousN(-n)
// 	}

// 	return it.IsValid()
// }

// func (it *TakeWhileIterator[TKey, TValue, TInner]) MoveTo(i int) bool {
// 	if i > 0 {
// 		return it.NextN(i)
// 	} else if i < 0 {
// 		return it.PreviousN(-i)
// 	}

// 	return it.IsValid()
// }

// func (it *TakeWhileIterator[TKey, TValue, TInner]) GetAt(i int) (value TValue, found bool) {
// 	iBack, _ := it.Index()
// 	if i > 0 {
// 		it.NextN(i)

// 		value, found := it.Get()
// 		it.MoveTo(iBack)

// 		return value, found
// 	} else if i < 0 {
// 		it.PreviousN(-i)
// 		value, found := it.Get()

// 		it.MoveTo(iBack)

// 		return value, found
// 	}

// 	return it.Get()
// }

// func (it *TakeWhileIterator[TKey, TValue, TInner]) SetAt(i int, value TValue) (found bool) {
// 	iBack, _ := it.Index()
// 	if i > 0 {
// 		it.NextN(i)

// 		found := it.Set(value)
// 		it.MoveTo(iBack)

// 		return found
// 	} else if i < 0 {
// 		it.PreviousN(-i)
// 		found := it.Set(value)

// 		it.MoveTo(iBack)

// 		return found
// 	}

// 	return it.Set(value)
// }

// type TakeWhileIterator[TKey comparable, TValue any] struct {
// 	ds.ReadWriteOrdCompBidRandCollMapIterator[TKey, TValue]
// 	unaryPredicate func(value TValue) bool
// }

// func NewTakeWhileIterator[TKey comparable, TValue any](inner ds.ReadWriteOrdCompBidRandCollMapIterator[TKey, TValue], unaryPredicate func(TValue) bool) ds.ReadWriteOrdCompBidRandCollMapIterator[TKey, TValue] {
// 	it := &TakeWhileIterator[TKey, TValue]{
// 		ReadWriteOrdCompBidRandCollMapIterator: inner,
// 		unaryPredicate:                         unaryPredicate,
// 	}

// 	return it
// }

// func (it *TakeWhileIterator[TKey, TValue]) Next() bool {
// 	found := it.ReadWriteOrdCompBidRandCollMapIterator.Next()

// 	value, _ := it.Get()

// 	if !it.unaryPredicate(value) {
// 		end := it.Size()
// 		it.MoveTo(end)

// 		return false
// 	}

// 	return found
// }

// func (it *TakeWhileIterator[TKey, TValue]) NextN(n int) bool {
// 	if !it.IsValid() {
// 		return false
// 	}

// 	for i := 0; i < n; i++ {
// 		found := it.Next()

// 		if !found {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (it *TakeWhileIterator[TKey, TValue]) Previous() bool {
// 	found := it.ReadWriteOrdCompBidRandCollMapIterator.Previous()

// 	value, _ := it.Get()

// 	if !it.unaryPredicate(value) {
// 		it.MoveTo(-1)

// 		return false
// 	}

// 	return found
// }

// func (it *TakeWhileIterator[TKey, TValue]) PreviousN(n int) bool {
// 	if !it.IsValid() {
// 		return false
// 	}

// 	for i := 0; i < n; i++ {
// 		found := it.Previous()

// 		if !found {
// 			return false
// 		}
// 	}

// 	return true
// }

// func (it *TakeWhileIterator[TKey, TValue]) MoveBy(n int) bool {
// 	if n > 0 {
// 		return it.NextN(n)
// 	} else if n < 0 {
// 		return it.PreviousN(-n)
// 	}

// 	return it.IsValid()
// }

// func (it *TakeWhileIterator[TKey, TValue]) MoveTo(i int) bool {
// 	if i > 0 {
// 		return it.NextN(i)
// 	} else if i < 0 {
// 		return it.PreviousN(-i)
// 	}

// 	return it.IsValid()
// }

// func (it *TakeWhileIterator[TKey, TValue]) GetAt(i int) (value TValue, found bool) {
// 	iBack, _ := it.Index()
// 	if i > 0 {
// 		it.NextN(i)

// 		value, found := it.Get()
// 		it.MoveTo(iBack)

// 		return value, found
// 	} else if i < 0 {
// 		it.PreviousN(-i)
// 		value, found := it.Get()

// 		it.MoveTo(iBack)

// 		return value, found
// 	}

// 	return it.Get()
// }

// func (it *TakeWhileIterator[TKey, TValue]) SetAt(i int, value TValue) (found bool) {
// 	iBack, _ := it.Index()
// 	if i > 0 {
// 		it.NextN(i)

// 		found := it.Set(value)
// 		it.MoveTo(iBack)

// 		return found
// 	} else if i < 0 {
// 		it.PreviousN(-i)
// 		found := it.Set(value)

// 		it.MoveTo(iBack)

// 		return found
// 	}

// 	return it.Set(value)
// }
