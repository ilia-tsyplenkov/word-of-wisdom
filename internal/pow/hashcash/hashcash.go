package hashcash

import (
	"crypto/rand"
	"fmt"

	hcash "github.com/repbin/repbin/hashcash"
)

const (
	complexityLimit = 25
)

type HashCash struct {
	complexity byte
}

func New(complexity uint8) (*HashCash, error) {
	if complexity == 0 || complexity > complexityLimit {
		return nil, fmt.Errorf("invalid complexity level")
	}

	return &HashCash{
		complexity: byte(complexity),
	}, nil
}

func (h *HashCash) Compute() ([]byte, error) {
	token := make([]byte, h.complexity)
	if _, err := rand.Read(token); err != nil {
		return nil, fmt.Errorf("failed to create challenge: %v", err)
	}

	return token, nil
}

func (h *HashCash) Solve(challenge []byte) ([]byte, error) {

	nonce, ok := hcash.ComputeNonce(challenge, h.complexity, 0, 0)
	if !ok {
		return nil, fmt.Errorf("failed to solve the challenge")
	}

	return nonce, nil

}

func (h *HashCash) Verify(challenge, solution []byte) error {
	ok, _ := hcash.TestNonce(challenge, solution, h.complexity)

	if !ok {
		return fmt.Errorf("failed to verify the challenge")
	}
	return nil
}
