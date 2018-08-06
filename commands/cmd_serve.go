package commands

import (
	"crypto/tls"
	"path/filepath"

	"github.com/hlts2/go-LB/config"
	"github.com/hlts2/go-LB/server"
	"github.com/urfave/cli"
)

const (

	// TSLCertFile is cert file name of TSL
	TSLCertFile = "cert.pen"

	// TSLKeyFile is key file name of TSL
	TSLKeyFile = "key.pen"
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
				certname = filepath.Join(tlspath, TSLCertFile)
				keyname  = filepath.Join(tlspath, TSLKeyFile)
			)

			cert, err := tls.LoadX509KeyPair(certname, keyname)
			if err != nil {
				return err
			}

			tlsConfig := tls.Config{
				Certificates: []tls.Certificate{
					cert,
				},
			}

			err = lb.ServeTLS(&tlsConfig, certname, keyname)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
