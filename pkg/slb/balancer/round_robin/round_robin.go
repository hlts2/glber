package roundrobin

import (
	"net/http"
	"net/url"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
	"github.com/hlts2/round-robin"
)

type roundrobinHandler struct {
	rr      roundrobin.RoundRobin
	proxier balancer.Proxier
	balancer.Handler
}

func (h *roundrobinHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.proxier.Proxy(h.rr.Next(), w, req)
}

func (h *roundrobinHandler) isBalaner() {}

// New returns balancer.Handler implementation(*roundrobin).
func New(urls []*url.URL, proxier balancer.Proxier) balancer.Handler {
	rr, _ := roundrobin.New(urls)

	return &roundrobinHandler{
		rr:      rr,
		proxier: proxier,
	}
}
