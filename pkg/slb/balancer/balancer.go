package balancer

import (
	"net/http"
	"net/url"
)

// Balancer --
type Balancer interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	isBalaner()
}

// Proxier --
type Proxier interface {
	Proxy(*url.URL, http.ResponseWriter, *http.Request)
}
