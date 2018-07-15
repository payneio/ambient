package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func die(err error) {
	fmt.Printf(`ERROR: %v`, err)
	os.Exit(1)
}

func main() {

	configPath := os.Getenv(`AMBIENTD`)
	if configPath == "" {
		die(errors.New("AMBIENTD must be specified in environment"))
	}

	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "command",
			Usage: "Execute a command",
			Action: func(c *cli.Context) error {
				fmt.Println("TODO: run a command against ambientd")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		die(err)
	}

}
