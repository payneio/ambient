package smartthings

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/payneio/ambient"
)

type System struct {
	client   *http.Client
	endpoint string
}

func (s *System) Authenticate(creds ambient.Credentials) {

	// Create the oauth2.config object and get a token
	config := NewOAuthConfig(creds.ClientID, creds.Secret)
	token, err := GetToken(creds.TokenFilePath, config)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a client with the token. This client will be used for all ST
	// API operations from here on.
	ctx := context.Background()
	s.client = config.Client(ctx, token)

	// Retrieve Endpoints URI. All future accesses to the smartthings API
	// for this session should use this URL, followed by the desired URL path.
	endpoint, err := GetEndPointsURI(s.client)
	if err != nil {
		log.Fatalln(err)
	}
	s.endpoint = endpoint

}

func (s *System) RegisterDevices() {

	devices := []string{}

	// List all info about devices
	devs, err := GetDevices(s.client, s.endpoint)
	if err != nil {
		log.Fatalln(err)
	}
	for _, d := range devs {
		devices = append(devices, d.ID)
	}

	for _, id := range devices {
		dev, err := GetDeviceInfo(s.client, s.endpoint, id)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("\nDevice ID:      %s\n", dev.ID)
		fmt.Printf("  Name:         %s\n", dev.Name)
		fmt.Printf("  Display Name: %s\n", dev.DisplayName)
		fmt.Printf("  Attributes:\n")
		for k, v := range dev.Attributes {
			fmt.Printf("    %v: %v\n", k, v)
		}

		fmt.Printf("  Commands & Parameters:\n")
		cmds, err := GetDeviceCommands(s.client, s.endpoint, id)
		for _, cmd := range cmds {
			fmt.Printf("    %s", cmd.Command)
			if len(cmd.Params) != 0 {
				fmt.Printf(" Parameters:")
				for k, v := range cmd.Params {
					fmt.Printf(" %s=%s", k, v)
				}
			}
			fmt.Println()
		}
	}

}
