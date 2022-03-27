package goaoi

import "errors"

func FindSlice[T comparable](haystack []T, needle T) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, value := range haystack {
		if value == needle {
			return i, nil
		}
	}

	return 0, errors.New("Could not find element")
}

func FindIfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) (TKey, error) {
	var zeroVal TKey

	if len(haystack) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	for key, value := range haystack {
		if comparator(value) {
			return key, nil
		}
	}

	return zeroVal, errors.New("Could not find element")
}

func FindIfSlice[T comparable](haystack []T, comparator func(T) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, value := range haystack {
		if comparator(value) {
			return i, nil
		}
	}

	return 0, errors.New("Could not find element")
}

func FindEndSlice[T comparable](super []T, sub []T, comparator func(T, T) bool) (int, error) {
	if len(super) == 0 || len(sub) == 0 {
		return 0, EmptyIterableError{}
	}
OUTER:
	for i := len(super) - 1; i >= len(sub)-1; i-- {
		for j := 0; j < len(sub); j++ {
			if !comparator(super[i-j], sub[len(sub)-1-j]) {
				continue OUTER
			}
		}
		return i - len(sub) + 1, nil
	}

	return 0, errors.New("Could not find element")
}

func FindFirstOfSlice[T comparable](haystack []T, needles []T, comparator func(T, T) bool) (int, error) {
	if len(haystack) == 0 || len(needles) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
			if comparator(haystackValue, needleValue) {
				return i, nil
			}
		}
	}

	return 0, errors.New("Could not find element")
}

func FindFirstOfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, needles []TValue, comparator func(TValue, TValue) bool) (TKey, error) {
	var zeroVal TKey
	if len(haystack) == 0 || len(needles) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
			if comparator(haystackValue, needleValue) {
				return i, nil
			}
		}
	}

	return zeroVal, errors.New("Could not find element")
}

func AllOfSlice[T comparable](haystack []T, comparator func(T) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range haystack {
		if !comparator(value) {
			return ComparisonError[int]{i}
		}
	}

	return nil
}

func AllOfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range haystack {
		if !comparator(value) {
			return ComparisonError[TKey]{key}
		}
	}

	return nil
}

func AnyOfSlice[T comparable](haystack []T, comparator func(T) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range haystack {
		if comparator(value) {
			return nil
		}
	}

	return ComparisonError[int]{len(haystack) - 1}
}

func AnyOfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range haystack {
		if comparator(value) {
			return nil
		}
	}

	return ComparisonError[TKey]{}
}

func NoneOfSlice[T comparable](haystack []T, comparator func(T) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range haystack {
		if comparator(value) {
			return ComparisonError[int]{len(haystack) - 1}
		}
	}

	return nil
}

func NoneOfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range haystack {
		if comparator(value) {
			return ComparisonError[TKey]{}
		}
	}

	return nil
}

func ForeachSlice[T comparable](haystack []T, comparator func(T) error) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range haystack {
		err := comparator(value)
		if err != nil {
			return ExecutionError[int]{i, err}
		}
	}

	return nil
}

func ForeachMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) error) error {
	if len(haystack) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range haystack {
		err := comparator(value)
		if err != nil {
			return ExecutionError[TKey]{key, err}
		}
	}

	return nil
}

func CountSlice[T comparable](haystack []T, wanted T) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range haystack {
		if value == wanted {
			counter++
		}
	}

	return counter, nil
}

func CountMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, wanted TValue) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range haystack {
		if value == wanted {
			counter++
		}
	}

	return counter, nil
}

func CountIfSlice[T comparable](haystack []T, comparator func(T) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range haystack {
		if comparator(value) {
			counter++
		}
	}

	return counter, nil
}

func CountIfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, comparator func(TValue) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range haystack {
		if comparator(value) {
			counter++
		}
	}

	return counter, nil
}

func MismatchSlice[T comparable](iterable1 []T, iterable2 []T) (int, error) {
	if len(iterable1) == 0 || len(iterable2) == 0 {
		return 0, EmptyIterableError{}
	}

	i := 0
	for ; i < min(len(iterable1), len(iterable2)); i++ {
		if iterable1[i] != iterable2[i] {
			return i, nil
		}
	}

	return 0, EqualIteratorsError{}
}

func AdjacentFindSlice[T comparable](haystack []T, comparator func(T, T) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i := 0; i < len(haystack)-1; i++ {
		if comparator(haystack[i], haystack[i+1]) {
			return i, nil
		}
	}

	return 0, errors.New("Could not find element")
}

func CopyReplaceSlice[T comparable](original []T, toReplace T, replacement T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if value == toReplace {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

func CopyReplaceMap[TKey comparable, TValue comparable](original map[TKey]TValue, toReplace TValue, replacement TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if value == toReplace {
			newContainer[key] = replacement
		} else {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

func CopyReplaceIfSlice[T comparable](original []T, comparator func(T) bool, replacement T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if comparator(value) {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

func CopyReplaceIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, comparator func(TValue) bool, replacement TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if comparator(value) {
			newContainer[key] = replacement
		} else {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

func CopyReplaceIfNotSlice[T comparable](original []T, comparator func(T) bool, replacement T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if !comparator(value) {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

func CopyReplaceIfNotMap[TKey comparable, TValue comparable](original map[TKey]TValue, comparator func(TValue) bool, replacement TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if !comparator(value) {
			newContainer[key] = replacement
		} else {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

func CopyExceptSlice[T comparable](original []T, toExclude T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if value != toExclude {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

func CopyExceptMap[TKey comparable, TValue comparable](original map[TKey]TValue, toExclude TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if value != toExclude {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

func CopyExceptIfSlice[T comparable](original []T, comparator func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if !comparator(value) {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

func CopyExceptIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, comparator func(TValue) bool) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if !comparator(value) {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

func CopyExceptIfNotSlice[T comparable](original []T, comparator func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if comparator(value) {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

func CopyExceptIfNotMap[TKey comparable, TValue comparable](original map[TKey]TValue, comparator func(TValue) bool) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if comparator(value) {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

func TransformMap[TKey comparable, TValue comparable](container map[TKey]TValue, transformer func(TValue) (TValue, error)) error {

	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key := range container {
		newValue, err := transformer(container[key])
		if err != nil {
			return ExecutionError[TValue]{container[key], err}
		}

		container[key] = newValue
	}

	return nil
}

func TransformSlice[T comparable](container []T, transformer func(*T) error) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		err := transformer(&container[i])
		if err != nil {
			return ExecutionError[T]{value, err}
		}
	}

	return nil
}

func TransformMapUnsafe[TKey comparable, TValue comparable](container map[TKey]TValue, transformer func(TValue) TValue) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key := range container {
		container[key] = transformer(container[key])
	}

	return nil
}

func TransformSliceUnsafe[T comparable](container []T, transformer func(*T)) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i := range container {
		transformer(&container[i])
	}

	return nil
}
