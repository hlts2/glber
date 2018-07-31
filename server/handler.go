package server

import (
	"net/http"

	iphash "github.com/hlts2/ip-hash"
	"github.com/hlts2/least-connections"
	"github.com/hlts2/round-robin"
)

func (s *LBServer) balancingLeastConnections(w http.ResponseWriter, req *http.Request) {
	lc := s.balancing.(leastconnections.LeastConnections)

	destAddr := lc.Next()

	lc.IncrementConnections(destAddr)
	s.passthrogh(w, req)
	lc.DecrementConnections(destAddr)

}

func (s *LBServer) balancingRoundRobin(w http.ResponseWriter, req *http.Request) {
	rr := s.balancing.(roundrobin.RoundRobin)
	_ = rr.Next()
	s.passthrogh(w, req)
}

func (s *LBServer) balancingIPHash(w http.ResponseWriter, req *http.Request) {
	ic := s.balancing.(iphash.IPHash)
	_ = ic.Next(req.RemoteAddr)
	s.passthrogh(w, req)
}

func (s *LBServer) passthrogh(w http.ResponseWriter, req *http.Request) {

}
