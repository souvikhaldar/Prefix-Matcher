package reverseindexer

import (
	"testing"

	"github.com/souvikhaldar/prefix-matcher/pkg/reader"
)

func TestReverseIndex(t *testing.T) {
	fr := reader.NewFileReader("firstNames.txt")
	in, _ := fr.Read()
	ri := NewInMemIndexer()
	rim := ri.ReverseIndex(in)
	if len(rim["m"]) == 0 {
		t.Fatal("Should not be empty")
	}
}
