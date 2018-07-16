package smartthings

import (
	"encoding/json"
	"net/http"
)

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

// GetDeviceCommands returns a slice of commands a specific device accepts.
func GetDeviceCommands(client *http.Client, endpoint string, id string) ([]STDeviceCommand, error) {
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
