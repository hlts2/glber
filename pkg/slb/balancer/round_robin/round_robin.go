package roundrobin

import (
	"net/http"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
)

type roundrobin struct {
	balancer.Balancer
}

func (h *roundrobin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO: not yet implemented
}

func (h *roundrobin) isBalaner() {}

// New --
func New(addrs []string, proxier balancer.Proxier) balancer.Balancer {
	return new(roundrobin)
}
