package cmd

import (
	"github.com/urfave/cli"
)

// Serve is the command that serve load balancer
func Serve() cli.Command {
	return cli.Command{
		Name:  "serve",
		Usage: "serve load balancer",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "set, s",
				Value: "config.yml",
				Usage: "set the configuration file",
			},
		},
		Action: func(c *cli.Context) error {
			// var cfg config.Config
			// err := config.Load(c.String("set"), &cfg)
			// if err != nil {
			// 	return errors.Wrap(err, "faild to load configuration file")
			// }
			//
			// lb := server.NewLB(c.String("host") + ":" + c.String("port")).Build(cfg)
			//
			// tlspath := c.String("tlspath")
			//
			// // NOT TLS Mode
			// if tlspath == "" {
			// 	err := lb.Serve()
			// 	if err != nil {
			// 		return errors.Wrap(err, "faild to run server")
			// 	}
			// 	return nil
			// }
			//
			// var (
			// 	certpath = filepath.Join(tlspath, TLSCertFileName)
			// 	keypath  = filepath.Join(tlspath, TLSKeyFileName)
			// )
			//
			// cert, err := tls.LoadX509KeyPair(certpath, keypath)
			// if err != nil {
			// 	return errors.Wrap(err, "faild to load certification file and key file")
			// }
			//
			// tlsConfig := tls.Config{
			// 	Certificates: []tls.Certificate{
			// 		cert,
			// 	},
			// }
			//
			// err = lb.ServeTLS(&tlsConfig, certpath, keypath)
			// if err != nil {
			// 	return errors.Wrap(err, "faild to run tls server")
			// }

			return nil
		},
	}
}
