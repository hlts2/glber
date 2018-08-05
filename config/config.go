package config

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config represents config file for loadbalancer
type Config struct {
	Servers   Servers `yaml:"servers"`
	Balancing string  `yaml:"balancing"`
}

// Server represents the server to connect to
type Server struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
}

// Servers is slice of Server
type Servers []Server

// ToStringSlice converts Servers to []string type
func (ss Servers) ToStringSlice() []string {
	hosts := make([]string, 0, len(ss))

	for _, s := range ss {
		hosts = append(hosts, s.Host)
	}
	return hosts
}

// LoadConfig loads config of loadbalancer
func LoadConfig(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
