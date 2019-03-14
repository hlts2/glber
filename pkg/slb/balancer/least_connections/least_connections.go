package leastconnections

import (
	"net/http"
	"net/url"

	"github.com/hlts2/least-connections"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
)

type leastconnectionsHandler struct {
	urls    []*url.URL
	proxier balancer.Proxier
	lc      leastconnections.LeastConnections

	balancer.Handler
}

func (h *leastconnectionsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	src, done := h.lc.Next()
	h.proxier.Proxy(src, w, req)
	done()
}

func (h *leastconnectionsHandler) isBalaner() {}

// New returns balancer.Handler implementation(*leastconnections).
func New(urls []*url.URL, proxier balancer.Proxier) balancer.Handler {
	lc, _ := leastconnections.New(urls)

	return &leastconnectionsHandler{
		urls:    urls,
		proxier: proxier,
		lc:      lc,
	}
}
