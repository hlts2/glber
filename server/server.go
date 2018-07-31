package server

import (
	"crypto/tls"
	"net/http"

	"github.com/hlts2/go-LB/config"
	iphash "github.com/hlts2/ip-hash"
	"github.com/hlts2/least-connections"
	"github.com/hlts2/round-robin"
)

// Balancing is custom type of balancing algorithm
type Balancing interface{}

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
	return lbs
}

// Build builds LB config
func (lbs *LBServer) Build(conf config.Config) *LBServer {
	lbs.lbConf = conf

	switch conf.Balancing {
	case "ip-hash":
		ih, err := iphash.New(conf.Servers.ToStringSlice())
		if err == nil {
			lbs.balancing = ih
		}
		lbs.Handler = http.HandlerFunc(lbs.balancingLeastConnections)
	case "round-robin":
		rr, err := roundrobin.New(conf.Servers.ToStringSlice())
		if err == nil {
			lbs.balancing = rr
		}
		lbs.Handler = http.HandlerFunc(lbs.balancingRoundRobin)
	case "least-connections":
		ll, err := leastconnections.New(conf.Servers.ToStringSlice())
		if err == nil {
			lbs.balancing = ll
		}
		lbs.Handler = http.HandlerFunc(lbs.balancingIPHash)
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
