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
			assert.Equal(t, err != nil, tc.hasError)
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
			assert.Equal(t, err != nil, tc.hasError)
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
			assert.Equal(t, err != nil, tc.hasError)
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
		exp        error
		name       string
	}{
		{map[string]int{"a": 1, "b": 2}, func(i int) error { i++; return nil }, nil, "Found"},
		{map[string]int{"a": 1, "b": 2}, func(i int) error { return assert.AnError }, goaoi.ExecutionError[string]{"a", assert.AnError}, "Not found"},
		{map[string]int{}, func(i int) error { return nil }, goaoi.EmptyIterableError{}, "Empty"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := goaoi.ForeachMap(tc.haystack, tc.comparator)

			assert.Equal(t, err, tc.exp)
		})
	}
}
