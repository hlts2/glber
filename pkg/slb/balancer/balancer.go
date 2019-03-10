package balancer

import (
	"net/http"
	"net/url"
)

// Handler is an interface for representing balancing handler implementation.
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	isBalaner()
}

// NopHandler --
type NopHandler struct{}

func (n *NopHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}
func (n *NopHandler) isBalaner()                                   {}

// Proxier is an interface for representing reverse proxy implementation.
type Proxier interface {
	Proxy(*url.URL, http.ResponseWriter, *http.Request)
}
