package server

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/pkg/errors"

	b "github.com/hlts2/go-LB/balancing"
	"github.com/hlts2/go-LB/config"
	iphash "github.com/hlts2/ip-hash"
	leastconnections "github.com/hlts2/least-connections"
	lockfree "github.com/hlts2/lock-free"
	roundrobin "github.com/hlts2/round-robin"
)

// ErrNotBalancingAlgorithm is error that balancing algorithm dose not found
var ErrNotBalancingAlgorithm = errors.New("balancing algorithm dose not found")

// LB represents load balancer
type LB struct {
	*http.Server
	balancing *b.Balancing
	lf        lockfree.LockFree
}

// NewLB returns LB object
func NewLB(addr string) *LB {
	return &LB{
		Server: &http.Server{
			Addr: addr,
		},
		lf: lockfree.New(),
	}
}

// Build builds config for load balancer
func (lb *LB) Build(conf config.Config) (*LB, error) {
	switch conf.Balancing {
	case "ip-hash":
		ih, err := iphash.New(conf.Servers.ToStringSlice())
		if err != nil {
			return nil, errors.Wrap(err, "ip-hash algorithm")
		}

		lb.balancing = b.New(ih)
		lb.Handler = http.HandlerFunc(lb.ipHashBalancing)
	case "round-robin":
		rr, err := roundrobin.New(conf.Servers.ToStringSlice())
		if err == nil {
			return nil, errors.Wrap(err, "round-robin algorithm")
		}

		lb.balancing = b.New(rr)
		lb.Handler = http.HandlerFunc(lb.roundRobinBalancing)
	case "least-connections":
		lc, err := leastconnections.New(conf.Servers.ToStringSlice())
		if err == nil {
			return nil, errors.Wrap(err, "least-connections algorithm")
		}

		lb.balancing = b.New(lc)
		lb.Handler = http.HandlerFunc(lb.ipHashBalancing)
	default:
		return nil, ErrNotBalancingAlgorithm
	}

	return lb, nil
}

// ServeTLS runs load balancer with TLS
func (lb *LB) ServeTLS(tlsConfig *tls.Config, certFile, keyFile string) error {
	lisner, err := net.Listen("tcp", lb.Addr)
	if err != nil {
		return err
	}

	lb.TLSConfig = tlsConfig

	err = lb.Server.ServeTLS(lisner, certFile, keyFile)
	if err != nil {
		return err
	}

	return nil
}

// Serve runs load balancer
func (lb *LB) Serve() error {
	lisner, err := net.Listen("tcp", lb.Addr)
	if err != nil {
		return err
	}

	err = lb.Server.Serve(lisner)
	if err != nil {
		return err
	}

	return nil
}
