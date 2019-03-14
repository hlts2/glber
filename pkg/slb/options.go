package slb

import (
	"net/url"

	"github.com/hlts2/glber/pkg/slb/balancer"
)

// Option configures serverLoadBalancer.
type Option func(*serverLoadBalancer)

// HandlerDirector type is director to generate handler.Handler implementation.
type HandlerDirector func(urls []*url.URL, proxier balancer.Proxier) balancer.Handler

// WithBalancingHandlerDirector returns an Option that sets the Balancer.Handler implementation.
func WithBalancingHandlerDirector(d HandlerDirector) func(*serverLoadBalancer) {
	return func(s *serverLoadBalancer) {
		if d != nil {
			s.HandlerDirector = d
		}
	}
}
