package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/payneio/ambient"
	"github.com/payneio/ambient/discovery"
)

func die(err error) {
	fmt.Printf(`ERROR: %v`, err)
	os.Exit(-1)
}

func main() {
	fmt.Println(`Hello`)

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
}
