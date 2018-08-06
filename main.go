package main

import (
	"os"

	"github.com/hlts2/go-LB/commands"
	"github.com/kpango/glg"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-LB"
	app.Usage = "Load Balancer"
	app.Version = "v0.0.1"
	app.Commands = cli.Commands{
		commands.ServeCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		glg.Fatalln(err)
	}
}
