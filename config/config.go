package config

import (
	"net/url"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// Config represents an application configuration content (config.yaml).
type Config struct {
	Servers   `yaml:"servers"`
	Balancing string `yaml:"balancing"`
}

// Server represents configuration content for server.
type Server struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
}

// Servers is Server slice.
type Servers []Server

func (s Server) address() string {
	return s.Scheme + "://" + s.Host + ":" + s.Port
}

func (ss Servers) validate() error {
	hostports := make([]string, len(ss))

	for i, s := range ss {
		if len(s.Scheme) == 0 || len(s.Host) == 0 || len(s.Port) == 0 {
			return errors.Errorf("invalid server configuration, scheme: %v, host: %v, port: %v", s.Scheme, s.Host, s.Port)
		}

		addr := s.address()

		_, err := url.ParseRequestURI(s.address())
		if err != nil {
			return errors.Wrapf(err, "invalid address: %s", addr)
		}

		// http://127.0.0.1:80 => 127.0.0.1:80
		hostports[i] = addr[len(s.Scheme)+3:]
	}

	ok := duplicateExists(hostports)
	if ok {
		return errors.Errorf("duplicate host and port exists: %v", hostports)
	}

	return nil
}

// duplicateExists returns true if there is duplicte.
func duplicateExists(vs []string) bool {
	m := make(map[string]bool, len(vs))

	for _, v := range vs {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}

// GetAddresses returns address of servers
func (ss Servers) GetAddresses() []string {
	addresses := make([]string, len(ss))

	for i, s := range ss {
		addresses[i] = s.address()
	}
	return addresses
}

// Load loads configuration content of the given the path.
func Load(path string, cfg *Config) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "faild to open configuration file")
	}

	err = yaml.NewDecoder(f).Decode(cfg)
	if err != nil {
		return errors.Wrap(err, "faild to decode")
	}

	err = cfg.Servers.validate()
	if err != nil {
		return errors.Wrap(err, "invalid server configuration")
	}

	return nil
}
