package balancer

import "net/http"

// Balancer --
type Balancer interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	isBalaner()
}

// Proxier --
type Proxier interface {
	Proxy(http.ResponseWriter, *http.Request)
}
