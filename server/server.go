package server

import (
	"crypto/tls"
	"net/http"

	"github.com/hlts2/go-LB/config"
)

// LBServer represents load balancing server object
type LBServer struct {
	*http.Server
}

// Build builds LB config
func (lbs *LBServer) Build(conf config.Config) {

}

// ListenAndServeTLS runs load balancing server with TLS
func (lbs *LBServer) ListenAndServeTLS(tlsConfig *tls.Config, certFile, keyFile string) error {
	lbs.TLSConfig = tlsConfig

	err := lbs.Server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		return err
	}

	return nil
}

// ListenAndServe runs load balancing server
func (lbs *LBServer) ListenAndServe() error {
	err := lbs.Server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// NewLBServer returns LBServer object
func NewLBServer(addr string) *LBServer {
	lbs := new(LBServer)
	lbs.Addr = addr
	lbs.Handler = http.HandlerFunc(lbs.passthrogh)
	return lbs
}
