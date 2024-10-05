package pow

import (
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/pow/hashcash"
)

// An interface choice was inspired by https://en.wikipedia.org/wiki/File:Proof_of_Work_solution_verification.svg
type POW interface {
	Compute() ([]byte, error)
	Solve(challenge []byte) ([]byte, error)
	Verify(challenge, solution []byte) error
}

var _ POW = &hashcash.HashCash{}
