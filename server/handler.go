package server

import (
	"bytes"
	"io"
	"net/http"

	"github.com/kpango/glg"
)

func (lb *LB) leastConnectionsBalancing(w http.ResponseWriter, req *http.Request) {
	lc := lb.balancing.GetLeastConnections()

	destAddr := lc.Next()

	lc.IncrementConnections(destAddr)
	lb.reverseProxy(destAddr, w, req)
	lc.DecrementConnections(destAddr)

	req.Body.Close()
}

func (lb *LB) roundRobinBalancing(w http.ResponseWriter, req *http.Request) {
	rr := lb.balancing.GetRoundRobin()

	destAddr := rr.Next()
	lb.reverseProxy(destAddr, w, req)

	req.Body.Close()
}

func (lb *LB) ipHashBalancing(w http.ResponseWriter, req *http.Request) {
	ih := lb.balancing.GetIPHash()

	destAddr := ih.Next(req.RemoteAddr)
	lb.reverseProxy(destAddr, w, req)

	req.Body.Close()
}

func (lb *LB) reverseProxy(destAddr string, w http.ResponseWriter, req *http.Request) {
	req.Host = destAddr

	lb.lf.Wait()
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		lb.lf.Signal()
		glg.Println(err)
		return
	}

	lb.lf.Signal()

	for _, cokie := range resp.Cookies() {
		http.SetCookie(w, cokie)
	}

	copyHeader(w, resp)

	w.WriteHeader(resp.StatusCode)

	data := readCloserToByte(resp.Body)
	w.Write(data)

	resp.Body.Close()
}

func readCloserToByte(readCloser io.ReadCloser) []byte {
	buf := new(bytes.Buffer)
	io.Copy(buf, readCloser)
	return buf.Bytes()
}

func copyHeader(dest http.ResponseWriter, src *http.Response) {
	for key, values := range src.Header {
		dest.Header().Del(key)
		for _, value := range values {
			dest.Header().Add(key, value)
		}
	}
}
