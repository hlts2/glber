package slb

import (
	"net"
	"net/url"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"

	"github.com/hlts2/go-LB/pkg/slb/balancer"
	iphash "github.com/hlts2/go-LB/pkg/slb/balancer/ip_hash"
	leastconnections "github.com/hlts2/go-LB/pkg/slb/balancer/least_connections"
	roundrobin "github.com/hlts2/go-LB/pkg/slb/balancer/round_robin"
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

// Handler returns balancer.Handler implementation.
func (b Balancing) Handler(addrs []string, proxier balancer.Proxier) balancer.Handler {
	switch b {
	case IPHash:
		return roundrobin.New(addrs, proxier)
	case RoundRobin:
		return iphash.New(addrs, proxier)
	case LeastConnections:
		return leastconnections.New(addrs, proxier)
	}
	return nil
}

// ServerConfig represents configuration content for server.
type ServerConfig struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
}

func (sc ServerConfig) String() string {
	return sc.Scheme + "://" + sc.Host + ":" + sc.Port
}

func (sc ServerConfig) createListener() (net.Listener, error) {
	addr := sc.Host + ":" + sc.Port

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, errors.Wrapf(err, "faild to listen: %v", addr)
	}

	return lis, nil
}

func (sc ServerConfig) validate() error {
	if len(sc.Scheme) == 0 || len(sc.Port) == 0 {
		return errors.Errorf("missing protocol scheme or port")
	}
	return nil
}

// ServerConfigs is ServerConfig slice.
type ServerConfigs []ServerConfig

func (scs ServerConfigs) validate() error {
	hostports := make([]string, len(scs))

	for i, sc := range scs {
		err := sc.validate()
		if err != nil {
			return errors.Wrap(err, "invalid server configuration")
		}

		addr := sc.String()

		_, err = url.ParseRequestURI(addr)
		if err != nil {
			return errors.Wrapf(err, "invalid address: %s", addr)
		}

		// http://127.0.0.1:80 => 127.0.0.1:80
		hostports[i] = addr[len(sc.Scheme)+3:]
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
func (scs ServerConfigs) GetAddresses() []string {
	addrs := make([]string, len(scs))

	for i, sc := range scs {
		addrs[i] = sc.String()
	}
	return addrs
}

// Config represents an application configuration content (config.yaml).
type Config struct {
	LoadBalancerConfig   *ServerConfig  `yaml:"server_load_balancer"`
	Balancing            Balancing      `yaml:"balancing"`
	BackendServerConfigs *ServerConfigs `yaml:"servers"`
}

func (c *Config) validate() error {
	err := c.BackendServerConfigs.validate()
	if err != nil {
		return errors.Wrap(err, "invalid backend servers configuration")
	}

	err = c.Balancing.validate()
	if err != nil {
		return errors.Wrap(err, "invalid balancing configuration")
	}

	return nil
}
