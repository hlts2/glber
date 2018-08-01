package server

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/pkg/errors"

	b "github.com/hlts2/go-LB/balancing"
	"github.com/hlts2/go-LB/config"
	iphash "github.com/hlts2/ip-hash"
	"github.com/hlts2/least-connections"
	"github.com/hlts2/round-robin"
)

// ErrNotBalancingAlgorithm is error that balancing algorithm dose not found
var ErrNotBalancingAlgorithm = errors.New("balancing algorithm dose not found")

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
func (lbs *LBServer) Build(conf config.Config) (*LBServer, error) {
	switch conf.Balancing {
	case "ip-hash":
		ih, err := iphash.New(conf.Servers.ToStringSlice())
		if err != nil {
			return nil, errors.Wrap(err, "ip-hash algorithm")
		}

		lbs.balancing = b.New(ih)
		lbs.Handler = http.HandlerFunc(lbs.ipHashBalancing)
	case "round-robin":
		rr, err := roundrobin.New(conf.Servers.ToStringSlice())
		if err == nil {
			return nil, errors.Wrap(err, "round-robin algorithm")
		}

		lbs.balancing = b.New(rr)
		lbs.Handler = http.HandlerFunc(lbs.roundRobinBalancing)
	case "least-connections":
		lc, err := leastconnections.New(conf.Servers.ToStringSlice())
		if err == nil {
			return nil, errors.Wrap(err, "least-connections algorithm")
		}

		lbs.balancing = b.New(lc)
		lbs.Handler = http.HandlerFunc(lbs.ipHashBalancing)
	default:
		return nil, ErrNotBalancingAlgorithm
	}

	return lbs, nil
}

// ServeTLS runs load balancing server with TLS
func (lbs *LBServer) ServeTLS(tlsConfig *tls.Config, certFile, keyFile string) error {
	lsn, err := net.Listen("tcp", lbs.Addr)
	if err != nil {
		return err
	}

	lbs.TLSConfig = tlsConfig

	err = lbs.Server.ServeTLS(lsn, certFile, keyFile)
	if err != nil {
		return err
	}

	return nil
}

// Serve runs load balancing server
func (lbs *LBServer) Serve() error {
	lsn, err := net.Listen("tcp", lbs.Addr)
	if err != nil {
		return err
	}

	err = lbs.Server.Serve(lsn)
	if err != nil {
		return err
	}

	return nil
}
