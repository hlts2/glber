package iphash

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
)

type iphash struct {
	urls    []*url.URL
	m       *sync.Map
	proxier balancer.Proxier

	balancer.Handler
}

func (h *iphash) ServeHTTP(http.ResponseWriter, *http.Request) {
	// TODO: not yet implemented
}

func (h *iphash) pick(url *url.URL) *url.URL {
	return nil
}

func (h *iphash) isBalaner() {}

// New returns balancer.Handler implementation(*iphash).
func New(urls []*url.URL, proxier balancer.Proxier) balancer.Handler {
	return &iphash{
		urls:    urls,
		m:       new(sync.Map),
		proxier: proxier,
	}
}
