package server

import (
	"net/http"
)

func (lbs *LBServer) leastConnectionsBalancing(w http.ResponseWriter, req *http.Request) {
	lc := lbs.getLeastConnections()

	destAddr := lc.Next()

	lc.IncrementConnections(destAddr)
	lbs.reverseProxy(destAddr, w, req)
	lc.DecrementConnections(destAddr)

}

func (lbs *LBServer) roundRobinBalancing(w http.ResponseWriter, req *http.Request) {
	rr := lbs.getRoundRobin()
	destAddr := rr.Next()
	lbs.reverseProxy(destAddr, w, req)
}

func (lbs *LBServer) ipHashBalancing(w http.ResponseWriter, req *http.Request) {
	ih := lbs.getIPHash()
	destAddr := ih.Next(req.RemoteAddr)
	lbs.reverseProxy(destAddr, w, req)
}

func (lbs *LBServer) reverseProxy(destAddr string, w http.ResponseWriter, req *http.Request) {

}
