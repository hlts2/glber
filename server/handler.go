package server

import (
	"bytes"
	"io"
	"net/http"

	"github.com/kpango/glg"
)

func (lbs *LBServer) leastConnectionsBalancing(w http.ResponseWriter, req *http.Request) {
	lc := lbs.balancing.GetLeastConnections()

	destAddr := lc.Next()

	lc.IncrementConnections(destAddr)
	lbs.reverseProxy(destAddr, w, req)
	lc.DecrementConnections(destAddr)
}

func (lbs *LBServer) roundRobinBalancing(w http.ResponseWriter, req *http.Request) {
	rr := lbs.balancing.GetRoundRobin()

	destAddr := rr.Next()
	lbs.reverseProxy(destAddr, w, req)
}

func (lbs *LBServer) ipHashBalancing(w http.ResponseWriter, req *http.Request) {
	ih := lbs.balancing.GetIPHash()

	destAddr := ih.Next(req.RemoteAddr)
	lbs.reverseProxy(destAddr, w, req)
}

// TODO copy header
// TODO add header for proxy
func (lbs *LBServer) reverseProxy(destAddr string, w http.ResponseWriter, req *http.Request) {
	req.Host = destAddr

	lbs.lf.Wait()
	resp, err := http.DefaultTransport.RoundTrip(req)
	lbs.lf.Signal()

	if err != nil {
		glg.Println(err)
		return
	}

	defer resp.Body.Close()

	for _, cokkie := range resp.Cookies() {
		http.SetCookie(w, cokkie)
	}

	contents := readCloserToByte(resp.Body)
	if err != nil {
		glg.Println(err)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(contents)
}

func readCloserToByte(readCloser io.ReadCloser) []byte {
	buf := new(bytes.Buffer)
	io.Copy(buf, readCloser)
	return buf.Bytes()
}
