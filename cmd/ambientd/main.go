package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/payneio/ambient"
	"github.com/payneio/ambient/discovery"
	"github.com/payneio/ambient/registry"
	"github.com/payneio/ambient/state"
	"github.com/urfave/cli"
)

func die(err error) {
	fmt.Printf(`ERROR: %v`, err)
	os.Exit(1)
}

func main() {

	// Load config
	configPath := os.Getenv(`AMBIENT_CONFIG`)
	if configPath == "" {
		die(errors.New("AMBIENT_CONFIG must be specified in environment"))
	}
	config, err := ambient.LoadConfig(configPath)
	if err != nil {
		die(err)
	}

	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start ambient",
			Action: func(c *cli.Context) error {

				registry := &registry.Registry{}

				discovery.Discover(config, registry)

				currentState := state.New(registry)

				fmt.Print(currentState)

				return nil

			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		die(err)
	}

}
