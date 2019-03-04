package iphash

import (
	"net/http"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
)

type iphash struct {
	balancer.Balancer
}

func (h *iphash) ServeHTTP(http.ResponseWriter, *http.Request) {
	// TODO: not yet implemented
}

func (h *iphash) isBalaner() {}

// New --
func New() balancer.Balancer {
	return nil
}
