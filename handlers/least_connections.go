package handlers

import "net/http"

type leastConnections struct{}

func (l *leastConnections) ServeHTTP(w http.ResponseWriter, req *http.Request) {}

// NewLeastConnections --
func NewLeastConnections() http.Handler {
	return new(leastConnections)
}
