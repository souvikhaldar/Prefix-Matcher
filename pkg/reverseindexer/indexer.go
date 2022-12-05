package reverseindexer

type ReverseIndexer interface {
	ReverseIndex([]string) map[rune][]string
}
