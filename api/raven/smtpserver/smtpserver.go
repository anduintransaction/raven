package smtpserver

import (
	"github.com/anduintransaction/raven/api/raven/config"
	"github.com/anduintransaction/raven/api/raven/servers"
	guerrilla "github.com/flashmob/go-guerrilla"
	"github.com/flashmob/go-guerrilla/backends"
	"github.com/flashmob/go-guerrilla/log"
)

// SMTPServer .
type SMTPServer struct {
	daemon   *guerrilla.Daemon
	config   *config.SMTPServerConfig
	doneChan chan error
	errChan  chan error
}

// NewSMTPServer .
func NewSMTPServer(config *config.SMTPServerConfig) servers.Server {
	return &SMTPServer{
		config:   config,
		doneChan: make(chan error, 1),
	}
}

// Start .
func (s *SMTPServer) Start() {
	sc := guerrilla.ServerConfig{
		ListenInterface: s.config.ListenAddress,
		IsEnabled:       true,
	}
	cfg := &guerrilla.AppConfig{
		AllowedHosts: []string{"."},
		LogFile:      log.OutputStdout.String(),
		Servers:      []guerrilla.ServerConfig{sc},
		BackendConfig: backends.BackendConfig{
			"save_process":      "Hasher|Postgres",
			"save_workers_size": 3,
		},
	}
	s.daemon = &guerrilla.Daemon{
		Config: cfg,
	}
	s.daemon.AddProcessor("Postgres", Postgres)
	err := s.daemon.Start()
	if err != nil && s.errChan != nil {
		s.errChan <- err
	}
}

// DoneChan .
func (s *SMTPServer) DoneChan() chan error {
	return s.doneChan
}

// ErrChan .
func (s *SMTPServer) ErrChan(errChan chan error) {
	s.errChan = errChan
}

// Stop .
func (s *SMTPServer) Stop() {
	s.daemon.Shutdown()
	s.doneChan <- nil
}
