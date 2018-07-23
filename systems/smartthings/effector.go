package smartthings

import (
	"fmt"

	"github.com/payneio/ambient/registry"
)

type STEffector struct {
	ID       string
	DeviceID string
	Commands []STDeviceCommand
	System   *System
}

func (e *STEffector) GetID() string {
	return e.ID
}

// ListCommands returns a map of registry commands for this effector
// by transforming the underlying STDeviceCommands
func (e *STEffector) ListCommands() map[string]*registry.Command {
	cmds := make(map[string]*registry.Command)
	for _, cmd := range e.Commands {
		regCmd := &registry.Command{
			ID:     cmd.Command,
			Params: cmd.Params,
		}
		cmds[regCmd.ID] = regCmd
	}
	return cmds
}

func (e *STEffector) Exec(cmd registry.Command) {
	err := SendCommand(e.System.client, e.System.endpoint, e.GetID(), cmd)
	if err != nil {
		// TODO: handle error conditions for consistency
		fmt.Println(err)
	}
}

type Command struct {
	Label  string
	Params map[string]interface{}
}
