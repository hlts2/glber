package handlers

import "net/http"

type iphash struct{}

func (i *iphash) ServeHTTP(w http.ResponseWriter, req *http.Request) {}

// NewIPHash --
func NewIPHash() http.Handler {
	return new(iphash)
}
