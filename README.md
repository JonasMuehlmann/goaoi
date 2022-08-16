[![codecov](https://codecov.io/gh/JonasMuehlmann/goaoi/branch/main/graph/badge.svg?token=0J2P8OAJ6Y)](https://codecov.io/gh/JonasMuehlmann/goaoi)
# goaoi
![img](https://img.shields.io/badge/semver-2.0.0-green) [![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

Conventient algorithms for processing iterables, inspired by the algorithm header from the C++ standard template library (STL for short).
Note that `*Iterator()` methods return lazy iterators wrapping the underlying ones instead of changing underlying data or allocation new data.

Please do not expect a stable API at this point.

Sister project: https://github.com/JonasMuehlmann/pyaoi

## Installation

```go get github.com/JonasMuehlmann/goaoi```

## How to use

All functions live in the ```goaoi``` namespace and most have separate implementations for maps, slices, strings and https://github.com/JonasMuehlmann/datastructures.go iterators.
For usage examples, refer to the test files.

API Documentation available at https://pkg.go.dev/github.com/JonasMuehlmann/goaoi.

### Predicates and functional operators
package [`functional`](https://pkg.go.dev/github.com/JonasMuehlmann/goaoi/functional) provides partially specializable predicates and functional operators.

Example:
```go
import (
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/goaoi/functional"
)


// Result: 15
sum := goaoi.AccumulateSlice([]int{1,2,3,4,5}, 0, functional.Add)
// Result: 1, nil
i, err := goaoi.FindIfSlice([]int{1,3,0,1,4,5}, 0, functional.AreEqualPartial(3))
```

### Lazy iterator adapters
package [`iteratoradapters`](https://pkg.go.dev/github.com/JonasMuehlmann/goaoi/iterator_adapters) provides lazy iterator adapters for efficient iterator processing.
The adapters wrap underlying ones and avoid altering, copying or allocating data.
This is especially useful for chaining them on the same container like this (untested example without fully implemented API):
```go
import (
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/goaoi/functional"
	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
)

// NOTE: Complexities refer to space and or time

// O(1)
valuesOrig := arraylist.NewFromSlice([]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15})
// O(1) because lazy, would be O(N) otherwise
valid := valuesOrig.Begin().TakeWhileIterator(functional.IsLessThanEqualPartial(12))

// O(m) because lazy, would be O(N) otherwise
m := 4
parts := increased.SplitNthIterator(m)

newParts := make([]goaoi.ReadForIndexIterator[int, int], len(parts))
partsIter := parts.Beign()
for partsIter.Next() {
    // Pretend the input is larger and this actually makes sense.
    go func() {
        part, _ := partsIter.Value()
        // O(1) because lazy, would be O(N) otherwise
        newparts = append(newParts, part.TransformIterator(functional.AddPartial(5))
    }
}

// O(m) because lazy, would be O(N)
joined := goaoi.JoinIterator(newParts...)
// O(n) because the adapter's functionality need to be applied 
// for materialization of the new data.
// valuesOrig left unchanged
valuesAfterCopy := arraylist.NewFromIterator(joined)
```

### Lazy generators
package [`generators`](https://pkg.go.dev/github.com/JonasMuehlmann/goaoi/generators) provides lazy generators, which work similar to iterator adapters, but they do not reference existing data, instead they generate it themselves (lazily).
Some generators have no defined end.

Example:
```go
import (
	"github.com/JonasMuehlmann/goaoi"
	"github.com/JonasMuehlmann/goaoi/functional"
	"github.com/JonasMuehlmann/goaoi/generators"
	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
)

// Generates 10 1s
repeater := generators.NewRepeat(1, 10)

// [1,1,1,1,1,1,1,1,1,1]
firstValues := arraylist.NewFromIterator(repeater)

// Generates infinite 1s
infiniteRepeater := generators.NewRepeat(1, -1)

// Infinite loop until OOM (Out of memory)
firstValues := arraylist.NewFromIterator(infiniteRepeater)
```

## License
Copyright (C) 2021-2022 [Jonas Muehlmann](https://github.com/JonasMuehlmann)
 
The project is licensed under the terms of the MIT license, you can view it [here](LICENSE.md).
