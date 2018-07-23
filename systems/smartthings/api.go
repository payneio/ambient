package smartthings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/payneio/ambient/registry"
)

const (
	apiURI = `https://graph.api.smartthings.com`

	// Endpoints URL
	endPointsURI = apiURI + "/api/smartapps/endpoints"
)

// GetEndPointsURI returns the smartthing endpoints URI. The endpoints
// URI is the base for all app requests.
func GetEndPointsURI(client *http.Client) (string, error) {
	// Fetch the JSON containing our endpoint URI
	resp, err := client.Get(endPointsURI)
	if err != nil {
		return "", fmt.Errorf("error getting endpoints URI %q", err)
	}
	contents, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if string(contents) == "[]" {
		return "", fmt.Errorf("endpoint URI returned no content")
	}

	// Only URI is fetched from JSON string.
	var ep []endpoints
	err = json.Unmarshal(contents, &ep)
	if err != nil {
		return "", fmt.Errorf("error decoding JSON: %q", err)
	}
	return ep[0].URI, nil
}

type STDevice struct {
	ID          string
	Name        string
	DisplayName string
	Attributes  map[string]interface{}
	Commands    []STDeviceCommand
}

// STDeviceList holds the list of devices returned by /devices
type STDeviceList struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}

// STDeviceInfo holds information about a specific device.
type STDeviceInfo struct {
	STDeviceList
	Attributes map[string]interface{} `json:"attributes"`
}

// STDeviceCommand holds one command a device can accept.
type STDeviceCommand struct {
	Command string                 `json:"command"`
	Params  map[string]interface{} `json:"params"`
}

// GetDevices returns the list of devices from smartthings using
// the specified http.client and endpoint URI.
func GetDevices(client *http.Client, endpoint string) ([]STDeviceList, error) {
	ret := []STDeviceList{}

	contents, err := IssueCommand(client, endpoint, "/devices")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(contents, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// GetDeviceInfo returns device specific information about a particular device.
func GetDeviceInfo(client *http.Client, endpoint string, id string) (*STDeviceInfo, error) {
	ret := &STDeviceInfo{}

	contents, err := IssueCommand(client, endpoint, "/devices/"+id)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(contents, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// ListDeviceCommands returns a slice of commands a specific device accepts.
func ListDeviceCommands(client *http.Client, endpoint string, id string) ([]STDeviceCommand, error) {
	ret := []STDeviceCommand{}

	contents, err := IssueCommand(client, endpoint, "/devices/"+id+"/commands")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(contents, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func SendCommand(client *http.Client, endpoint string, effectorID string, command registry.Command) error {
	_, err := IssueCommand(client, endpoint, "/devices/"+effectorID+"/"+command.ID)
	return err
}
