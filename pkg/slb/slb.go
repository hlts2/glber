package slb

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/kpango/glg"
	"github.com/pkg/errors"
)

// Server is an interface for representing server load balancer implementation.
type Server interface {
	Serve() error
	ServeTLS(certFile, keyFile string) error
	Shutdown()
}

type serverLoadBalancer struct {
	*Config
	*http.Server
	Director        func(*url.URL) func(*http.Request)
	HandlerDirector HandlerDirector
}

// CreateSLB returns Server implementation(*serverLoadBalancer) from the given Config.
func CreateSLB(cfg *Config, ops ...Option) (Server, error) {
	if err := cfg.validate(); err != nil {
		return nil, errors.Wrap(err, "invalid configuration")
	}

	sbl := &serverLoadBalancer{
		Config: cfg,
		Director: func(target *url.URL) func(*http.Request) {
			return func(req *http.Request) {
				req.URL.Scheme = target.Scheme
				req.URL.Host = target.Host
				req.URL.Path = target.Path

				if target.RawQuery == "" || req.URL.RawQuery == "" {
					req.URL.RawQuery = target.RawQuery + req.URL.RawQuery
				} else {
					req.URL.RawQuery = target.RawQuery + "&" + req.URL.RawQuery
				}
				if _, ok := req.Header["User-Agent"]; !ok {
					req.Header.Set("User-Agent", "")
				}
			}
		},
		HandlerDirector: cfg.Balancing.Handler,
	}
	sbl.apply(ops...)

	sbl.Server = &http.Server{
		Handler: sbl.HandlerDirector(cfg.BackendServerConfigs.getURLs(), sbl),
	}

	return sbl, nil
}

func (s *serverLoadBalancer) apply(ops ...Option) {
	for _, op := range ops {
		op(s)
	}
}

func (s *serverLoadBalancer) Proxy(target *url.URL, w http.ResponseWriter, req *http.Request) {
	(&httputil.ReverseProxy{Director: s.Director(target)}).ServeHTTP(w, req)
}

func (s *serverLoadBalancer) Serve() error {
	lis, err := s.LoadBalancerConfig.createListener()
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
	lis, err := s.LoadBalancerConfig.createListener()
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
