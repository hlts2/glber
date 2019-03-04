package handlers

import "net/http"

type roundrobin struct{}

func (r *roundrobin) ServeHTTP(w http.ResponseWriter, req *http.Request) {}

// NewRoundRobin --
func NewRoundRobin() http.Handler {
	return new(roundrobin)
}
