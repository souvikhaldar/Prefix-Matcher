package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/souvikhaldar/prefix-matcher/pkg/reader"
	"github.com/souvikhaldar/prefix-matcher/pkg/reverseindexer"
)

func main() {
	pref := flag.String("file", "sp.txt", "path to the file containing prefixes")
	flag.Parse()
	log.Println("Indexing the prefix list...")
	t := time.Now()
	read := reader.NewFileReader(*pref)
	prefixes, err := read.Read()
	if err != nil {
		log.Fatal(err)
	}
	indexer := reverseindexer.NewInMemIndexer()
	index := indexer.ReverseIndex(prefixes)
	log.Printf("Indexing completed in %f seconds.", time.Since(t).Seconds())

	for {
		var in string
		fmt.Println("Enter string to search: ")
		fmt.Scanln(&in)
		// start checking
		var longestPrefixes []string
		for i := 0; i < len(in); i++ {
			p := index[in[:i+1]]
			if len(p) == 0 {
				break
			}
			longestPrefixes = p
		}
		PrintPrefix(longestPrefixes)

		// check if user wants to continue
		fmt.Println("Continue to search? (y/n)")
		fmt.Scanln(&in)
		if strings.Contains(in, "n") || strings.Contains(in, "N") {
			break
		}
	}
}

func PrintPrefix(p []string) {
	max := 0
	var maxEl string
	for _, el := range p {
		if len(el) > max {
			max = len(el)
			maxEl = el
		}
	}
	log.Println("Longest matching prefix: ", maxEl)
}
