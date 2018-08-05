package config

import (
	"errors"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// ErrDuplicateHost is error that there is duplicte host
var ErrDuplicateHost = errors.New("duplicate host")

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

// GetAddress returns address of servers
func (ss Servers) GetAddress() []string {
	hosts := make([]string, 0, len(ss))

	for _, s := range ss {
		hosts = append(hosts, s.Scheme+"://"+s.Host+":"+s.Port)
	}
	return hosts
}

// GetHostWithPort returns host of servers
// i.e) 192.168.33.10:1111
func (ss Servers) GetHostWithPort() []string {
	hosts := make([]string, 0, len(ss))

	for _, s := range ss {
		hosts = append(hosts, s.Host+":"+s.Port)
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

	ok := existsDuplicateHost(conf.Servers.GetHostWithPort())
	if ok {
		return ErrDuplicateHost
	}

	return nil
}

// existsDuplicateHost returns true if there is duplicte host
func existsDuplicateHost(hosts []string) bool {
	m := make(map[string]int, len(hosts))

	for _, host := range hosts {
		if _, ok := m[host]; ok {
			return true
		}
		m[host] = 0
	}
	return false
}
