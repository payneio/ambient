package smartthings

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/payneio/ambient/registry"
)

type System struct {
	client   *http.Client
	endpoint string
}

func (s *System) Authenticate(creds registry.Credentials) {

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

// IssueCommand sends a given command to an URI and returns the contents
func IssueCommand(client *http.Client, endpoint string, cmd string) ([]byte, error) {
	uri := endpoint + cmd
	resp, err := client.Get(uri)
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (s *System) ListDevices() ([]*STDevice, error) {

	var devices []*STDevice

	// List all info about devices
	devs, err := GetDevices(s.client, s.endpoint)
	if err != nil {
		return nil, err
	}
	for _, d := range devs {

		device := &STDevice{
			ID:          d.ID,
			Name:        d.Name,
			DisplayName: d.DisplayName,
			Attributes:  make(map[string]interface{}),
		}

		// Get device info
		info, err := GetDeviceInfo(s.client, s.endpoint, device.ID)
		if err != nil {
			return nil, err
		}

		for k, v := range info.Attributes {
			device.Attributes[k] = v
		}

		cmds, err := ListDeviceCommands(s.client, s.endpoint, device.ID)
		if err != nil {
			return nil, err
		}
		device.Commands = cmds

		devices = append(devices, device)
	}
	return devices, nil

}
