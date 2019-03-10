package leastconnections

import (
	"net/http"
	"net/url"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
)

type leastconnections struct {
	balancer.Handler
}

func (h *leastconnections) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO: not yet implemented
}

func (h *leastconnections) isBalaner() {}

// New returns balancer.Handler implementation(*leastconnections).
func New(urls []url.URL, proxier balancer.Proxier) balancer.Handler {
	return new(leastconnections)
}
