package smartthings

import "github.com/payneio/ambient/registry"

type STEffector struct {
	ID       string
	DeviceID string
	Commands []*STDeviceCommand
	System   *System
}

func (e *STEffector) GetID() string {
	return e.ID
}

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
	err := SendCommand(e.System.client, e.System.endpoint, cmd)
	if err != nil {
		// TODO: handle error conditions for consistency
	}
}

type Command struct {
	Label  string
	Params map[string]interface{}
}

type Effector interface {
	ID() string
	Commands() map[string]Command
	Exec(Command)
}

type Sensor interface {
	ID() string
	Read() interface{}
	ReadString() string
	ReadBytes() []byte
	ReadInt() int
}
