package slb

import (
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

	return nil
}

// Balancing is custom type for balancing algorithm name.
type Balancing string

// CreateHandler returns balancer.Handler implementation.
// If set invalid balancing algorithm, the default balancing algorithm(round-robin) is used.
func (b Balancing) CreateHandler(urls []*url.URL, proxier balancer.Proxier) balancer.Handler {
	switch b {
	case RoundRobin:
		return roundrobin.New(urls, proxier)
	case IPHash:
		return iphash.New(urls, proxier)
	case LeastConnections:
		return leastconnections.New(urls, proxier)
	default:
		return roundrobin.New(urls, proxier)
	}
}

// ServerConfig represents configuration content for server.
type ServerConfig struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	url    *url.URL
}

func (sc ServerConfig) String() string {
	return sc.Scheme + "://" + sc.Host + ":" + sc.Port
}

func (sc ServerConfig) validate() error {
	if len(sc.Scheme) == 0 || len(sc.Port) == 0 {
		return errors.Errorf("missing protocol scheme or port")
	}
	return nil
}

// ServerConfigs represents ServerConfig slice.
type ServerConfigs []ServerConfig

func (scs ServerConfigs) validate() error {
	if len(scs) == 0 {
		return errors.New("backend servers dose not exist")
	}

	hostports := make([]string, len(scs))

	for i, sc := range scs {
		err := sc.validate()
		if err != nil {
			return errors.Wrap(err, "invalid server configuration")
		}

		addr := sc.String()

		scs[i].url, err = url.ParseRequestURI(addr)
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

// getURLs returns url of servers.
func (scs ServerConfigs) getURLs() []*url.URL {
	urls := make([]*url.URL, len(scs))

	for i, sc := range scs {
		urls[i] = sc.url
	}
	return urls
}

// TLSConfig represents the TLS configuration for server load balancer.
type TLSConfig struct {
	// Enable represents the server load balancer enable TLS or not.
	Enabled bool `yaml:"enabled"`

	// CertKey represents the environment variable name having the certificate file path of server load balancer.
	CertKey string `yaml:"cert_key"`

	// KeyKey represents the environment variable name having the private key file path of server load balancer certificate.
	KeyKey string `yaml:"key_key"`
}

// Config represents an application configuration content (config.yaml).
type Config struct {
	ServerConfig         `yaml:",inline"`
	TLSConfig            `yaml:"tls"`
	Balancing            Balancing     `yaml:"balancing"`
	BackendServerConfigs ServerConfigs `yaml:"servers"`
}

// Validate validates configuration content(*Config).
func (c *Config) validate() error {
	err := c.ServerConfig.validate()
	if err != nil {
		return errors.Wrap(err, "invalid server configuration")
	}

	err = c.BackendServerConfigs.validate()
	if err != nil {
		return errors.Wrap(err, "invalid backend servers configuration")
	}

	return nil
}
