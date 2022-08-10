package intermediatelist

import (
	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
)

type List[TValue any] struct {
	*arraylist.List[TValue]
}
