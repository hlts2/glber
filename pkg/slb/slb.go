package slb

import (
	"context"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// Server is an interface for representing server load balancer implementation.
type Server interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

type serverLoadBalancer struct {
	*Config
	*http.Server
	RequestDirector func(target *url.URL) func(*http.Request)
	HandlerDirector HandlerDirector
}

// CreateSLB returns Server implementation(*serverLoadBalancer) from the given Config.
func CreateSLB(cfg *Config, ops ...Option) (Server, error) {
	err := cfg.validate()
	if err != nil {
		return nil, errors.Wrap(err, "invalid configuration")
	}

	sbl := &serverLoadBalancer{
		Config: cfg,
		RequestDirector: func(target *url.URL) func(*http.Request) {
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
	(&httputil.ReverseProxy{
		Director: s.RequestDirector(target),
	}).ServeHTTP(w, req)
}

func (s *serverLoadBalancer) Serve() error {
	lis, err := createListener(s.ServerConfig.Host, s.ServerConfig.Port)
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
	lis, err := createListener(s.ServerConfig.Host, s.ServerConfig.Port)
	if err != nil {
		return errors.Wrap(err, "faild to create listener")
	}

	err = s.Server.ServeTLS(lis, certFile, keyFile)
	if err != nil {
		return errors.Wrap(err, "faild to serve with TLS")
	}

	return nil
}

func createListener(host, port string) (net.Listener, error) {
	addr := host + ":" + port

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, errors.Wrapf(err, "faild to listen: %v", addr)
	}

	return lis, nil
}

func (s *serverLoadBalancer) ListenAndServe() error {
	return nil
}

func (s *serverLoadBalancer) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err := s.Server.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, "faild to shutdown")
	}
	return nil
}
