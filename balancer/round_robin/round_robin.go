package roundrobin

import (
	"net/http"

	"google.golang.org/grpc/balancer"
)

type roundrobin struct {
	balancer.Balancer
}

func (h *roundrobin) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO: not yet implemented
}

func (h *roundrobin) isBalaner() {}

// New --
func New() balancer.Balancer {
	return new(roundrobin)
}
