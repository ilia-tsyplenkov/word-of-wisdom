package client

import (
	"fmt"
	"net"

	"github.com/ilia-tsyplenkov/word-of-wisdom/config"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/pow"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/utils"
	log "github.com/sirupsen/logrus"
)

type internalClient struct {
	pow    pow.POW
	config *config.ClientConfig
}

func New(cfg *config.ClientConfig, pow pow.POW) *internalClient {
	return &internalClient{
		pow:    pow,
		config: cfg,
	}
}

func (c *internalClient) Run(iterations int) {
	for i := 0; i < iterations; i++ {
		quote, err := c.getQuote()
		if err != nil {
			log.Errorf("error to get quote: %v", err)
		} else {
			log.Infof("quote: %s", string(quote))
		}
	}
}

func (c *internalClient) getQuote() ([]byte, error) {
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
