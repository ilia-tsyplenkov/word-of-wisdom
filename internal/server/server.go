package server

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/ilia-tsyplenkov/word-of-wisdom/config"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/pow"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/repository"
	"github.com/ilia-tsyplenkov/word-of-wisdom/internal/utils"
	log "github.com/sirupsen/logrus"
)

type server struct {
	pow      pow.POW
	config   *config.ServerConfig
	listener net.Listener
	wg       sync.WaitGroup
	cancel   context.CancelFunc
	limiter  chan struct{}

	repo repository.Repository
}

func New(cfg *config.ServerConfig, pow pow.POW, repo repository.Repository) *server {
	return &server{
		pow:     pow,
		config:  cfg,
		limiter: make(chan struct{}, cfg.RequestLimiter),
		repo:    repo,
	}
}

func (s *server) Run() error {
	var (
		err error
		ctx context.Context
	)
	ctx, s.cancel = context.WithCancel(context.Background())
	defer s.cancel()

	s.listener, err = net.Listen("tcp", s.config.Addr)
	if err != nil {
		return fmt.Errorf("failed to start server on %s: %v", s.config.Addr, err)
	}

	log.Infof("server started on: %s", s.config.Addr)
	s.serve(ctx)
	log.Info("server stopped")
	return nil
}

func (s *server) Stop() {
	log.Infof("stopping server")
	s.cancel()
}

func (s *server) serve(ctx context.Context) {

	defer s.wg.Wait()

	for {
		select {
		case <-ctx.Done():
			s.listener.Close()
			return
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				log.Errorf("error to accept connection: %v", err)
				break
			}
			s.wg.Add(1)
			// limit number of simultaneously handled connection
			s.limiter <- struct{}{}
			go func(conn net.Conn) {
				defer func() {
					<-s.limiter
					s.wg.Done()
				}()
				if err := s.handle(conn); err != nil {
					log.Errorf("error to handle connection: %v", err)
				}
			}(conn)

		}
	}
}

func (s *server) handle(conn net.Conn) error {
	defer conn.Close()

	// receive challenge request
	if _, err := utils.ReceiveMessage(conn); err != nil {
		return fmt.Errorf("error receiving challenge request: %v", err)
	}

	challenge, err := s.pow.Compute()
	if err != nil {
		return fmt.Errorf("error computing challenge: %v", err)
	}

	// send challenge
	if err = utils.SendMessage(conn, challenge); err != nil {
		return fmt.Errorf("error sending challenge: %v", err)
	}

	// receive solution
	solution, err := utils.ReceiveMessage(conn)
	if err != nil {
		return fmt.Errorf("error receiving solution")
	}

	// check the solution
	if err = s.pow.Verify(challenge, solution); err != nil {
		return fmt.Errorf("failed to verify solution")
	}

	// get quote
	quote, err := s.repo.GetRecord()
	if err != nil {
		return fmt.Errorf("error to get quote from repository: %v", err)
	}

	// send quote
	if err := utils.SendMessage(conn, []byte(quote)); err != nil {
		return fmt.Errorf("error to send quote: %v", err)
	}

	return nil
}
