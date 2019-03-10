package slb

import (
	"net/url"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
)

// Option configures serverLoadBalancer.
type Option func(*serverLoadBalancer)

// HandlerDirector --
type HandlerDirector func(urls []url.URL, proxier balancer.Proxier) balancer.Handler

// WithBalancingHandlerDirector returns an Option that sets the Balancer.Handler implementation.
func WithBalancingHandlerDirector(f HandlerDirector) func(*serverLoadBalancer) {
	return func(s *serverLoadBalancer) {
		// TODO:
	}
}
