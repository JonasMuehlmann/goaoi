package goaoi_test

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/stretchr/testify/assert"
)

func Test_FindIfSlice(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

	tcs := []struct {
		super      []int
		sub        []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2, 3, 4}, []int{3, 4}, goaoi.AreEqual[int], 2, nil, "Found at end"},
		{[]int{1, 2, 3, 4}, []int{1, 2}, goaoi.AreEqual[int], 0, nil, "Found at beginning"},
		{[]int{1, 2, 3, 4}, []int{2, 3}, goaoi.AreEqual[int], 1, nil, "Found in middle"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, goaoi.AreEqual[int], 0, nil, "Found equal"},
		{[]int{1, 2, 3}, []int{1, 4}, goaoi.AreEqual[int], 0, goaoi.ElementNotFoundError{}, "Not found"},
		{[]int{}, []int{1, 2, 3}, goaoi.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Super empty"},
		{[]int{1, 2, 3}, []int{}, goaoi.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Sub empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

	tcs := []struct {
		haystack   []int
		needles    []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2, 3, 4}, []int{4, 10}, goaoi.AreEqual[int], 3, nil, "Found at end"},
		{[]int{1, 2, 3, 4}, []int{1, 2}, goaoi.AreEqual[int], 0, nil, "Found at beginning"},
		{[]int{1, 2, 3, 4}, []int{3, 4}, goaoi.AreEqual[int], 2, nil, "Found in middle"},
		{[]int{1}, []int{1}, goaoi.AreEqual[int], 0, nil, "Found equal"},
		{[]int{1, 2, 3}, []int{4, 5}, goaoi.AreEqual[int], 0, goaoi.ElementNotFoundError{}, "Not found"},
		{[]int{}, []int{1, 2, 3}, goaoi.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Super empty"},
		{[]int{1, 2, 3}, []int{}, goaoi.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Sub empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

	tcs := []struct {
		haystack   map[string]int
		needles    []int
		comparator func(int, int) bool
		exp        string
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, []int{4, 10}, goaoi.AreEqual[int], "d", nil, "Found at end"},
		{map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, []int{1, -1}, goaoi.AreEqual[int], "a", nil, "Found at beginning"},
		{map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, []int{3, 5}, goaoi.AreEqual[int], "c", nil, "Found in middle"},
		{map[string]int{"a": 1}, []int{1}, goaoi.AreEqual[int], "a", nil, "Found equal"},
		{map[string]int{"a": 1, "b": 2, "c": 3}, []int{4, 5}, goaoi.AreEqual[int], "", goaoi.ElementNotFoundError{}, "Not found"},
		{map[string]int{}, []int{1, 2, 3}, goaoi.AreEqual[int], "", goaoi.EmptyIterableError{}, "Super empty"},
		{map[string]int{"a": 1, "b": 2, "c": 3}, []int{}, goaoi.AreEqual[int], "", goaoi.EmptyIterableError{}, "Sub empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.AllOfSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
		})
	}
}

func Test_AllOfMap(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.AnyOfSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
		})
	}
}

func Test_AnyOfMap(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.AnyOfMap(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
		})
	}
}

func Test_NoneOfSlice(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.NoneOfSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
		})
	}
}

func Test_NoneOfMap(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.NoneOfMap(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
		})
	}
}

func Test_ForeachSlice(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.ForeachSlice(tc.haystack, tc.comparator)

			assert.Equal(t, err, tc.exp)
		})
	}
}

func Test_ForeachMap(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.ForeachSliceUnsafe(tc.haystack, tc.comparator)

			assert.Equal(t, err, tc.exp)
		})
	}
}

func Test_ForeachMapUnsafe(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

	tcs := []struct {
		iterable1  []int
		iterable2  []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2, 3, 4}, []int{1, 1, 3, 4}, goaoi.AreEqual[int], 1, nil, "Found"},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, goaoi.AreEqual[int], 0, goaoi.ElementNotFoundError{}, "Not found"},
		{[]int{}, []int{1, 2}, goaoi.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

	tcs := []struct {
		haystack   []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 1, 2, 3}, goaoi.AreEqual[int], 0, nil, "Found at beginning"},
		{[]int{1, 2, 2}, goaoi.AreEqual[int], 1, nil, "Found at end"},
		{[]int{1, 2, 2, 3, 4}, goaoi.AreEqual[int], 1, nil, "Found in middle"},
		{[]int{1, 2, 3, 4}, goaoi.AreEqual[int], 0, goaoi.ElementNotFoundError{}, "Not found"},
		{[]int{}, goaoi.AreEqual[int], 0, goaoi.EmptyIterableError{}, "Empty"},
		{[]int{1}, goaoi.AreEqual[int], 0, goaoi.ElementNotFoundError{}, "Only one element"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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

func Test_DropWhileSlice(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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

func Test_CopyIfSlice(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyIfSlice(tc.original, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_CopyIfMap(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyIfMap(tc.original, tc.comparator)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_CopyIfIterator(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.CopyIfIterator[int, int](it, tc.comparator)
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

func Test_CopyReplaceIfSlice(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyReplaceIfSlice(tc.original, tc.comparator, tc.replacement)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}

func Test_CopyReplaceIfMap(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyReplaceIfMap(tc.original, tc.comparator, tc.replacement)

			assert.Equal(t, tc.exp, res)
			if tc.err == nil {
				assert.Nil(t, err)
			} else {
				assert.ErrorAs(t, err, &tc.err)
			}

		})
	}
}
func Test_CopyReplaceIfIterator(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			it := arraylist.NewFromSlice(tc.original).Begin()
			outIter, err := goaoi.CopyReplaceIfIterator[int, int](it, tc.comparator, tc.replacement)
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
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

func Test_FillSlice(t *testing.T) {
	t.Parallel()

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
		t.Run(tc.name, func(t *testing.T) {
			res := goaoi.FillSlice(&tc.original, tc.filler)

			assert.Equal(t, tc.original, res)
			assert.Equal(t, tc.exp, res)
		})
	}
}

func Test_MinSlicePred(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack   []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2, 3}, goaoi.IsLessThan[int], 1, nil, "Found at beginning"},
		{[]int{2, 3, 1}, goaoi.IsLessThan[int], 1, nil, "Found at end"},
		{[]int{}, goaoi.IsLessThan[int], 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

	tcs := []struct {
		haystack   map[string]int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3}, goaoi.IsLessThan[int], 1, nil, "Found at beginning"},
		{map[string]int{"a": 2, "b": 3, "c": 1}, goaoi.IsLessThan[int], 1, nil, "Found at end"},
		{map[string]int{}, goaoi.IsLessThan[int], 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

	tcs := []struct {
		haystack   []int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{[]int{1, 2, 3}, goaoi.IsGreaterThan[int], 3, nil, "Found at beginning"},
		{[]int{2, 3, 1}, goaoi.IsGreaterThan[int], 3, nil, "Found at end"},
		{[]int{}, goaoi.IsGreaterThan[int], 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

	tcs := []struct {
		haystack   map[string]int
		comparator func(int, int) bool
		exp        int
		err        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3}, goaoi.IsGreaterThan[int], 3, nil, "Found at beginning"},
		{map[string]int{"a": 2, "b": 3, "c": 1}, goaoi.IsGreaterThan[int], 3, nil, "Found at end"},
		{map[string]int{}, goaoi.IsGreaterThan[int], 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

	tcs := []struct {
		haystack       []int
		comparator_min func(int, int) bool
		comparator_max func(int, int) bool
		exp_min        int
		exp_max        int
		err            error
		name           string
	}{
		{[]int{1, 2, 3}, goaoi.IsLessThan[int], goaoi.IsGreaterThan[int], 1, 3, nil, "Found at beginning"},
		{[]int{2, 3, 1}, goaoi.IsLessThan[int], goaoi.IsGreaterThan[int], 1, 3, nil, "Found at end"},
		{[]int{}, goaoi.IsLessThan[int], goaoi.IsGreaterThan[int], 0, 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
	t.Parallel()

	tcs := []struct {
		haystack       map[string]int
		comparator_min func(int, int) bool
		comparator_max func(int, int) bool
		exp_min        int
		exp_max        int
		err            error
		name           string
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3}, goaoi.IsLessThan[int], goaoi.IsGreaterThan[int], 1, 3, nil, "Found at beginning"},
		{map[string]int{"a": 2, "b": 3, "c": 1}, goaoi.IsLessThan[int], goaoi.IsGreaterThan[int], 1, 3, nil, "Found at end"},
		{map[string]int{}, goaoi.IsLessThan[int], goaoi.IsGreaterThan[int], 0, 0, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
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
