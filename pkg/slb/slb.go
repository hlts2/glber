package slb

import (
	"net/http"

	"github.com/pkg/errors"
)

// Server --
type Server interface {
	Serve() error
	ServeTLS(certFile, keyFile string) error
	Shutdown()
}

// serverLoadBalancer --
type serverLoadBalancer struct {
	*Config
	Server *http.Server
}

// New --
func New(cfg *Config) Server {
	return &serverLoadBalancer{
		Config: cfg,
		Server: nil,
	}
}

func (s *serverLoadBalancer) Serve() error {
	lis, err := s.LoadBalancer.createListener()
	if err != nil {
		return errors.Wrap(err, "faild to create listener")
	}

	err = s.Server.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "faild to serve")
	}
	return nil
}

func (s *serverLoadBalancer) ServeTLS(certFile, keyFile string) error {
	lis, err := s.LoadBalancer.createListener()
	if err != nil {
		return errors.Wrap(err, "faild to create listener")
	}

	err = s.Server.ServeTLS(lis, certFile, keyFile)
	if err != nil {
		return errors.Wrap(err, "faild to serve with TLS")
	}

	return nil
}

func (s *serverLoadBalancer) Shutdown() {}
