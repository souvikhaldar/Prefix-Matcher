package reverseindexer

import (
	"log"
)

type ReverseCharMap map[rune][]string

type InMemIndexer struct{}

func NewInMemIndexer() *InMemIndexer {
	return &InMemIndexer{}
}

func (i *InMemIndexer) ReverseIndex(input []string) []ReverseCharMap {
	if len(input) == 0 {
		return nil
	}

	// find out the length of largest element in input
	var maxLen int
	for _, el := range input {
		if len(el) > maxLen {
			maxLen = len(el)
		}
	}
	log.Println("Maxlen: ", maxLen)
	// all the unique characters present at each position of the input list
	// for eg, all the unique characters present in 0th position of the all
	// the inputs
	posUniqueChars := make([][]rune, maxLen)

	for i := 0; i < maxLen; i++ {
		unique := make(map[rune]struct{})
		for _, el := range input {
			if i >= len(el) {
				continue
			}
			k := rune(el[i])
			if _, ok := unique[k]; !ok {
				unique[k] = struct{}{}
			}
		}
		for key := range unique {
			posUniqueChars[i] = append(posUniqueChars[i], key)
		}
	}

	// create the reversemap slice of length maxLen
	reverseMap := make([]ReverseCharMap, maxLen)
	for pos, puc := range posUniqueChars {
		var rcm ReverseCharMap = make(map[rune][]string)
		for _, p := range puc {
			for _, in := range input {
				if pos >= len(in) {
					continue
				}
				k := rune(in[pos])
				if k == p {
					rcm[p] = append(rcm[p], in)
				}
			}
		}
		reverseMap[pos] = rcm
	}
	return reverseMap
}
