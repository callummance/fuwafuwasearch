package main

import (
	"fmt"

	"github.com/callummance/fuwafuwasearch/levenshteinmatrix"
)

func main() {
	searcher := levenshteinmatrix.NewLMatrixSearch([]string{"test1", "blargle", "plop", "sitting", "sittings"})
	fmt.Printf("%v", searcher.SearchForSubstring("sitting"))
}
