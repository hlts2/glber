package roundrobin

import (
	"net/http"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
)

type roundrobin struct {
	balancer.Handler
}

func (h *roundrobin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO: not yet implemented
}

func (h *roundrobin) isBalaner() {}

// New --
func New(addrs []string, proxier balancer.Proxier) balancer.Handler {
	return new(roundrobin)
}
