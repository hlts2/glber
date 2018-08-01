package server

import (
	"crypto/tls"
	"net/http"

	b "github.com/hlts2/go-LB/balancing"
	"github.com/hlts2/go-LB/config"
	iphash "github.com/hlts2/ip-hash"
	"github.com/hlts2/least-connections"
	"github.com/hlts2/round-robin"
	"github.com/kpango/glg"
)

// LBServer represents load balancing server object
type LBServer struct {
	*http.Server
	balancing *b.Balancing
}

// NewLBServer returns LBServer object
func NewLBServer(addr string) *LBServer {
	lbs := new(LBServer)
	lbs.Addr = addr
	return lbs
}

// Build builds LB config
func (lbs *LBServer) Build(conf config.Config) *LBServer {
	switch conf.Balancing {
	case "ip-hash":
		ih, err := iphash.New(conf.Servers.ToStringSlice())
		if err != nil {
			return nil
		}

		lbs.balancing = b.New(ih)
		lbs.Handler = http.HandlerFunc(lbs.ipHashBalancing)
	case "round-robin":
		rr, err := roundrobin.New(conf.Servers.ToStringSlice())
		if err == nil {
			return nil
		}

		lbs.balancing = b.New(rr)
		lbs.Handler = http.HandlerFunc(lbs.roundRobinBalancing)
	case "least-connections":
		lc, err := leastconnections.New(conf.Servers.ToStringSlice())
		if err == nil {
			return nil
		}

		lbs.balancing = b.New(lc)
		lbs.Handler = http.HandlerFunc(lbs.ipHashBalancing)
	default:
		glg.Fatal("balancing algorithm dose not found")
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
