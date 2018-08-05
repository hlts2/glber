package main

import (
	"flag"

	"github.com/hlts2/go-LB/config"
	"github.com/hlts2/go-LB/server"
	"github.com/kpango/glg"
)

var (
	configFileN string
	host        string
	port        string
)

func init() {
	flag.StringVar(&configFileN, "s", "config.yaml", "set a config file of load balancer")
	flag.StringVar(&host, "a", "127.0.0.1", "set a host name or IP address of load balancer")
	flag.StringVar(&port, "p", "8080", "set a port number of load balancer")
	flag.Parse()
}

func main() {
	var conf config.Config

	err := config.LoadConfig(configFileN, &conf)
	if err != nil {
		glg.Fatalln(err)
	}

	lb := server.NewLB(host + ":" + port)

	err = lb.Build(conf).Serve()
	if err != nil {
		glg.Fatalln(err)
	}
}
