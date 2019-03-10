package roundrobin

import (
	"net/http"
	"net/url"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
)

type roundrobin struct {
	balancer.Handler
}

func (h *roundrobin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO: not yet implemented
}

func (h *roundrobin) isBalaner() {}

// New returns balancer.Handler implementation(*roundrobin).
func New(addrs []url.URL, proxier balancer.Proxier) balancer.Handler {
	return new(roundrobin)
}
