package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"

	"github.com/filecoin-project/go-address"
	locli "github.com/tddey01/switfs-auth/cli"
	"github.com/tddey01/switfs-auth/core"
)

func main() {
	address.CurrentNetwork = address.Mainnet
	app := newApp()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newApp() (app *cli.App) {
	app = &cli.App{
		Version:  core.Version,
		Commands: locli.Commands,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "config dir path",
			},
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"VENUS_AUTH_PATH"},
				Value:   "~/.switfs-auth",
			},
			&cli.StringFlag{
				Name:  "listen",
				Value: "127.0.0.1:8989",
			},
		},
	}
	return app
}
