package generators_test

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
	testCommon "github.com/JonasMuehlmann/datastructures.go/tests"
	"github.com/JonasMuehlmann/goaoi/generators"
	"github.com/stretchr/testify/assert"
)

func Test_RangeGenerator(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name   string
		args   []int
		output []int
		error  string
	}{
		{name: "no args", args: []int{}, output: []int{}},
		{name: "one arg, negative", args: []int{-5}, output: []int{}, error: generators.ErrorNegativeRange},
		{name: "one arg, implicit 0 to 0", args: []int{0}, output: []int{0}},
		{name: "one arg, implicit 0 to 1", args: []int{1}, output: []int{0, 1}},
		{name: "one arg, implicit 0 to 10", args: []int{10}, output: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},

		{name: "two args, equal", args: []int{1, 1}, output: []int{1}},
		{name: "two args, stop before start", args: []int{1, -1}, output: []int{}, error: generators.ErrorNegativeRange},
		{name: "two args, negative start", args: []int{-1, 1}, output: []int{-1, 0, 1}},
		{name: "two args, negative start, negative stop", args: []int{-5, -2}, output: []int{-5, -4, -3, -2}},

		{name: "three args, equal", args: []int{1, 1, 1}, output: []int{1}},
		{name: "three args, negative step", args: []int{1, 5, -1}, output: []int{}, error: generators.ErrorNegativeRange},
		{name: "three args, 'unclean' step", args: []int{1, 5, 3}, output: []int{}, error: generators.ErrorUncleanStep},
		{name: "three args, step higher than stop - start", args: []int{1, 3, 5}, output: []int{}, error: generators.ErrorUncleanStep},
		{name: "three args, proper stepping", args: []int{1, 5, 2}, output: []int{1, 3, 5}},
		{name: "four args", args: []int{1, 5, 2, 3}, output: []int{}, error: generators.ErrorTooManyArgs},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			defer testCommon.HandlePanic(t, tc.name)

			if tc.error != "" {
				assert.PanicsWithValuef(t, tc.error, func() {
					generators.NewRange(tc.args...)
				}, tc.name, " construction error")
			} else {
				out := generators.NewRange(tc.args...)
				result := arraylist.NewFromIterator[int](out).GetSlice()

				assert.Equalf(t, tc.output, result, tc.name, " output")
				assert.Equalf(t, len(tc.output), out.Size(), tc.name, " suze")
			}
		})
	}
}
