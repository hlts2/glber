package iphash

import (
	"net/http"
	"net/url"

	"github.com/hlts2/ip-hash"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
)

type iphashHandler struct {
	urls    []*url.URL
	proxier balancer.Proxier
	iphash  iphash.IPHash

	balancer.Handler
}

func (h *iphashHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	src := h.iphash.Next(req.URL)
	h.proxier.Proxy(src, w, req)
}

func (h *iphashHandler) isBalaner() {}

// New returns balancer.Handler implementation(*iphash).
func New(urls []*url.URL, proxier balancer.Proxier) balancer.Handler {
	i, _ := iphash.New(urls)

	return &iphashHandler{
		urls:    urls,
		proxier: proxier,
		iphash:  i,
	}
}
