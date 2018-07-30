package server

import (
	"crypto/tls"
	"net/http"

	"github.com/hlts2/go-LB/config"
)

// LBServerTLS represents base TLS load balancing server interface
type LBServerTLS interface {
	ServeTLS(string, string) error
}

// LBServer represents base load balancing server interface
type LBServer interface {
	Serve() error
}

// lbServer represents load balancing server object
type lbServer struct {
	s *http.Server
}

// ServeTLS runs load balancing server with TLS
func (lbs *lbServer) ServeTLS(certFile, keyFile string) error {
	// TODO output loging mesasge

	err := lbs.s.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		return err
	}

	return nil
}

// Serv runs load balancing server
func (lbs *lbServer) Serve() error {
	// TODO output loging mesasge

	err := lbs.s.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// NewLBServerByTLS returns Server(*server) object
func NewLBServerByTLS(addr string, tlsConfig *tls.Config, lbConf config.Config) LBServerTLS {
	lbs := new(lbServer)

	lbs.s = &http.Server{
		Addr:      addr,
		TLSConfig: tlsConfig,
		Handler:   http.HandlerFunc(lbs.handler),
	}

	return lbs
}

// NewLBServer returns Server(*server) object
func NewLBServer(addr string, lbConf config.Config) LBServer {
	lbs := new(lbServer)

	lbs.s = &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(lbs.handler),
	}

	return lbs
}
