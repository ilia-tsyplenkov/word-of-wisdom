package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ilia-tsyplenkov/word-of-wisdom/config"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/pow/hashcash"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/repository/file"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/server"
	log "github.com/sirupsen/logrus"
)

const repoFile = "quotes.txt"

func main() {

	cfg, err := config.NewServerConfig()
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
		log.Fatalf("failed to initialize hashcash pow: %v", err)
	}
	repo, err := file.NewRepo(repoFile)
	if err != nil {
		log.Fatalf("failed to initialize repository: %v", err)
	}

	srv := server.New(cfg, pow, repo)

	signalCh := make(chan os.Signal, 1)

	go signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalCh
		srv.Stop()
	}()

	if err := srv.Run(); err != nil {
		log.Errorf("server: %v", err)
	}
}
