package commands

import (
	"crypto/tls"
	"path/filepath"

	"github.com/hlts2/go-LB/config"
	"github.com/hlts2/go-LB/server"
	"github.com/urfave/cli"
)

const (

	// TLSCertFile is cert file name for TLS
	TLSCertFile = "cert.pem"

	// TLSKeyFile is key file name for TLS
	TLSKeyFile = "key.pem"
)

// ServeCommand is the command that serve load balancer
func ServeCommand() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "serve load balancer",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "set, s",
				Value: "config.yml",
				Usage: "set a config file of load balancer",
			},
			cli.StringFlag{
				Name:  "host, H",
				Value: "127.0.0.1",
				Usage: "set a host name or IP of load balancer",
			},
			cli.StringFlag{
				Name:  "port, p",
				Value: "8080",
				Usage: "set a port number of load balancer",
			},
			cli.StringFlag{
				Name:  "tlspath",
				Value: "",
				Usage: "set a TLS directory of load balancer",
			},
		},
		Action: func(c *cli.Context) error {
			var conf config.Config
			err := config.LoadConfig(c.String("set"), &conf)
			if err != nil {
				return err
			}

			lb := server.NewLB(c.String("host") + ":" + c.String("port")).Build(conf)

			tlspath := c.String("tlspath")

			// NOT TLS Mode
			if tlspath == "" {
				err := lb.Serve()
				if err != nil {
					return err
				}
				return nil
			}

			var (
				certpath = filepath.Join(tlspath, TLSCertFile)
				keypath  = filepath.Join(tlspath, TLSKeyFile)
			)

			cert, err := tls.LoadX509KeyPair(certpath, keypath)
			if err != nil {
				return err
			}

			tlsConfig := tls.Config{
				Certificates: []tls.Certificate{
					cert,
				},
			}

			err = lb.ServeTLS(&tlsConfig, certpath, keypath)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
