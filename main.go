package main

import (
	"flag"

	"github.com/hlts2/go-LB/config"
	"github.com/hlts2/go-LB/server"
	"github.com/kpango/glg"
)

var (
	configFileN string
	addr        string
	port        string
)

func init() {
	flag.StringVar(&configFileN, "s", "config.yaml", "set a config file of load balancer")
	flag.StringVar(&addr, "a", "0.0.0.0", "set a host address of load balancer")
	flag.StringVar(&port, "p", "80", "set a port number of load balancer")
	flag.Parse()
}

func main() {
	conf, err := config.LoadConfig(configFileN)
	if err != nil {
		glg.Fatalln(err)
	}

	lb := server.NewLB(addr + ":" + port)

	err = lb.Build(*conf).Serve()
	if err != nil {
		glg.Fatalln(err)
	}
}
