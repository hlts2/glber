package slb

import (
	"net"
	"net/http"

	"github.com/pkg/errors"
)

// Server --
type Server interface {
	Serve(l net.Listener) error
	ServeTLS(l net.Listener, certFile, keyFile string) error
	Shutdown()
}

// serverLoadBalancer --
type serverLoadBalancer struct {
	server *http.Server
}

// New --
func New(cfg Config) Server {
	return new(serverLoadBalancer)
}

func (s *serverLoadBalancer) Serve(l net.Listener) error {
	err := s.server.Serve(l)
	if err != nil {
		return errors.Wrap(err, "faild to serve")
	}
	return nil
}

func (s *serverLoadBalancer) ServeTLS(l net.Listener, certFile, keyFile string) error {
	err := s.server.ServeTLS(l, certFile, keyFile)
	if err != nil {
		return errors.Wrap(err, "faild to serve with TLS")
	}

	return nil
}

func (s *serverLoadBalancer) Shutdown() {}
