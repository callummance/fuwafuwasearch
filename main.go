package main

import (
	"fmt"

	"github.com/callummance/fuwafuwasearch/levenshteinmatrix"
)

func main() {
	searcher := levenshteinmatrix.NewLMatrixSearch([]string{"test1", "blargle", "plop", "sitting", "sittings"}, []string{"a", "b", "c", "d", "e"}, false)
	fmt.Printf("%v", searcher.GetMatchingKeys("sitting", 5))
}
