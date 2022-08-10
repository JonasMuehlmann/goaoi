package intermediatemap

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
	"github.com/barweiss/go-tuple"
)

type Map[TKey comparable, TValue any] struct {
	*arraylist.List[tuple.T2[TKey, TValue]]
}

func (m *Map[TKey, TValue]) Begin() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewIterator(-1, m.Size())
}
