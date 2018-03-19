package servers

import (
	"context"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
)

// HTTPServer .
type HTTPServer struct {
	httpServer *http.Server
	errChan    chan error
	doneChan   chan error
}

// NewHTTPServer .
func NewHTTPServer(listenAddress string, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		httpServer: &http.Server{
			Addr:    listenAddress,
			Handler: handler,
		},
		doneChan: make(chan error, 1),
	}
}

// Start .
func (s *HTTPServer) Start() {
	logrus.Infof("Starting HTTP server at %s", s.httpServer.Addr)
	err := s.httpServer.ListenAndServe()
	if err != nil && s.errChan != nil {
		s.errChan <- err
	}
}

// ErrChan .
func (s *HTTPServer) ErrChan(errChan chan error) {
	s.errChan = errChan
}

// DoneChan .
func (s *HTTPServer) DoneChan() chan error {
	return s.doneChan
}

// Stop .
func (s *HTTPServer) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	logrus.Infof("Shutting down HTTP Server at %s", s.httpServer.Addr)
	s.doneChan <- s.httpServer.Shutdown(ctx)
}
