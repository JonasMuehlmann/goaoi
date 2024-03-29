package goaoi_test

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/goaoi/functional"
	"github.com/stretchr/testify/assert"
)

func Test_FindIfSlice(t *testing.T) {
	tcs := []struct {
		haystack   []int
		comparator func(int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2}, func(i int) bool { return i == 1 }, 0, nil, "Found"},
		{[]int{1, 2}, func(i int) bool { return i == 0 }, 0, goaoi.ElementNotFoundError{}, "Not found"},
		{[]int{}, func(i int) bool { return i == 0 }, 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.FindIfSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_FindIfMap(t *testing.T) {
	tcs := []struct {
		haystack   map[string]int
		comparator func(int) bool
		exp        string
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i == 1 }, "a", nil, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i == 0 }, "", goaoi.ElementNotFoundError{}, "Not found"},
		{map[string]int{}, func(i int) bool { return i == 0 }, "", goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.FindIfMap(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_FindEndSlicePred(t *testing.T) {
	tcs := []struct {
		super      []int
		sub        []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2, 3, 4}, []int{3, 4}, functional.AreEqual[int], 2, nil, "Found at end"},
		{[]int{1, 2, 3, 4}, []int{1, 2}, functional.AreEqual[int], 0, nil, "Found at beginning"},
		{[]int{1, 2, 3, 4}, []int{2, 3}, functional.AreEqual[int], 1, nil, "Found in middle"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, functional.AreEqual[int], 0, nil, "Found equal"},
		{[]int{1, 2, 3}, []int{1, 4}, functional.AreEqual[int], 0, goaoi.ElementNotFoundError{}, "Not found"},
		{[]int{}, []int{1, 2, 3}, functional.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Super empty"},
		{[]int{1, 2, 3}, []int{}, functional.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Sub empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.FindEndSlicePred(tc.super, tc.sub, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_FindFirstOfSlicePred(t *testing.T) {
	tcs := []struct {
		haystack   []int
		needles    []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2, 3, 4}, []int{4, 10}, functional.AreEqual[int], 3, nil, "Found at end"},
		{[]int{1, 2, 3, 4}, []int{1, 2}, functional.AreEqual[int], 0, nil, "Found at beginning"},
		{[]int{1, 2, 3, 4}, []int{3, 4}, functional.AreEqual[int], 2, nil, "Found in middle"},
		{[]int{1}, []int{1}, functional.AreEqual[int], 0, nil, "Found equal"},
		{[]int{1, 2, 3}, []int{4, 5}, functional.AreEqual[int], 0, goaoi.ElementNotFoundError{}, "Not found"},
		{[]int{}, []int{1, 2, 3}, functional.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Super empty"},
		{[]int{1, 2, 3}, []int{}, functional.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Sub empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.FindFirstOfSlicePred(tc.haystack, tc.needles, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_FindFirstOfMapPred(t *testing.T) {
	tcs := []struct {
		haystack   map[string]int
		needles    []int
		comparator func(int, int) bool
		exp        string
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, []int{4, 10}, functional.AreEqual[int], "d", nil, "Found at end"},
		{map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, []int{1, -1}, functional.AreEqual[int], "a", nil, "Found at beginning"},
		{map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, []int{3, 5}, functional.AreEqual[int], "c", nil, "Found in middle"},
		{map[string]int{"a": 1}, []int{1}, functional.AreEqual[int], "a", nil, "Found equal"},
		{map[string]int{"a": 1, "b": 2, "c": 3}, []int{4, 5}, functional.AreEqual[int], "", goaoi.ElementNotFoundError{}, "Not found"},
		{map[string]int{}, []int{1, 2, 3}, functional.AreEqual[int], "", goaoi.EmptyIterableError{}, "Super empty"},
		{map[string]int{"a": 1, "b": 2, "c": 3}, []int{}, functional.AreEqual[int], "", goaoi.EmptyIterableError{}, "Sub empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.FindFirstOfMapPred(tc.haystack, tc.needles, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_AllOfSlice(t *testing.T) {
	tcs := []struct {
		haystack   []int
		comparator func(int) bool
		exp        error
		name       string
	}{
		{[]int{1, 2}, func(i int) bool { return i > 0 }, nil, "Found"},
		{[]int{1, 2}, func(i int) bool { return i < 0 }, goaoi.ComparisonError[int, int]{0, 1}, "Not found"},
		{[]int{}, func(i int) bool { return i > 0 }, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.AllOfSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
		})
	}
}

func Test_AllOfMap(t *testing.T) {
	tcs := []struct {
		haystack   map[string]int
		comparator func(int) bool
		exp        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i > 0 }, nil, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i < 0 }, goaoi.ComparisonError[string, int]{}, "Not found"},
		{map[string]int{}, func(i int) bool { return i > 0 }, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.AllOfMap(tc.haystack, tc.comparator)

			if tc.exp == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.exp)
			}
		})
	}
}

func Test_AnyOfSlice(t *testing.T) {
	tcs := []struct {
		haystack   []int
		comparator func(int) bool
		exp        error
		name       string
	}{
		{[]int{1, 2}, func(i int) bool { return i > 0 }, nil, "Found"},
		{[]int{1, 2}, func(i int) bool { return i < 0 }, goaoi.ElementNotFoundError{}, "Not found"},
		{[]int{}, func(i int) bool { return i > 0 }, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.AnyOfSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
		})
	}
}

func Test_AnyOfMap(t *testing.T) {
	tcs := []struct {
		haystack   map[string]int
		comparator func(int) bool
		exp        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i > 0 }, nil, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i < 0 }, goaoi.ElementNotFoundError{}, "Not found"},
		{map[string]int{}, func(i int) bool { return i > 0 }, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.AnyOfMap(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
		})
	}
}

func Test_NoneOfSlice(t *testing.T) {
	tcs := []struct {
		haystack   []int
		comparator func(int) bool
		exp        error
		name       string
	}{
		{[]int{1, 2}, func(i int) bool { return i < 0 }, nil, "Found"},
		{[]int{1, 2}, func(i int) bool { return i > 0 }, goaoi.ComparisonError[int, int]{0, 1}, "Not found"},
		{[]int{}, func(i int) bool { return i > 0 }, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.NoneOfSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
		})
	}
}

func Test_NoneOfMap(t *testing.T) {
	tcs := []struct {
		haystack   map[string]int
		comparator func(int) bool
		exp        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i < 0 }, nil, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i > 0 }, goaoi.ComparisonError[string, int]{"a", 1}, "Not found"},
		{map[string]int{}, func(i int) bool { return i > 0 }, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.NoneOfMap(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
		})
	}
}

func Test_ForeachSlice(t *testing.T) {
	tcs := []struct {
		haystack   []int
		comparator func(int) error
		exp        error
		name       string
	}{
		{[]int{1, 2}, func(i int) error { i++; return nil }, nil, "Found"},
		{[]int{1, 2}, func(i int) error { return assert.AnError }, goaoi.ExecutionError[int, int]{0, 1, assert.AnError}, "Not found"},
		{[]int{}, func(i int) error { return nil }, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.ForeachSlice(tc.haystack, tc.comparator)

			assert.Equal(t, err, tc.exp)
		})
	}
}

func Test_ForeachMap(t *testing.T) {
	tcs := []struct {
		haystack   map[string]int
		comparator func(int) error
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) error { i++; return nil }, nil, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(i int) error { return assert.AnError }, goaoi.ElementNotFoundError{}, "Not found"},
		{map[string]int{}, func(i int) error { return nil }, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.ForeachMap(tc.haystack, tc.comparator)

			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}
		})
	}
}

func Test_ForeachSliceUnsafe(t *testing.T) {
	tcs := []struct {
		haystack   []int
		comparator func(int)
		exp        error
		name       string
	}{
		{[]int{1, 2}, func(i int) { i++ }, nil, "Found"},
		{[]int{}, func(i int) {}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.ForeachSliceUnsafe(tc.haystack, tc.comparator)

			assert.Equal(t, err, tc.exp)
		})
	}
}

func Test_ForeachMapUnsafe(t *testing.T) {
	tcs := []struct {
		haystack   map[string]int
		comparator func(int)
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) { i++ }, nil, "Found"},
		{map[string]int{}, func(i int) {}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.ForeachMapUnsafe(tc.haystack, tc.comparator)

			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}
		})
	}
}

func Test_CountIfSlice(t *testing.T) {
	tcs := []struct {
		haystack   []int
		comparator func(int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2}, func(i int) bool { return i > 0 }, 2, nil, "Found"},
		{[]int{1, 2}, func(i int) bool { return i < 0 }, 0, nil, "Not found"},
		{[]int{}, func(i int) bool { return i == 0 }, 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.CountIfSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_CountIfMap(t *testing.T) {
	tcs := []struct {
		haystack   map[string]int
		comparator func(int) bool
		exp        int
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i > 0 }, 2, nil, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i < 0 }, 0, nil, "Not found"},
		{map[string]int{}, func(i int) bool { return i == 0 }, 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.CountIfMap(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_MismatchSlicePred(t *testing.T) {
	tcs := []struct {
		iterable1  []int
		iterable2  []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2, 3, 4}, []int{1, 1, 3, 4}, functional.AreEqual[int], 1, nil, "Found"},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, functional.AreEqual[int], 0, goaoi.ElementNotFoundError{}, "Not found"},
		{[]int{}, []int{1, 2}, functional.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.MismatchSlicePred(tc.iterable1, tc.iterable2, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_AdjacentFindPred(t *testing.T) {
	tcs := []struct {
		haystack   []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 1, 2, 3}, functional.AreEqual[int], 0, nil, "Found at beginning"},
		{[]int{1, 2, 2}, functional.AreEqual[int], 1, nil, "Found at end"},
		{[]int{1, 2, 2, 3, 4}, functional.AreEqual[int], 1, nil, "Found in middle"},
		{[]int{1, 2, 3, 4}, functional.AreEqual[int], 0, goaoi.ElementNotFoundError{}, "Not found"},
		{[]int{}, functional.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Empty"},
		{[]int{1}, functional.AreEqual[int], 0, goaoi.ElementNotFoundError{}, "Only one element"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.AdjacentFindSlicePred(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TakeWhileSlice(t *testing.T) {
	tcs := []struct {
		original   []int
		comparator func(int) bool
		exp        []int
		err        error
		name       string
	}{
		{[]int{1, 2}, func(x int) bool { return x > 0 }, []int{1, 2}, nil, "Found"},
		{[]int{1, 2}, func(x int) bool { return x == 1 }, []int{1}, nil, "Found 1"},
		{[]int{1, 2}, func(x int) bool { return x < 0 }, []int{}, nil, "Not Found"},
		{[]int{}, func(x int) bool { return x == -1 }, []int(nil), goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.TakeWhileSlice(tc.original, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TakeWhileIterator(t *testing.T) {
	tcs := []struct {
		original   []int
		comparator func(int) bool
		exp        []int
		err        error
		name       string
	}{
		{[]int{1, 2}, func(x int) bool { return x > 0 }, []int{1, 2}, nil, "Found"},
		{[]int{1, 2}, func(x int) bool { return x == 1 }, []int{1}, nil, "Found 1"},
		{[]int{1, 2}, func(x int) bool { return x < 0 }, []int{}, nil, "Not Found"},
		{[]int{}, func(x int) bool { return x == -1 }, []int{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.TakeWhileIterator[int, int](it, tc.comparator)
			res := arraylist.NewFromIterator[int](outIter).GetSlice()

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TakeNIterator(t *testing.T) {
	tcs := []struct {
		original []int
		n        int
		exp      []int
		err      error
		name     string
	}{
		{[]int{1, 2}, 0, []int{}, nil, "take 0"},
		{[]int{1, 2}, 2, []int{1, 2}, nil, "take all"},
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{1, 2, 3}, nil, "take half"},
		{[]int{1, 2, 3}, 6, []int{1, 2, 3}, nil, "take more than exists"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.TakeNIterator[int, int](it, tc.n)
			res := arraylist.NewFromIterator[int](outIter).GetSlice()

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_DropWhileSlice(t *testing.T) {
	tcs := []struct {
		original   []int
		comparator func(int) bool
		exp        []int
		err        error
		name       string
	}{
		{[]int{1, 2}, func(x int) bool { return x < 0 }, []int{1, 2}, nil, "Found"},
		{[]int{1, 2}, func(x int) bool { return x == 1 }, []int{2}, nil, "Found 1"},
		{[]int{1, 2}, func(x int) bool { return x < 10 }, []int{}, nil, "Not Found"},
		{[]int{}, func(x int) bool { return x == -1 }, []int(nil), goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.DropWhileSlice(tc.original, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_DropWhileIterator(t *testing.T) {
	tcs := []struct {
		original   []int
		comparator func(int) bool
		exp        []int
		err        error
		name       string
	}{
		{[]int{1, 2}, func(x int) bool { return x < 0 }, []int{1, 2}, nil, "Found"},
		{[]int{1, 2}, func(x int) bool { return x == 1 }, []int{2}, nil, "Found 1"},
		{[]int{1, 2}, func(x int) bool { return x < 10 }, []int{}, nil, "Not Found"},
		{[]int{}, func(x int) bool { return x == -1 }, []int{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.DropWhileIterator[int, int](it, tc.comparator)
			res := arraylist.NewFromIterator[int](outIter).GetSlice()

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_DropNIterator(t *testing.T) {
	tcs := []struct {
		original []int
		n        int
		exp      []int
		err      error
		name     string
	}{
		{[]int{1, 2}, 2, []int{}, nil, "drop all"},
		{[]int{1, 2}, 0, []int{1, 2}, nil, "drop none"},
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{4, 5, 6}, nil, "drop half"},
		{[]int{1, 2, 3}, 6, []int{}, nil, "drop more than exists"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.DropNIterator[int, int](it, tc.n)
			res := arraylist.NewFromIterator[int](outIter).GetSlice()

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func TestStridedIterator(t *testing.T) {
	tcs := []struct {
		original []int
		n        int
		exp      []int
		err      error
		name     string
	}{
		{[]int{1, 2}, 0, []int{}, nil, "zero stride"},
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 3, 5}, nil, "every second"},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}, nil, "regual stride"},
		{[]int{1, 2, 3}, 6, []int{1}, nil, "stride larger than count"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.StridedIterator[int, int](it, tc.n)
			res := arraylist.NewFromIterator[int](outIter).GetSlice()

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func TestJoinIterator(t *testing.T) {
	tcs := []struct {
		originals [][]int
		exp       []int
		err       error
		name      string
	}{
		{[][]int{}, []int{}, nil, "empty"},
		{[][]int{{}, {}, {}}, []int{}, nil, "empty originals"},
		{[][]int{{1, 2, 3}}, []int{1, 2, 3}, nil, "single original"},
		{[][]int{{1}, {2}, {3}}, []int{1, 2, 3}, nil, "three originals with one element each"},
		{[][]int{{1, 2}, {3}, {4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}, nil, "three originals with different number of elements"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			iterators := make([]ds.ReadForIndexIterator[int, int], 0)

			for _, original := range tc.originals {
				it := arraylist.NewFromSlice(original).Begin()
				iterators = append(iterators, it)
			}

			outIter, err := goaoi.JoinIterator[int, int](iterators...)
			res := arraylist.NewFromIterator[int](outIter).GetSlice()

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TakeIfSlice(t *testing.T) {
	tcs := []struct {
		original   []int
		comparator func(int) bool
		exp        []int
		err        error
		name       string
	}{
		{[]int{1, 2}, func(x int) bool { return x != 1 }, []int{2}, nil, "Found"},
		{[]int{1, 2}, func(x int) bool { return x != -1 }, []int{1, 2}, nil, "Not Found"},
		{[]int{}, func(x int) bool { return x != -1 }, []int(nil), goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.TakeIfSlice(tc.original, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TakeIfMap(t *testing.T) {
	tcs := []struct {
		original   map[string]int
		comparator func(int) bool
		exp        map[string]int
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x != 1 }, map[string]int{"b": 2}, nil, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x != -1 }, map[string]int{"a": 1, "b": 2}, nil, "Not Found"},
		{map[string]int{}, func(x int) bool { return x != -1 }, map[string]int(nil), goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.TakeIfMap(tc.original, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TakeIfIterator(t *testing.T) {
	tcs := []struct {
		original   []int
		comparator func(int) bool
		exp        []int
		err        error
		name       string
	}{
		{[]int{1, 2}, func(x int) bool { return x > 0 }, []int{1, 2}, nil, "Found"},
		{[]int{1, 2}, func(x int) bool { return x == 1 }, []int{1}, nil, "Found 1"},
		{[]int{1, 2}, func(x int) bool { return x > 10 }, []int{}, nil, "Not Found"},
		{[]int{}, func(x int) bool { return x == -1 }, []int{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.TakeIfIterator[int, int](it, tc.comparator)
			res := arraylist.NewFromIterator[int](outIter).GetSlice()

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})

	}
}

func Test_ReplaceIfSlice(t *testing.T) {
	tcs := []struct {
		original    []int
		comparator  func(int) bool
		replacement int
		exp         []int
		err         error
		name        string
	}{
		{[]int{1, 2}, func(x int) bool { return x == 1 }, 0, []int{0, 2}, nil, "Found"},
		{[]int{1, 2}, func(x int) bool { return x == -1 }, 0, []int{1, 2}, nil, "Not Found"},
		{[]int{}, func(x int) bool { return x == -1 }, 0, []int(nil), goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.ReplaceIfSlice(tc.original, tc.comparator, tc.replacement)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_ReplaceIfMap(t *testing.T) {
	tcs := []struct {
		original    map[string]int
		comparator  func(int) bool
		replacement int
		exp         map[string]int
		err         error
		name        string
	}{
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x == 1 }, 0, map[string]int{"a": 0, "b": 2}, nil, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x == -1 }, 0, map[string]int{"a": 1, "b": 2}, nil, "Not Found"},
		{map[string]int{}, func(x int) bool { return x == -1 }, 0, map[string]int(nil), goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.ReplaceIfMap(tc.original, tc.comparator, tc.replacement)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}
func Test_ReplaceIfIterator(t *testing.T) {
	tcs := []struct {
		original    []int
		comparator  func(int) bool
		replacement int
		exp         []int
		err         error
		name        string
	}{
		{[]int{1, 2}, func(x int) bool { return x == 1 }, 0, []int{0, 2}, nil, "Found"},
		{[]int{1, 2}, func(x int) bool { return x == -1 }, 0, []int{1, 2}, nil, "Not Found"},
		{[]int{}, func(x int) bool { return x == -1 }, 0, []int{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.ReplaceIfIterator[int, int](it, tc.comparator, tc.replacement)
			res := arraylist.NewFromIterator[int](outIter).GetSlice()

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TransformSlice(t *testing.T) {
	tcs := []struct {
		original    []int
		transformer func(*int) error
		exp         []int
		err         error
		name        string
	}{
		{[]int{1, 2}, func(i *int) error { *i++; return nil }, []int{2, 3}, nil, "Found"},
		{[]int{}, func(i *int) error { *i++; return nil }, []int{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.TransformSlice(tc.original, tc.transformer)

			assert.Equal(t, tc.exp, tc.original)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TransformIterator(t *testing.T) {
	tcs := []struct {
		original    []int
		transformer func(int) (int, error)
		exp         []int
		err         error
		name        string
	}{
		{[]int{1, 2}, func(i int) (int, error) { return i + 1, nil }, []int{2, 3}, nil, "Found"},
		{[]int{}, func(i int) (int, error) { return i + 1, nil }, []int{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.TransformIterator[int, int](it, tc.transformer)
			res := arraylist.NewFromIterator[int](outIter).GetSlice()

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TransformCopySlice(t *testing.T) {
	tcs := []struct {
		original    []int
		transformer func(int) (float32, error)
		exp         []float32
		err         error
		name        string
	}{
		{[]int{1, 2}, func(i int) (float32, error) { return float32(i) + 1.0, nil }, []float32{2.0, 3.0}, nil, "Found"},
		{[]int{}, func(i int) (float32, error) { return float32(i) + 1.0, nil }, []float32{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.TransformCopySlice(tc.original, tc.transformer)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TransformMCopyap(t *testing.T) {
	tcs := []struct {
		original    map[string]int
		transformer func(int) (float32, error)
		exp         map[string]float32
		err         error
		name        string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) (float32, error) { return float32(i) + 1, nil }, map[string]float32{"a": 2.0, "b": 3.0}, nil, "Found"},
		{map[string]int{}, func(i int) (float32, error) { return float32(i) + 1, nil }, map[string]float32{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.TransformCopyMap(tc.original, tc.transformer)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TransformSCopyliceUnsafe(t *testing.T) {
	tcs := []struct {
		original    []int
		transformer func(int) float32
		exp         []float32
		err         error
		name        string
	}{
		{[]int{1, 2}, func(i int) float32 { return float32(i) + 1 }, []float32{2.0, 3.0}, nil, "Found"},
		{[]int{}, func(i int) float32 { return float32(i) + 1 }, []float32{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.TransformCopySliceUnsafe(tc.original, tc.transformer)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TransformMCopyapUnsafe(t *testing.T) {
	tcs := []struct {
		original    map[string]int
		transformer func(int) float32
		exp         map[string]float32
		err         error
		name        string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) float32 { return float32(i) + 1 }, map[string]float32{"a": 2.0, "b": 3.0}, nil, "Found"},
		{map[string]int{}, func(i int) float32 { return float32(i) + 1 }, map[string]float32{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res, err := goaoi.TransformCopyMapUnsafe(tc.original, tc.transformer)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TransformMap(t *testing.T) {
	tcs := []struct {
		original    map[string]int
		transformer func(int) (int, error)
		exp         map[string]int
		err         error
		name        string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) (int, error) { return i + 1, nil }, map[string]int{"a": 2, "b": 3}, nil, "Found"},
		{map[string]int{}, func(i int) (int, error) { return i + 1, nil }, map[string]int{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.TransformMap(tc.original, tc.transformer)

			assert.Equal(t, tc.exp, tc.original)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TransformSliceUnsafe(t *testing.T) {
	tcs := []struct {
		original    []int
		transformer func(*int)
		exp         []int
		err         error
		name        string
	}{
		{[]int{1, 2}, func(i *int) { *i++ }, []int{2, 3}, nil, "Found"},
		{[]int{}, func(i *int) { *i++ }, []int{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.TransformSliceUnsafe(tc.original, tc.transformer)

			assert.Equal(t, tc.exp, tc.original)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TransformMapUnsafe(t *testing.T) {
	tcs := []struct {
		original    map[string]int
		transformer func(int) int
		exp         map[string]int
		err         error
		name        string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) int { return i + 1 }, map[string]int{"a": 2, "b": 3}, nil, "Found"},
		{map[string]int{}, func(i int) int { return i + 1 }, map[string]int{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := goaoi.TransformMapUnsafe(tc.original, tc.transformer)

			assert.Equal(t, tc.exp, tc.original)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_TransformIteratorUnsafe(t *testing.T) {
	tcs := []struct {
		original    []int
		transformer func(int) int
		exp         []int
		err         error
		name        string
	}{
		{[]int{1, 2}, func(i int) int { return i + 1 }, []int{2, 3}, nil, "Found"},
		{[]int{}, func(i int) int { return i + 1 }, []int{}, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.TransformIteratorUnsafe[int, int](it, tc.transformer)
			res := arraylist.NewFromIterator[int](outIter).GetSlice()

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_FillSlice(t *testing.T) {
	tcs := []struct {
		original []int
		filler   int
		exp      []int
		name     string
	}{
		{make([]int, 4), 1, []int{1, 1, 1, 1}, "Found, len eq cap"},
		{make([]int, 0, 4), 1, []int{1, 1, 1, 1}, "Found, len neq cap"},
		{make([]int, 0), 1, []int{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			res := goaoi.FillSlice(&tc.original, tc.filler)

			assert.Equal(t, tc.original, res)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func Test_MinSlicePred(t *testing.T) {
	tcs := []struct {
		haystack   []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2, 3}, functional.IsLessThan[int], 1, nil, "Found at beginning"},
		{[]int{2, 3, 1}, functional.IsLessThan[int], 1, nil, "Found at end"},
		{[]int{}, functional.IsLessThan[int], 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, res, err := goaoi.MinSlicePred(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_MinMapPred(t *testing.T) {
	tcs := []struct {
		haystack   map[string]int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3}, functional.IsLessThan[int], 1, nil, "Found at beginning"},
		{map[string]int{"a": 2, "b": 3, "c": 1}, functional.IsLessThan[int], 1, nil, "Found at end"},
		{map[string]int{}, functional.IsLessThan[int], 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, res, err := goaoi.MinMapPred(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_MaxSlicePred(t *testing.T) {
	tcs := []struct {
		haystack   []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2, 3}, functional.IsGreaterThan[int], 3, nil, "Found at beginning"},
		{[]int{2, 3, 1}, functional.IsGreaterThan[int], 3, nil, "Found at end"},
		{[]int{}, functional.IsGreaterThan[int], 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, res, err := goaoi.MaxSlicePred(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_MaxMapPred(t *testing.T) {
	tcs := []struct {
		haystack   map[string]int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3}, functional.IsGreaterThan[int], 3, nil, "Found at beginning"},
		{map[string]int{"a": 2, "b": 3, "c": 1}, functional.IsGreaterThan[int], 3, nil, "Found at end"},
		{map[string]int{}, functional.IsGreaterThan[int], 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, res, err := goaoi.MaxMapPred(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_MinMaxSlicePred(t *testing.T) {
	tcs := []struct {
		haystack       []int
		comparator_min func(int, int) bool
		comparator_max func(int, int) bool
		exp_min        int
		exp_max        int
		err            error
		name           string
	}{
		{[]int{1, 2, 3}, functional.IsLessThan[int], functional.IsGreaterThan[int], 1, 3, nil, "Found at beginning"},
		{[]int{2, 3, 1}, functional.IsLessThan[int], functional.IsGreaterThan[int], 1, 3, nil, "Found at end"},
		{[]int{}, functional.IsLessThan[int], functional.IsGreaterThan[int], 0, 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, _, min, max, err := goaoi.MinMaxSlicePred(tc.haystack, tc.comparator_min, tc.comparator_max)

			assert.Equal(t, tc.exp_min, min)
			assert.Equal(t, tc.exp_max, max)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_MinMaxMapPred(t *testing.T) {
	tcs := []struct {
		haystack       map[string]int
		comparator_min func(int, int) bool
		comparator_max func(int, int) bool
		exp_min        int
		exp_max        int
		err            error
		name           string
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3}, functional.IsLessThan[int], functional.IsGreaterThan[int], 1, 3, nil, "Found at beginning"},
		{map[string]int{"a": 2, "b": 3, "c": 1}, functional.IsLessThan[int], functional.IsGreaterThan[int], 1, 3, nil, "Found at end"},
		{map[string]int{}, functional.IsLessThan[int], functional.IsGreaterThan[int], 0, 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, _, min, max, err := goaoi.MinMaxMapPred(tc.haystack, tc.comparator_min, tc.comparator_max)

			assert.Equal(t, tc.exp_min, min)
			assert.Equal(t, tc.exp_max, max)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}
