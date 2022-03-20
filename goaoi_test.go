package goaoi_test

import (
	"testing"

	"github.com/JonasMuehlmann/goaoi"
	"github.com/stretchr/testify/assert"
)

func Test_Find(t *testing.T) {
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
			res, err := goaoi.Find(tc.haystack, tc.needle)

			assert.Equal(t, tc.exp, res)
			assert.Equal(t, err != nil, tc.hasError)
		})
	}
}
