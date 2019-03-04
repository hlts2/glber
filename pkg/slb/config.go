package slb

import (
	"net/url"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// Represents name of balancing algorithm.
const (
	IPHash           = "ip-hash"
	RoundRobin       = "round-robin"
	LeastConnections = "least-connections"
)

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

	err = cfg.validate()
	if err != nil {
		return errors.Wrap(err, "invalid configuration")
	}

	return nil
}

// Balancing is custom type for balancing algorithm name.
type Balancing string

func (b Balancing) validate() error {
	switch b {
	case IPHash, RoundRobin, LeastConnections:
		return nil
	default:
		return errors.Errorf("invalid balancing algorithm: %s", b)
	}
}

// ServerConfig represents configuration content for server.
type ServerConfig struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
}

func (s ServerConfig) String() string {
	return s.Scheme + "://" + s.Host + ":" + s.Port
}

// ServersConfig is Server slice.
type ServersConfig []ServerConfig

func (ss ServersConfig) validate() error {
	hostports := make([]string, len(ss))

	for i, s := range ss {
		if len(s.Scheme) == 0 || len(s.Host) == 0 || len(s.Port) == 0 {
			return errors.Errorf("empty scheme: %v or host: %v or port: %v", s.Scheme, s.Host, s.Port)
		}

		addr := s.String()

		_, err := url.ParseRequestURI(addr)
		if err != nil {
			return errors.Wrapf(err, "invalid address: %s", addr)
		}

		// http://127.0.0.1:80 => 127.0.0.1:80
		hostports[i] = addr[len(s.Scheme)+3:]
	}

	ok := duplicateExists(hostports)
	if ok {
		return errors.New("duplicate host and port exists")
	}

	return nil
}

// duplicateExists returns true if there is duplicte in to values.
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
func (ss ServersConfig) GetAddresses() []string {
	addrs := make([]string, len(ss))

	for i, s := range ss {
		addrs[i] = s.String()
	}
	return addrs
}

// Config represents an application configuration content (config.yaml).
type Config struct {
	ServersConfig ServersConfig `yaml:"servers"`
	Balancing     Balancing     `yaml:"balancing"`
}

func (c *Config) validate() error {
	err := c.ServersConfig.validate()
	if err != nil {
		return errors.Wrap(err, "invalid server configuration")
	}

	err = c.Balancing.validate()
	if err != nil {
		return errors.Wrap(err, "invalid balancing configuration")
	}

	return nil
}
