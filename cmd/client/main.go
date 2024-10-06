package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/ilia-tsyplenkov/word-of-wisdom/config"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/client"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/pow/hashcash"
)

func main() {

	cfg, err := config.NewClientConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Errorf("failed to set log level: %v", err)
	}

	log.SetLevel(logLevel)

	pow, err := hashcash.New(cfg.Complexity)
	if err != nil {
		log.Fatalf("failed to initialize hashcash pow")
	}

	cl := client.New(cfg, pow)
	cl.Run(cfg.Iterations)
}
