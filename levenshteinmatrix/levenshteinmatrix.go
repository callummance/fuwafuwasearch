package levenshteinmatrix

import (
	"strings"
)

//LMatrixSearch contains the search target and and data required for a levenshtein
//matrix search
type LMatrixSearch struct {
	SearchLibrary []string
	Keys          []interface{}
	CaseSensitive bool
}

//NewLMatrixSearch returns a new LMatrixSerachStruct
func NewLMatrixSearch(library []string, keys []interface{}, caseSensitive bool) *LMatrixSearch {
	s := new(LMatrixSearch)
	for _, searchString := range library {
		if caseSensitive {
			s.SearchLibrary = append(s.SearchLibrary, searchString)
		} else {
			s.SearchLibrary = append(s.SearchLibrary, strings.ToLower(searchString))
		}
	}

	s.Keys = keys
	s.CaseSensitive = caseSensitive
	return s
}

//GetMatchingKeys returns a map of keys which are sufficiently close to the
//searchTerm paired with their distance
func (s *LMatrixSearch) GetMatchingKeys(searchTerm string, maxDiff int) map[interface{}]int {
	distances := s.SearchForSubstring(searchTerm)
	res := make(map[interface{}]int)
	shouldUseKeys := len(s.SearchLibrary) == len(s.Keys)
	for i, d := range distances {
		if shouldUseKeys {
			if d <= maxDiff {
				res[s.Keys[i]] = d
			}
		} else {
			if d <= maxDiff {
				res[s.SearchLibrary[i]] = d
			}
		}
	}
	return res
}

//SearchForSubstring returns a slice containing match distance for each of the
//items in the SearchLibrary
func (s *LMatrixSearch) SearchForSubstring(searchTerm string) []int {
	casedSearchTerm := func() string {
		if !s.CaseSensitive {
			return strings.ToLower(searchTerm)
		}
		return searchTerm
	}()
	res := make([]int, len(s.SearchLibrary), len(s.SearchLibrary))
	matches := make(chan ([]int))

	for i, searchTarget := range s.SearchLibrary {
		potMatch := searchTarget
		index := i
		go func() {
			distance := computeMatchVal(casedSearchTerm, potMatch)
			matches <- []int{index, distance}
		}()
	}
	for i := 0; i < len(s.SearchLibrary); i++ {
		match := <-matches
		index := match[0]
		distance := match[1]
		res[index] = distance
	}
	return res
}

func computeMatchVal(needle string, haystack string) int {
	needleSize := len(needle)
	haystackSize := len(haystack)
	minDistance := needleSize
	d := make([]int, (needleSize+1)*(haystackSize+1), (needleSize+1)*(haystackSize+1))

	for j := 0; j <= (needleSize); j++ {
		d[j] = j
	}

	for i := 0; i <= (haystackSize); i++ {
		d[i*(needleSize+1)] = 0
	}

	for i := 0; i < haystackSize; i++ {
		for j := 0; j < needleSize; j++ {
			if needle[j] == haystack[i] {
				d[(i+1)*(needleSize+1)+j+1] = d[i*(needleSize+1)+j]
			} else {
				deletion := d[i*(needleSize+1)+(j+1)] + 1
				insertion := d[(i+1)*(needleSize+1)+j] + 1
				substitution := d[i*(needleSize+1)+j] + 1
				d[(i+1)*(needleSize+1)+j+1] = min3(deletion, insertion, substitution)
			}
		}
		minDistance = min(minDistance, d[(i+1)*(needleSize+1)+needleSize])
	}
	return minDistance
}

func min3(x, y, z int) int {
	if x < y {
		return min(x, z)
	}
	return min(y, z)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
