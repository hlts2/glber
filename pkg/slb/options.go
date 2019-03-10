package slb

import "github.com/hlts2/go-LB/pkg/slb/balancer"

// Option configures serverLoadBalancer.
type Option func(*serverLoadBalancer)

// WithBalancerHandler returns an Option that sets the Balancer.Handler implementation.
func WithBalancerHandler(handler balancer.Handler) func(*serverLoadBalancer) {
	return func(s *serverLoadBalancer) {
		// TODO:
	}
}
