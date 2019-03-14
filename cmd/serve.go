package cmd

import (
	"github.com/kpango/glg"
	"github.com/pkg/errors"
	"github.com/urfave/cli"

	"github.com/hlts2/glber/pkg/slb"
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
			var cfg slb.Config
			err := slb.Load(c.String("set"), &cfg)
			if err != nil {
				return errors.Wrap(err, "faild to load configuration file")
			}

			s, err := slb.CreateSLB(&cfg)
			if err != nil {
				return errors.Wrap(err, "faild to create server load balancer")
			}

			glg.Infof("Starting Server Load Balancer on %s", cfg.Host+":"+cfg.Port)
			err = s.ListenAndServe()
			if err != nil {
				return errors.Wrap(err, "faild to listen and serve")
			}

			return nil
		},
	}
}
