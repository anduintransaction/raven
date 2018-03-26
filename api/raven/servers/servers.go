package servers

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

// Server .
type Server interface {
	Start()
	DoneChan() chan error
	ErrChan(errChan chan error)
	Stop()
}

// Servers .
type Servers struct {
	servers    map[string]Server
	errChan    chan error
	signalChan chan os.Signal
}

// NewServers .
func NewServers() *Servers {
	s := &Servers{
		servers:    make(map[string]Server),
		errChan:    make(chan error, 1),
		signalChan: make(chan os.Signal, 1),
	}
	return s
}

// AddServer .
func (s *Servers) AddServer(name string, server Server) *Servers {
	server.ErrChan(s.errChan)
	s.servers[name] = server
	return s
}

// ListenAndServe .
func (s *Servers) ListenAndServe() {
	s.start()
	s.wait()
}

// Start .
func (s *Servers) start() {
	logrus.Info("Starting servers")
	signal.Notify(s.signalChan, syscall.SIGINT, syscall.SIGTERM)
	for _, server := range s.servers {
		go server.Start()
	}
}

// Wait .
func (s *Servers) wait() {
	select {
	case sig := <-s.signalChan:
		logrus.Debugf("Got signal %v", sig)
		logrus.Info("Stopping servers")
		for name, server := range s.servers {
			server.Stop()
			timer := time.NewTimer(30 * time.Second)
			select {
			case err := <-server.DoneChan():
				if err != nil {
					logrus.Error(err)
				}
			case <-timer.C:
				logrus.Errorf("Timeout waiting for server %q", name)
			}
		}
	case err := <-s.errChan:
		logrus.Error(err)
		os.Exit(1)
	}
}
