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

// Server is address of address
type Server string

// Servers is slice of Server
type Servers []string

// ToStringSlice converts Servers to []string type
func (ss Servers) ToStringSlice() []string {
	return []string(ss)
}

// LoadConfig loads config for loadbalancer
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
