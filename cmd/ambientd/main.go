package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/payneio/ambient"
	"github.com/payneio/ambient/discovery"
	"github.com/urfave/cli"
)

func die(err error) {
	fmt.Printf(`ERROR: %v`, err)
	os.Exit(1)
}

func main() {

	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start ambient",
			Action: func(c *cli.Context) error {

				// TODO: replace this with env var
				configPath := os.Getenv(`AMBIENT_CONFIG`)
				if configPath == "" {
					die(errors.New("AMBIENT_CONFIG must be specified in environment"))
				}
				config, err := ambient.LoadConfig(configPath)
				if err != nil {
					die(err)
				}

				discovery.Discover(config)
				return nil

			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		die(err)
	}

}
