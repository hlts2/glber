package server

// func (lb *LB) reverseProxy(scheme, destHost string, w http.ResponseWriter, req *http.Request) {
// 	req.URL.Scheme = scheme
// 	req.URL.Host = destHost
//
// 	lb.lf.Wait()
// 	resp, err := http.DefaultTransport.RoundTrip(req)
// 	if err != nil {
// 		lb.lf.Signal()
// 		glg.Println(err)
// 		return
// 	}
//
// 	lb.lf.Signal()
//
// 	copyHeader(w, resp)
//
// 	w.WriteHeader(resp.StatusCode)
//
// 	data := readCloserToByte(resp.Body)
// 	w.Write(data)
//
// 	resp.Body.Close()
// }
//
// func readCloserToByte(readCloser io.ReadCloser) []byte {
// 	buf := new(bytes.Buffer)
// 	io.Copy(buf, readCloser)
// 	return buf.Bytes()
// }
//
// func copyHeader(dest http.ResponseWriter, src *http.Response) {
// 	for key, values := range src.Header {
// 		dest.Header().Del(key)
// 		for _, value := range values {
// 			dest.Header().Add(key, value)
// 		}
// 	}
// }
