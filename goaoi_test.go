package goaoi_test

import (
	"testing"

	"github.com/JonasMuehlmann/goaoi"
	"github.com/stretchr/testify/assert"
)

func Test_FindSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack []int
		needle   int
		exp      int
		hasError bool
		name     string
	}{
		{[]int{1, 2}, 1, 0, false, "Found"},
		{[]int{1, 2}, -1, 0, true, "Not found"},
		{[]int{}, -1, 0, true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.FindSlice(tc.haystack, tc.needle)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_FindIfSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack   []int
		comparator func(int) bool
		exp        int
		hasError   bool
		name       string
	}{
		{[]int{1, 2}, func(i int) bool { return i == 1 }, 0, false, "Found"},
		{[]int{1, 2}, func(i int) bool { return i == 0 }, 0, true, "Not found"},
		{[]int{}, func(i int) bool { return i == 0 }, 0, true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.FindIfSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_FindIfMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack   map[string]int
		comparator func(int) bool
		exp        string
		hasError   bool
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i == 1 }, "a", false, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i == 0 }, "", true, "Not found"},
		{map[string]int{}, func(i int) bool { return i == 0 }, "", true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.FindIfMap(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_FindEndSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		super      []int
		sub        []int
		comparator func(int, int) bool
		exp        int
		hasError   bool
		name       string
	}{
		{[]int{1, 2, 3, 4}, []int{3, 4}, goaoi.AreEqual[int], 2, false, "Found at end"},
		{[]int{1, 2, 3, 4}, []int{1, 2}, goaoi.AreEqual[int], 0, false, "Found at beginning"},
		{[]int{1, 2, 3, 4}, []int{2, 3}, goaoi.AreEqual[int], 1, false, "Found in middle"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, goaoi.AreEqual[int], 0, false, "Found equal"},
		{[]int{1, 2, 3}, []int{1, 4}, goaoi.AreEqual[int], 0, true, "Not found"},
		{[]int{}, []int{1, 2, 3}, goaoi.AreEqual[int], 0, true, "Super empty"},
		{[]int{1, 2, 3}, []int{}, goaoi.AreEqual[int], 0, true, "Sub empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.FindEndSlice(tc.super, tc.sub, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_FindFirstOfSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack   []int
		needles    []int
		comparator func(int, int) bool
		exp        int
		hasError   bool
		name       string
	}{
		{[]int{1, 2, 3, 4}, []int{4, 10}, goaoi.AreEqual[int], 3, false, "Found at end"},
		{[]int{1, 2, 3, 4}, []int{1, 2}, goaoi.AreEqual[int], 0, false, "Found at beginning"},
		{[]int{1, 2, 3, 4}, []int{3, 4}, goaoi.AreEqual[int], 2, false, "Found in middle"},
		{[]int{1}, []int{1}, goaoi.AreEqual[int], 0, false, "Found equal"},
		{[]int{1, 2, 3}, []int{4, 5}, goaoi.AreEqual[int], 0, true, "Not found"},
		{[]int{}, []int{1, 2, 3}, goaoi.AreEqual[int], 0, true, "Super empty"},
		{[]int{1, 2, 3}, []int{}, goaoi.AreEqual[int], 0, true, "Sub empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.FindFirstOfSlice(tc.haystack, tc.needles, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_FindFirstOfMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack   map[string]int
		needles    []int
		comparator func(int, int) bool
		exp        string
		hasError   bool
		name       string
	}{
		{map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, []int{4, 10}, goaoi.AreEqual[int], "d", false, "Found at end"},
		{map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, []int{1, 2}, goaoi.AreEqual[int], "a", false, "Found at beginning"},
		{map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, []int{3, 4}, goaoi.AreEqual[int], "c", false, "Found in middle"},
		{map[string]int{"a": 1}, []int{1}, goaoi.AreEqual[int], "a", false, "Found equal"},
		{map[string]int{"a": 1, "b": 2, "c": 3}, []int{4, 5}, goaoi.AreEqual[int], "", true, "Not found"},
		{map[string]int{}, []int{1, 2, 3}, goaoi.AreEqual[int], "", true, "Super empty"},
		{map[string]int{"a": 1, "b": 2, "c": 3}, []int{}, goaoi.AreEqual[int], "", true, "Sub empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.FindFirstOfMap(tc.haystack, tc.needles, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
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
		{[]int{1, 2}, func(i int) bool { return i < 0 }, goaoi.ComparisonError[int]{0}, "Not found"},
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
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i < 0 }, goaoi.ComparisonError[string]{"a"}, "Not found"},
		{map[string]int{}, func(i int) bool { return i > 0 }, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.AllOfMap(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, err)
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
		{[]int{1, 2}, func(i int) bool { return i < 0 }, goaoi.ComparisonError[int]{1}, "Not found"},
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
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i < 0 }, goaoi.ComparisonError[string]{}, "Not found"},
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
		{[]int{1, 2}, func(i int) bool { return i > 0 }, goaoi.ComparisonError[int]{1}, "Not found"},
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
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i > 0 }, goaoi.ComparisonError[string]{}, "Not found"},
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
		{[]int{1, 2}, func(i int) error { return assert.AnError }, goaoi.ExecutionError[int]{0, assert.AnError}, "Not found"},
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
		hasError   bool
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) error { i++; return nil }, false, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(i int) error { return assert.AnError }, true, "Not found"},
		{map[string]int{}, func(i int) error { return nil }, true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.ForeachMap(tc.haystack, tc.comparator)

			assert.Equal(t, err != nil, tc.hasError)
		})
	}
}

func Test_CountSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack []int
		wanted   int
		exp      int
		hasError bool
		name     string
	}{
		{[]int{1, 1}, 1, 2, false, "Found"},
		{[]int{2, 2}, 1, 0, false, "Not found"},
		{[]int{}, 1, 0, true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CountSlice(tc.haystack, tc.wanted)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CountMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack map[string]int
		wanted   int
		exp      int
		hasError bool
		name     string
	}{
		{map[string]int{"a": 1, "b": 1}, 1, 2, false, "Found"},
		{map[string]int{"a": 2, "b": 2}, 1, 0, false, "Not found"},
		{map[string]int{}, 1, 0, true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CountMap(tc.haystack, tc.wanted)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CountIfSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack   []int
		comparator func(int) bool
		exp        int
		hasError   bool
		name       string
	}{
		{[]int{1, 2}, func(i int) bool { return i > 0 }, 2, false, "Found"},
		{[]int{1, 2}, func(i int) bool { return i < 0 }, 0, false, "Not found"},
		{[]int{}, func(i int) bool { return i == 0 }, 0, true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CountIfSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CountIfMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack   map[string]int
		comparator func(int) bool
		exp        int
		hasError   bool
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i > 0 }, 2, false, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(i int) bool { return i < 0 }, 0, false, "Not found"},
		{map[string]int{}, func(i int) bool { return i == 0 }, 0, true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CountIfMap(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_MismatchSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		iterable1 []int
		iterable2 []int
		exp       int
		hasError  bool
		name      string
	}{
		{[]int{1, 2, 3, 4}, []int{1, 1, 3, 4}, 1, false, "Found"},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}, 0, true, "Not found"},
		{[]int{}, []int{1, 2}, 0, true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.MismatchSlice(tc.iterable1, tc.iterable2)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_AdjacentFind(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		haystack   []int
		comparator func(int, int) bool
		exp        int
		hasError   bool
		name       string
	}{
		{[]int{1, 1, 2, 3}, goaoi.AreEqual[int], 0, false, "Found at beginning"},
		{[]int{1, 2, 2}, goaoi.AreEqual[int], 1, false, "Found at end"},
		{[]int{1, 2, 2, 3, 4}, goaoi.AreEqual[int], 1, false, "Found in middle"},
		{[]int{1, 2, 3, 4}, goaoi.AreEqual[int], 0, true, "Not found"},
		{[]int{}, goaoi.AreEqual[int], 0, true, "Empty"},
		{[]int{1}, goaoi.AreEqual[int], 0, true, "Only one element"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.AdjacentFindSlice(tc.haystack, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CopyReplaceSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original    []int
		toReplace   int
		replacement int
		exp         []int
		hasError    bool
		name        string
	}{
		{[]int{1, 2}, 1, 0, []int{0, 2}, false, "Found"},
		{[]int{1, 2}, -1, 0, []int{1, 2}, false, "Not Found"},
		{[]int{}, 1, 0, []int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyReplaceSlice(tc.original, tc.toReplace, tc.replacement)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CopyReplaceMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original    map[string]int
		toReplace   int
		replacement int
		exp         map[string]int
		hasError    bool
		name        string
	}{
		{map[string]int{"a": 1, "b": 2}, 1, 0, map[string]int{"a": 0, "b": 2}, false, "Found"},
		{map[string]int{"a": 1, "b": 2}, -1, 0, map[string]int{"a": 1, "b": 2}, false, "Not Found"},
		{map[string]int{}, -1, 0, map[string]int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyReplaceMap(tc.original, tc.toReplace, tc.replacement)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
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
		hasError    bool
		name        string
	}{
		{[]int{1, 2}, func(x int) bool { return x == 1 }, 0, []int{0, 2}, false, "Found"},
		{[]int{1, 2}, func(x int) bool { return x == -1 }, 0, []int{1, 2}, false, "Not Found"},
		{[]int{}, func(x int) bool { return x == -1 }, 0, []int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyReplaceIfSlice(tc.original, tc.comparator, tc.replacement)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
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
		hasError    bool
		name        string
	}{
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x == 1 }, 0, map[string]int{"a": 0, "b": 2}, false, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x == -1 }, 0, map[string]int{"a": 1, "b": 2}, false, "Not Found"},
		{map[string]int{}, func(x int) bool { return x == -1 }, 0, map[string]int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyReplaceIfMap(tc.original, tc.comparator, tc.replacement)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CopyReplaceIfNotSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original    []int
		comparator  func(int) bool
		replacement int
		exp         []int
		hasError    bool
		name        string
	}{
		{[]int{1, 2}, func(x int) bool { return x > 1 }, 0, []int{0, 2}, false, "Found"},
		{[]int{1, 2}, func(x int) bool { return x > -1 }, 0, []int{1, 2}, false, "Not Found"},
		{[]int{}, func(x int) bool { return x > -1 }, 0, []int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyReplaceIfNotSlice(tc.original, tc.comparator, tc.replacement)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CopyReplaceIfNotMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original    map[string]int
		comparator  func(int) bool
		replacement int
		exp         map[string]int
		hasError    bool
		name        string
	}{
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x > 1 }, 0, map[string]int{"a": 0, "b": 2}, false, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x > -1 }, 0, map[string]int{"a": 1, "b": 2}, false, "Not Found"},
		{map[string]int{}, func(x int) bool { return x > -1 }, 0, map[string]int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyReplaceIfNotMap(tc.original, tc.comparator, tc.replacement)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CopyExceptSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original  []int
		toExclude int
		exp       []int
		hasError  bool
		name      string
	}{
		{[]int{1, 2}, 1, []int{2}, false, "Found"},
		{[]int{1, 2}, -1, []int{1, 2}, false, "Not Found"},
		{[]int{}, 1, []int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyExceptSlice(tc.original, tc.toExclude)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CopyExceptMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original  map[string]int
		toExclude int
		exp       map[string]int
		hasError  bool
		name      string
	}{
		{map[string]int{"a": 1, "b": 2}, 1, map[string]int{"b": 2}, false, "Found"},
		{map[string]int{"a": 1, "b": 2}, -1, map[string]int{"a": 1, "b": 2}, false, "Not Found"},
		{map[string]int{}, -1, map[string]int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyExceptMap(tc.original, tc.toExclude)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CopyExceptIfSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original   []int
		comparator func(int) bool
		exp        []int
		hasError   bool
		name       string
	}{
		{[]int{1, 2}, func(x int) bool { return x == 1 }, []int{2}, false, "Found"},
		{[]int{1, 2}, func(x int) bool { return x == -1 }, []int{1, 2}, false, "Not Found"},
		{[]int{}, func(x int) bool { return x == -1 }, []int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyExceptIfSlice(tc.original, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CopyExceptIfMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original   map[string]int
		comparator func(int) bool
		exp        map[string]int
		hasError   bool
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x == 1 }, map[string]int{"b": 2}, false, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x == -1 }, map[string]int{"a": 1, "b": 2}, false, "Not Found"},
		{map[string]int{}, func(x int) bool { return x == -1 }, map[string]int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyExceptIfMap(tc.original, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CopyExceptIfNotSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original   []int
		comparator func(int) bool
		exp        []int
		hasError   bool
		name       string
	}{
		{[]int{1, 2}, func(x int) bool { return x > 1 }, []int{2}, false, "Found"},
		{[]int{1, 2}, func(x int) bool { return x > -1 }, []int{1, 2}, false, "Not Found"},
		{[]int{}, func(x int) bool { return x > -1 }, []int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyExceptIfNotSlice(tc.original, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_CopyExceptIfNotMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original   map[string]int
		comparator func(int) bool
		exp        map[string]int
		hasError   bool
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x > 1 }, map[string]int{"b": 2}, false, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(x int) bool { return x > -1 }, map[string]int{"a": 1, "b": 2}, false, "Not Found"},
		{map[string]int{}, func(x int) bool { return x > -1 }, map[string]int(nil), true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			res, err := goaoi.CopyExceptIfNotMap(tc.original, tc.comparator)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_TransformSlice(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original    []int
		transformer func(*int) error
		exp         []int
		hasError    bool
		name        string
	}{
		{[]int{1, 2}, func(i *int) error { *i++; return nil }, []int{2, 3}, false, "Found"},
		{[]int{}, func(i *int) error { *i++; return nil }, []int{}, true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.TransformSlice(tc.original, tc.transformer)

			assert.Equal(t, tc.exp, tc.original)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}

func Test_TransformMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		original    map[string]int
		transformer func(int) (int, error)
		exp         map[string]int
		hasError    bool
		name        string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) (int, error) { return i + 1, nil }, map[string]int{"a": 2, "b": 3}, false, "Found"},
		{map[string]int{}, func(i int) (int, error) { return i + 1, nil }, map[string]int{}, true, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.TransformMap(tc.original, tc.transformer)

			assert.Equal(t, tc.exp, tc.original)
			assert.Equal(t, tc.hasError, err != nil)
		})
	}
}
