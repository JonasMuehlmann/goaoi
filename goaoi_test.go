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
		// {[]int{1, 2, 3, 4}, []int{3, 4}, goaoi.AreEqual[int], 2, false, "Found at end"},
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
