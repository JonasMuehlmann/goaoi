package generators

import (
	"github.com/JonasMuehlmann/datastructures.go/utils"
	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
)

type Repeat[TValue any] struct {
	index     int
	limit     int
	generator func() (TValue, bool)
}

func NewRepeat[TValue any](generator func() (TValue, bool), limit int) compounditerators.ReadForIndexIterator[int, TValue] {
	it := &Repeat[TValue]{
		index:     -1,
		limit:     limit,
		generator: generator,
	}

	return it
}

func (it *Repeat[TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *Repeat[TValue]) IsEnd() bool {
	return it.index == it.limit
}

func (it *Repeat[TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *Repeat[TValue]) IsLast() bool {
	return it.index == it.limit-1
}

func (it *Repeat[TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *Repeat[TValue]) Get() (value TValue, found bool) {
	return it.generator()
}

func (it *Repeat[TValue]) Next() bool {
	if it.limit < 0 {
		it.index++
	} else {
		it.index = utils.Min(it.index+1, it.limit)
	}

	return it.IsValid()
}

func (it *Repeat[TValue]) NextN(n int) bool {
	if it.limit < 0 {
		it.index += n
	} else {
		it.index = utils.Min(it.index+n, it.limit)
	}

	return it.IsValid()
}

func (it *Repeat[TValue]) Size() int {
	return it.limit
}

func (it *Repeat[TValue]) Index() (int, bool) {
	return it.index, true
}
func (it *Repeat[TValue]) GetKey() (int, bool) {
	return it.Index()
}
