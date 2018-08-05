package server

import (
	"bytes"
	"io"
	"net/http"

	"github.com/kpango/glg"
)

func (lb *LB) leastConnectionsBalancing(w http.ResponseWriter, req *http.Request) {
	lc := lb.balancing.GetLeastConnections()

	dest := lc.Next()
	scheme, host := getSchemeAndHost(dest)

	lc.IncrementConnections(dest)
	lb.reverseProxy(scheme, host, w, req)
	lc.DecrementConnections(dest)

	req.Body.Close()
}

func (lb *LB) roundRobinBalancing(w http.ResponseWriter, req *http.Request) {
	rr := lb.balancing.GetRoundRobin()

	scheme, host := getSchemeAndHost(rr.Next())
	lb.reverseProxy(scheme, host, w, req)

	req.Body.Close()
}

func (lb *LB) ipHashBalancing(w http.ResponseWriter, req *http.Request) {
	ih := lb.balancing.GetIPHash()

	scheme, host := getSchemeAndHost(ih.Next(req.RemoteAddr))
	lb.reverseProxy(scheme, host, w, req)

	req.Body.Close()
}

func (lb *LB) reverseProxy(scheme, destHost string, w http.ResponseWriter, req *http.Request) {
	req.URL.Scheme = scheme
	req.URL.Host = destHost

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

func getSchemeAndHost(url string) (string, string) {
	if url[:5] == "https" {
		return url[:5], url[8:]
	} else if url[:4] == "http" {
		return url[:4], url[7:]
	} else {
		return "", ""
	}
}

func copyHeader(dest http.ResponseWriter, src *http.Response) {
	for key, values := range src.Header {
		dest.Header().Del(key)
		for _, value := range values {
			dest.Header().Add(key, value)
		}
	}
}
