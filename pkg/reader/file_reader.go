package reader

import (
	"bufio"
	"os"
)

type FileReader struct {
	fileName string
}

func (f *FileReader) Read() ([]string, error) {
	file, err := os.Open(f.fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	prefixes := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prefixes = append(prefixes, scanner.Text())
	}
	err = scanner.Err()
	return prefixes, err
}
