package file

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type fileRepo struct {
	quotes []string
}

func NewRepo(fname string) (*fileRepo, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("failed to open source file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	quotes := make([]string, 0)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading source file")
	}
	return &fileRepo{
		quotes: quotes,
	}, nil
}

func (r *fileRepo) GetRecord() (string, error) {

	return r.quotes[rand.Intn(len(r.quotes))], nil
}
