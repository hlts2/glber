package leastconnections

import (
	"net/http"

	"github.com/hlts2/go-LB/balancer"
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
