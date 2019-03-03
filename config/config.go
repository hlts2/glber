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
	for _, s := range ss {
		if len(s.Scheme) == 0 || len(s.Host) == 0 || len(s.Port) == 0 {
			return errors.New("empty scheme or host or port")
		}

		_, err := url.ParseRequestURI(s.address())
		if err != nil {
			return nil
		}
	}
	return nil
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

	addrs := cfg.Servers.GetAddresses()

	ok := duplicateAddressExists(addrs)
	if ok {
		return errors.New("duplicate host in yaml file")
	}

	return nil
}

// duplicateAddressExists returns true if there is duplicte address.
func duplicateAddressExists(hosts []string) bool {
	m := make(map[string]bool, len(hosts))

	for _, host := range hosts {
		if _, ok := m[host]; ok {
			return true
		}
		m[host] = true
	}
	return false
}
