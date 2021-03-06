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
	reg := registry.New()

	discovery.Discover(config, reg)

	currentState := state.New(reg)
	fmt.Print(currentState)

	// Start HTTP
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/state", func(c *gin.Context) {

		// This isn't really state. This is just printing out
		// the registry temorarily for testing
		// This will be replaced with actual state after
		// I develop the polling/update mechanism
		c.JSON(200, reg)
	})

	r.POST("/effector/:id/command", func(c *gin.Context) {

		// This is a temporary testing endpoint for manually
		// controlling devices

		// Get effector by ID
		id := c.Param(`id`)
		if id == "" {
			c.JSON(400, `id required`)
			return
		}
		effector, ok := reg.GetEffector(id)
		if !ok {
			c.JSON(404, `id not found`)
			return
		}

		// Unmarshal command from POST data
		var cmd registry.Command
		err := c.BindJSON(&cmd)
		if err != nil {
			c.JSON(400, `invalid body`)
		}

		// Execute command
		effector.Exec(cmd)

		c.JSON(200, "command executed")
	})

	r.POST("/desire", func(c *gin.Context) {

		// ambient needs to know how to accomplish a required state.
		// This means each state variable needs to be associated with
		// particular effectors, meaning "change this effector to change
		// this state". This is a simple 1:1 with things like switches,
		// though we still need to worry about non-compliance. Other
		// effectors (e.g. heaters) have a more control feedback
		// relationship with the desired state (the thermostat). Also,
		// oftentimes effectors have impacts on multiple state variables,
		// which leads to planning and multi-variable optimization.

		// Discrete:
		// body: [{
		//   variable: <state_variable>
		//   equals: <value>
		// }]

		// Continuous:
		// [{
		//   variable: <state_variable>
		//   gt: 65,
		//   lt: 75
		// }]

		// Return value should be whether or not desired state was
		// received. An optional callback may be sent which will
		// be used to notify when the state has been reached.

		// TODO: load params
		// TODO: find command from registry
		// TODO: execute command with params
		// TODO: return results

		c.JSON(200, gin.H{})
	})

	r.Run() // listen and serve on 0.0.0.0:8080

	return nil
}
