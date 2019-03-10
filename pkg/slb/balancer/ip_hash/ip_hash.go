package iphash

import (
	"net/http"
	"net/url"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
)

type iphash struct {
	balancer.Handler
}

func (h *iphash) ServeHTTP(http.ResponseWriter, *http.Request) {
	// TODO: not yet implemented
}

func (h *iphash) isBalaner() {}

// New returns balancer.Handler implementation(*iphash).
func New(addrs []url.URL, proxier balancer.Proxier) balancer.Handler {
	return nil
}
