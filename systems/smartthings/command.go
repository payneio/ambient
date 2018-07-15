package smartthings

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// DeviceCommand holds one command a device can accept.
type DeviceCommand struct {
	Command string                 `json:"command"`
	Params  map[string]interface{} `json:"params"`
}

// GetDeviceCommands returns a slice of commands a specific device accepts.
func GetDeviceCommands(client *http.Client, endpoint string, id string) ([]DeviceCommand, error) {
	ret := []DeviceCommand{}

	contents, err := IssueCommand(client, endpoint, "/devices/"+id+"/commands")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(contents, &ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// issueCommand sends a given command to an URI and returns the contents
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
