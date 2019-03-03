package main

import (
	"os"

	"github.com/kpango/glg"
	"github.com/pkg/errors"
	"github.com/urfave/cli"

	"github.com/hlts2/go-LB/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-LB"
	app.Usage = "Load Balancer"
	app.Version = "v0.0.1"
	app.Commands = cli.Commands{
		cmd.ServeCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		glg.Fatalln(errors.Wrap(err, "exit app because an error occurred"))
	}
}
