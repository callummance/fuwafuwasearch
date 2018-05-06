package main

import (
	"fmt"

	"github.com/callummance/fuwafuwasearch/levenshteinmatrix"
)

func main() {
	fmt.Printf("%v", levenshteinmatrix.ComputeMatchVal("sitting", "kitten", 1))
}
