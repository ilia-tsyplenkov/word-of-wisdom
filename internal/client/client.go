package client

import (
	"context"
	"fmt"
	"net"

	"github.com/ilia-tsyplenkov/word-of-wisdom/config"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/pow"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/utils"
)

type client struct {
	pow    pow.POW
	config *config.ClientConfig
}

func New(cfg *config.ClientConfig, pow pow.POW) *client {
	return &client{
		pow:    pow,
		config: cfg,
	}
}

func (c *client) Start(ctx context.Context, iterations int) error {
	for i := 0; i < iterations; i++ {

	}
	return nil
}

func (c *client) getQuote() ([]byte, error) {
	conn, err := net.Dial("tcp", c.config.ServerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed connect to the server: %v", err)
	}

	defer conn.Close()

	if err = utils.SendMessage(conn, []byte("get challenge")); err != nil {
		return nil, fmt.Errorf("error to get challenge: %v", err)

	}
	challenge, err := utils.ReceiveMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("error to receive challenge: %v", err)
	}

	solution, err := c.pow.Solve(challenge)
	if err != nil {
		return nil, fmt.Errorf("error to solve challenge: %v", err)
	}
	if err = utils.SendMessage(conn, solution); err != nil {
		return nil, fmt.Errorf("error to send solution: %v", err)
	}

	quote, err := utils.ReceiveMessage(conn)
	if err != nil {
		return nil, fmt.Errorf("error to receive quote: %v", err)
	}

	return quote, nil
}
