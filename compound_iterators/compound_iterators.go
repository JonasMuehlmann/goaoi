package compounditerators

import "github.com/JonasMuehlmann/datastructures.go/ds"

type ReadForIndexIterator[TKey any, TValue any] interface {
	ds.ReadForIterator[TValue]
	ds.IndexedIterator[TKey]
}
