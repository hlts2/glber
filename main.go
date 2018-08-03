package main

import (
	"flag"
)

var (
	cofingFileN string
)

func init() {
	flag.StringVar(&cofingFileN, "s", "config.yaml", "set a config file of load balancer")
	flag.Parse()
}

func main() {
}
