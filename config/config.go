package config

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config represents config file for load balancer
type Config struct {
	Servers   Servers `yaml:"servers"`
	Balancing string  `yaml:"balancing"`
}

// Server represents domain name for server
type Server string

// Servers is Server(string) slice
type Servers []Server

// LoadConfig load config for load balancer
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
