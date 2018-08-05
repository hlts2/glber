package config

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config represents config file of load balancer
type Config struct {
	Servers   Servers `yaml:"servers"`
	Balancing string  `yaml:"balancing"`
}

// Server represents the server to connect to
type Server struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
}

// Servers is slice of Server
type Servers []Server

// GetAddresses returns address of servers
func (ss Servers) GetAddresses() []string {
	hosts := make([]string, 0, len(ss))

	for _, s := range ss {
		hosts = append(hosts, s.Scheme+"://"+s.Host+":"+s.Port)
	}
	return hosts
}

// LoadConfig loads config of load balancer
func LoadConfig(filename string, conf *Config) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	err = yaml.NewDecoder(f).Decode(conf)
	if err != nil {
		return err
	}

	return nil
}
