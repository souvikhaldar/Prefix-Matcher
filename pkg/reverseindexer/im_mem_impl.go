package reverseindexer

var maxLen int

func GetMaxIndexLen() int {
	return maxLen
}

type InMemIndexer struct{}

func NewInMemIndexer() *InMemIndexer {
	return &InMemIndexer{}
}

func (i *InMemIndexer) ReverseIndex(input []string) map[string][]string {
	if len(input) == 0 {
		return nil
	}

	// find out the length of largest element in input
	for _, el := range input {
		if len(el) > maxLen {
			maxLen = len(el)
		}
	}

	// find out all the unique characters possible at a given position
	// eg. all the unique characters present throughout the input
	// at position 0 maybe [a,5,G,x,Y]
	uniqueCharPos := make([][]string, maxLen)

	for i := 0; i < maxLen; i++ {
		unique := make(map[string]struct{})
		for _, el := range input {
			if i >= len(el) {
				continue
			}
			unique[string(el[:i+1])] = struct{}{}
		}
		for unEl := range unique {
			uniqueCharPos[i] = append(uniqueCharPos[i], string(unEl))
		}
	}

	subStrMap := make(map[string][]string)
	for _, in := range input {
		for _, unEl := range uniqueCharPos {
			for _, un := range unEl {
				if len(in) <= len(un) {
					continue
				}
				if in[:len(un)] == un {
					subStrMap[un] = append(subStrMap[un], in)
				}
			}
		}
	}

	return subStrMap
}
