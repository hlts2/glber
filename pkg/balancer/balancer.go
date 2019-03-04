package balancer

import "net/http"

// Balancer --
type Balancer interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
	isBalaner()
}

type balancer struct {
	next http.Handler
}

// New --
func New(next http.Handler) http.Handler {
	return &balancer{
		next: next,
	}
}

func (b *balancer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	b.next.ServeHTTP(w, req)
}
