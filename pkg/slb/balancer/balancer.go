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

// Proxier is an interface for representing reverse proxy implementation.
type Proxier interface {
	Proxy(*url.URL, http.ResponseWriter, *http.Request)
}
