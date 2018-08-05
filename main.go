package main

import (
	"crypto/tls"
	"flag"
	"path/filepath"

	"github.com/hlts2/go-LB/config"
	"github.com/hlts2/go-LB/server"
	"github.com/kpango/glg"
)

var (
	configFileN string
	host        string
	port        string
	tlspath     string
)

func init() {
	flag.StringVar(&configFileN, "s", "config.yaml", "set a config file of load balancer")
	flag.StringVar(&host, "h", "127.0.0.1", "set a host name or IP address of load balancer")
	flag.StringVar(&port, "p", "8080", "set a port number of load balancer")
	flag.StringVar(&tlspath, "tlf-path", "", "set a TLS directory")
	flag.Parse()
}

func main() {
	var conf config.Config

	err := config.LoadConfig(configFileN, &conf)
	if err != nil {
		glg.Fatalln(err)
	}

	lb := server.NewLB(host + ":" + port)
	lb.Build(conf)

	// NOT TLS Mode
	if tlspath == "" {
		err = lb.Serve()
		if err != nil {
			glg.Fatalln(err)
		}
	} else {
		var (
			certfile = filepath.Join(tlspath, "server.pem")
			keyfile  = filepath.Join(tlspath, "key.pen")
		)
		cer, err := tls.LoadX509KeyPair(certfile, keyfile)
		if err != nil {
			glg.Fatalln(err)
		}

		tlsConfig := &tls.Config{
			Certificates: []tls.Certificate{cer},
		}

		lb.ServeTLS(tlsConfig, certfile, keyfile)
	}
}
