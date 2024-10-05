package file

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const testQuotesFile = "quotes.txt"

func TestQuotesNumber(t *testing.T) {

	expected := readQuotes(t)
	repo, err := NewRepo(testQuotesFile)
	require.NoError(t, err)
	require.Equal(t, len(expected), len(repo.quotes))
}

func TestGetRecord(t *testing.T) {
	repo, err := NewRepo(testQuotesFile)
	require.NoError(t, err)

	rec, err := repo.GetRecord()
	require.NoError(t, err)
	require.NotEqual(t, "", rec)
}

func readQuotes(t *testing.T) []string {
	content, err := os.ReadFile(testQuotesFile)
	require.NoError(t, err)
	res := strings.Split(string(content), "\n")
	// some text editors may add additional '\n' to the file end
	if res[len(res)-1] == "" {
		res = res[:len(res)-1]
	}
	return res
}
