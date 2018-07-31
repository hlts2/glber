package server

import (
	"crypto/tls"
	"net/http"

	"github.com/hlts2/go-LB/config"
	"github.com/hlts2/least-connections"
	"github.com/hlts2/round-robin"
)

// Balancing is base balancing interface
type Balancing interface {
	Next() string
}

// LBServer represents load balancing server object
type LBServer struct {
	*http.Server
	lbConf    config.Config
	balancing Balancing
}

// NewLBServer returns LBServer object
func NewLBServer(addr string) *LBServer {
	lbs := new(LBServer)
	lbs.Addr = addr
	lbs.Handler = http.HandlerFunc(lbs.passthrogh)
	return lbs
}

// Build builds LB config
func (lbs *LBServer) Build(conf config.Config) *LBServer {
	lbs.lbConf = conf

	switch conf.Balancing {
	case "ip-hash":
		// TODO Load ip-hash balancing algorithm
	case "round-robin":
		rr, err := roundrobin.New(conf.Servers.ToStringSlice())
		if err == nil {
			lbs.balancing = rr
		}
	case "least-connections":
		ll, err := leastconnections.New(conf.Servers.ToStringSlice())
		if err == nil {
			lbs.balancing = ll
		}
	default:
		// TODO proxy
	}

	return lbs
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
