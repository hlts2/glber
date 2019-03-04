package leastconnections

import (
	"net/http"

	"google.golang.org/grpc/balancer"
)

type leastconnections struct {
	balancer.Balancer
}

func (h *leastconnections) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO: not yet implemented
}

func (h *leastconnections) isBalaner() {}

// New --
func New() balancer.Balancer {
	return new(leastconnections)
}
