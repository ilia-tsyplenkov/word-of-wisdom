package hashcash_test

import (
	"encoding/hex"
	"testing"

	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/pow/hashcash"
	"github.com/stretchr/testify/require"
)

func TestNewOK(t *testing.T) {
	_, err := hashcash.New(uint8(10))
	require.NoError(t, err)
}

func TestNewFail(t *testing.T) {
	for _, c := range []uint8{0, 26} {
		_, err := hashcash.New(c)
		require.Error(t, err)
	}
}

func TestCompute(t *testing.T) {
	complexity := uint8(11)
	hc, _ := hashcash.New(complexity)
	token, err := hc.Compute()
	require.NoError(t, err)
	require.Equal(t, complexity, uint8(len(token)))
}

func TestSolve(t *testing.T) {
	challenge := []byte("hi there!")
	hc, _ := hashcash.New(uint8(len(challenge)))
	solution, err := hc.Solve(challenge)
	require.NoError(t, err)
	require.Equal(t, "1f01000000000000", hex.EncodeToString(solution))
}

func TestVerify(t *testing.T) {

	for _, tc := range []struct {
		title     string
		challenge []byte
		solution  string
		success   bool
	}{
		{
			title:     "pass",
			challenge: []byte("hi there!"),
			solution:  "1f01000000000000",
			success:   true,
		},

		{
			title:     "fail",
			challenge: []byte("it must fail"),
			solution:  "1f01000000000000",
			success:   false,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			hc, _ := hashcash.New(uint8(len(tc.challenge)))
			solution, err := hex.DecodeString(tc.solution)
			require.NoError(t, err)
			err = hc.Verify(tc.challenge, solution)
			require.Equal(t, tc.success, err == nil)
		})
	}

}
