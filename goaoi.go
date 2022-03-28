package goaoi

func FindSlice[T comparable](haystack []T, needle T) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, value := range haystack {
		if value == needle {
			return i, nil
		}
	}

	return 0, ElementNotFoundError{}
}

func FindIfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, unary_predicate func(TValue) bool) (TKey, error) {
	var zeroVal TKey

	if len(haystack) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	for key, value := range haystack {
		if unary_predicate(value) {
			return key, nil
		}
	}

	return zeroVal, ElementNotFoundError{}
}

func FindIfSlice[T comparable](haystack []T, unary_predicate func(T) bool) (int, error) {
	if len(haystack) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, value := range haystack {
		if unary_predicate(value) {
			return i, nil
		}
	}

	return 0, ElementNotFoundError{}
}

func FindEndSlice[T comparable](super []T, sub []T, binary_predicate func(T, T) bool) (int, error) {
	if len(super) == 0 || len(sub) == 0 {
		return 0, EmptyIterableError{}
	}
OUTER:
	for i := len(super) - 1; i >= len(sub)-1; i-- {
		for j := 0; j < len(sub); j++ {
			if !binary_predicate(super[i-j], sub[len(sub)-1-j]) {
				continue OUTER
			}
		}
		return i - len(sub) + 1, nil
	}

	return 0, ElementNotFoundError{}
}

func FindFirstOfSlice[T comparable](haystack []T, needles []T, binary_predicate func(T, T) bool) (int, error) {
	if len(haystack) == 0 || len(needles) == 0 {
		return 0, EmptyIterableError{}
	}

	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
			if binary_predicate(haystackValue, needleValue) {
				return i, nil
			}
		}
	}

	return 0, ElementNotFoundError{}
}

func FindFirstOfMap[TKey comparable, TValue comparable](haystack map[TKey]TValue, needles []TValue, binary_predicate func(TValue, TValue) bool) (TKey, error) {
	var zeroVal TKey
	if len(haystack) == 0 || len(needles) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	for i, haystackValue := range haystack {
		for _, needleValue := range needles {
			if binary_predicate(haystackValue, needleValue) {
				return i, nil
			}
		}
	}

	return zeroVal, ElementNotFoundError{}
}

func AllOfSlice[T comparable](container []T, unary_predicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		if !unary_predicate(value) {
			return ComparisonError[int]{i}
		}
	}

	return nil
}

func AllOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range container {
		if !unary_predicate(value) {
			return ComparisonError[TKey]{key}
		}
	}

	return nil
}

func AnyOfSlice[T comparable](container []T, unary_predicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unary_predicate(value) {
			return nil
		}
	}

	return ComparisonError[int]{len(container) - 1}
}

func AnyOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unary_predicate(value) {
			return nil
		}
	}

	return ComparisonError[TKey]{}
}

func NoneOfSlice[T comparable](container []T, unary_predicate func(T) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unary_predicate(value) {
			return ComparisonError[int]{len(container) - 1}
		}
	}

	return nil
}

func NoneOfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		if unary_predicate(value) {
			return ComparisonError[TKey]{}
		}
	}

	return nil
}

func ForeachSlice[T comparable](container []T, unary_func func(T) error) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for i, value := range container {
		err := unary_func(value)
		if err != nil {
			return ExecutionError[int]{i, err}
		}
	}

	return nil
}

func ForeachMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_func func(TValue) error) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for key, value := range container {
		err := unary_func(value)
		if err != nil {
			return ExecutionError[TKey]{key, err}
		}
	}

	return nil
}

func ForeachSliceUnsafe[T comparable](container []T, unary_func func(T)) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		unary_func(value)
	}

	return nil
}

func ForeachMapUnsafe[TKey comparable, TValue comparable](container map[TKey]TValue, unary_func func(TValue)) error {
	if len(container) == 0 {
		return EmptyIterableError{}
	}

	for _, value := range container {
		unary_func(value)
	}

	return nil
}

func CountSlice[T comparable](container []T, wanted T) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if value == wanted {
			counter++
		}
	}

	return counter, nil
}

func CountMap[TKey comparable, TValue comparable](container map[TKey]TValue, wanted TValue) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if value == wanted {
			counter++
		}
	}

	return counter, nil
}

func CountIfSlice[T comparable](container []T, unary_predicate func(T) bool) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if unary_predicate(value) {
			counter++
		}
	}

	return counter, nil
}

func CountIfMap[TKey comparable, TValue comparable](container map[TKey]TValue, unary_predicate func(TValue) bool) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	counter := 0
	for _, value := range container {
		if unary_predicate(value) {
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

func AdjacentFindSlice[T comparable](container []T, binary_predicate func(T, T) bool) (int, error) {
	if len(container) == 0 {
		return 0, EmptyIterableError{}
	}

	for i := 0; i < len(container)-1; i++ {
		if binary_predicate(container[i], container[i+1]) {
			return i, nil
		}
	}

	return 0, ElementNotFoundError{}
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

func CopyReplaceIfSlice[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if unary_predicate(value) {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

func CopyReplaceIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool, replacement TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if unary_predicate(value) {
			newContainer[key] = replacement
		} else {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

func CopyReplaceIfNotSlice[T comparable](original []T, unary_predicate func(T) bool, replacement T) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if !unary_predicate(value) {
			newContainer = append(newContainer, replacement)
		} else {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

func CopyReplaceIfNotMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool, replacement TValue) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if !unary_predicate(value) {
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

func CopyExceptIfSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if !unary_predicate(value) {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

func CopyExceptIfMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if !unary_predicate(value) {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

func CopyExceptIfNotSlice[T comparable](original []T, unary_predicate func(T) bool) ([]T, error) {
	var zeroVal []T

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make([]T, 0, len(original))

	for _, value := range original {
		if unary_predicate(value) {
			newContainer = append(newContainer, value)
		}
	}

	return newContainer, nil
}

func CopyExceptIfNotMap[TKey comparable, TValue comparable](original map[TKey]TValue, unary_predicate func(TValue) bool) (map[TKey]TValue, error) {
	var zeroVal map[TKey]TValue

	if len(original) == 0 {
		return zeroVal, EmptyIterableError{}
	}

	newContainer := make(map[TKey]TValue, len(original))

	for key, value := range original {
		if unary_predicate(value) {
			newContainer[key] = value
		}
	}

	return newContainer, nil
}

func FillSlice[T comparable](arr *[]T, filler T) []T {
	for i := range *arr {
		(*arr)[i] = filler
	}

	n_unfilled := cap(*arr) - len(*arr)

	for i := 0; i < n_unfilled; i++ {
		*arr = append(*arr, filler)
	}

	return *arr
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
