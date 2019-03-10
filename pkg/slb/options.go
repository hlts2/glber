package slb

import "github.com/hlts2/go-LB/pkg/slb/balancer"

// Option configures Config.
type Option func(*Config)

// WithBalancingHandler returns an Option that sets the Balancer.Handler implementation.
func WithBalancingHandler(handler balancer.Handler) func(*Config) {
	return func(cfg *Config) {}
}
