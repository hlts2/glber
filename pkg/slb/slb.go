package slb

import (
	"context"
	"net/http"
	"time"

	"github.com/kpango/glg"
	"github.com/pkg/errors"
)

// Server --
type Server interface {
	Serve() error
	ServeTLS(certFile, keyFile string) error
	Shutdown()
}

// Proxier --
type Proxier interface {
	Proxy(http.ResponseWriter, *http.Request)
}

// serverLoadBalancer --
type serverLoadBalancer struct {
	*Config
	*http.Server
}

// NewServerLoadBalancer --
func NewServerLoadBalancer(cfg *Config) Server {
	sbl := &serverLoadBalancer{
		Config: cfg,
	}

	sbl.Server = &http.Server{
		Handler: cfg.Balancing.Balancer(sbl),
	}

	return sbl
}

func (s *serverLoadBalancer) Proxy(w http.ResponseWriter, req *http.Request) {
	// TODO: not yet implemented
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

func (s *serverLoadBalancer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := s.Server.Shutdown(ctx)

	glg.Info("All http(s) requets finished")

	if err != nil {
		glg.Errorf("faild to shutdown server load balancer: %v", err)
	}
}
