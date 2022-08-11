package generators

import (
	"math"

	compounditerators "github.com/JonasMuehlmann/goaoi/compound_iterators"
	"golang.org/x/exp/constraints"
)

const (
	ErrorNegativeRange = "range must be increasing"
	ErrorUncleanStep   = "step would overshoot stop"
	ErrorTooManyArgs   = "the number of args must be in the range of [0,3]"
)

type Range[TValue constraints.Float | constraints.Integer] struct {
	i       TValue
	start   TValue
	stop    TValue
	step    TValue
	index   int
	isBegin bool
}

func NewRange[TValue constraints.Float | constraints.Integer](args ...TValue) compounditerators.ReadForIndexIterator[int, TValue] {
	it := &Range[TValue]{isBegin: true}

	// [0, stop]
	if len(args) == 1 {
		it.step = 1
		it.stop = args[0] + it.step
		// [start, stop]
	} else if len(args) == 2 {
		it.start = args[0]
		it.step = 1
		it.stop = args[1] + it.step
		// [start, stop]
	} else if len(args) == 3 {
		it.start = args[0]
		it.step = args[2]
		it.stop = args[1] + it.step

		_, mod := math.Modf(float64(it.stop-it.start-it.step) / float64(it.step))
		epsilon := math.Nextafter(1.0, 2.0) - 1

		if math.Abs(mod) > epsilon {
			panic(ErrorUncleanStep)
		}
	} else if len(args) > 3 {
		panic(ErrorTooManyArgs)
	}

	// We don't support reverse counting atm
	if it.step < 0 || it.stop < it.start {
		panic(ErrorNegativeRange)
	}

	return it
}

func (it *Range[TValue]) IsBegin() bool {
	return it.isBegin
}

func (it *Range[TValue]) IsEnd() bool {
	return it.i == it.stop && it.i > it.start || it.start == it.stop
}

func (it *Range[TValue]) IsFirst() bool {
	return it.i == it.start
}

func (it *Range[TValue]) IsLast() bool {
	return it.i == it.stop-it.step
}

func (it *Range[TValue]) IsValid() bool {
	return !it.IsBegin() && !it.IsEnd()
}

func (it *Range[TValue]) Get() (value TValue, found bool) {
	return it.i, it.IsValid()
}

func (it *Range[TValue]) GetKey() (value int, found bool) {
	return 0, false
}

func (it *Range[TValue]) Next() bool {
	if it.isBegin {
		it.i = it.start
		it.isBegin = false
	} else {
		it.i += it.step
	}
	it.index++

	return it.IsValid()
}

func (it *Range[TValue]) NextN(n int) bool {
	return false
}

func (it *Range[TValue]) Size() int {
	if it.start == it.stop {
		return 0
	}

	return int(float64((it.stop - it.start)) / float64(it.step))
}

func (it *Range[TValue]) Index() (int, bool) {
	return it.index, it.IsValid()
}
