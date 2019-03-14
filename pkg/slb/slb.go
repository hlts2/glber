package slb

import (
	"context"
	"crypto/tls"
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
		HandlerDirector: cfg.Balancing.CreateHandler,
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

func (s *serverLoadBalancer) ListenAndServe() error {
	addr := s.Config.Host + ":" + s.Config.Port

	var (
		ls  net.Listener
		err error
	)

	if s.Config.TLSConfig.Enabled {
		ls, err = createTLSListenter(addr, s.Config.TLSConfig.CertKey, s.Config.TLSConfig.KeyKey)
		if err != nil {
			return errors.Wrap(err, "faild to create tls lisner")
		}
	} else {
		ls, err = createListener(addr)
		if err != nil {
			return errors.Wrap(err, "faild to create listener")
		}
	}

	err = s.Server.Serve(ls)
	if err != nil {
		return errors.Wrap(err, "faild to serve")
	}
	return nil
}

func createListener(addr string) (net.Listener, error) {
	ls, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, errors.Wrapf(err, "faild to create lisner, network: tcp, addr: %s", addr)
	}
	return ls, nil
}

func createTLSListenter(addr string, certFile, keyFile string) (net.Listener, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, errors.Wrapf(err, "faild to load 509 key parir, certFile: %s, keyFile: %s", certFile, keyFile)
	}

	cfg := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	ls, err := tls.Listen("tcp", addr, cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "faild to create listener, network: tcp, addr: %s", addr)
	}

	return ls, nil
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
