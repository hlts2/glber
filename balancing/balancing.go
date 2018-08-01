package balancing

import (
	iphash "github.com/hlts2/ip-hash"
	leastconnections "github.com/hlts2/least-connections"
	roundrobin "github.com/hlts2/round-robin"
)

// Balancing is wrapper object of balancing algorithm for reverse proxy
type Balancing struct {
	algorithm interface{}
}

// New returns Balancing object
func New(algorithm interface{}) *Balancing {
	switch algorithm.(type) {
	case leastconnections.LeastConnections, roundrobin.RoundRobin, iphash.IPHash:
		return &Balancing{
			algorithm: algorithm,
		}
	default:
		return nil
	}
}

// GetLeastConnections returns LeastConnections balancing algorithm interface
func (b *Balancing) GetLeastConnections() leastconnections.LeastConnections {
	return b.algorithm.(leastconnections.LeastConnections)
}

// GetRoundRobin returns RoundRobin balancing algorithm interface
func (b *Balancing) GetRoundRobin() roundrobin.RoundRobin {
	return b.algorithm.(roundrobin.RoundRobin)
}

// GetIPHash returns IPHash balancing algorithm interface
func (b *Balancing) GetIPHash() iphash.IPHash {
	return b.algorithm.(iphash.IPHash)
}
