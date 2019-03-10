package balancer

import (
	"net/http"
	"net/url"
)

// Handler represents an interface for balancing handler.
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	isBalaner()
}

// Proxier represents an interface for reverse proxy.
type Proxier interface {
	Proxy(*url.URL, http.ResponseWriter, *http.Request)
}
