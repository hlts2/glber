package slb

// func (lb *LB) leastConnectionsBalancing(w http.ResponseWriter, req *http.Request) {
// 	lc := lb.balancing.GetLeastConnections()
//
// 	destAddr := lc.Next()
// 	scheme, host := getSchemeAndHostWithPort(destAddr)
//
// 	lc.IncrementConnections(destAddr)
// 	lb.reverseProxy(scheme, host, w, req)
// 	lc.DecrementConnections(destAddr)
//
// 	req.Body.Close()
// }
//
// func (lb *LB) roundRobinBalancing(w http.ResponseWriter, req *http.Request) {
// 	rr := lb.balancing.GetRoundRobin()
//
// 	scheme, host := getSchemeAndHostWithPort(rr.Next())
// 	lb.reverseProxy(scheme, host, w, req)
//
// 	req.Body.Close()
// }
//
// func (lb *LB) ipHashBalancing(w http.ResponseWriter, req *http.Request) {
// 	ih := lb.balancing.GetIPHash()
//
// 	destAddr := ih.Next(req.RemoteAddr)
// 	scheme, host := getSchemeAndHostWithPort(destAddr)
// 	lb.reverseProxy(scheme, host, w, req)
//
// 	req.Body.Close()
// }
//
// // getSchemeAndHostWithPort returns the scheme name and host name with port
// // i.e) http://192.168.33.10:1111 => http, 192.168.33.10:1111
// func getSchemeAndHostWithPort(addr string) (string, string) {
// 	if addr[:5] == "https" {
// 		return addr[:5], addr[8:]
// 	} else if addr[:4] == "http" {
// 		return addr[:4], addr[7:]
// 	} else {
// 		return "", ""
// 	}
// }
