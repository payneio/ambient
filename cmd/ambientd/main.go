package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
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
				return boot(config)
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		die(err)
	}

}

func boot(config ambient.Config) error {
	registry := &registry.Registry{}

	discovery.Discover(config, registry)

	currentState := state.New(registry)
	fmt.Print(currentState)

	// Start HTTP
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/state", func(c *gin.Context) {
		c.JSON(200, currentState)
	})

	r.POST("/command", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	r.Run() // listen and serve on 0.0.0.0:8080

	return nil
}
