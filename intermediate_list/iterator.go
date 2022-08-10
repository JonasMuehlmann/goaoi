package intermediatelist

import "github.com/JonasMuehlmann/datastructures.go/lists/arraylist"

type Iterator[TValue any] struct {
	*arraylist.Iterator[TValue]
}
