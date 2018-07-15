package ambient

import "github.com/payneio/ambient/registry"

type Authenticator interface {
	Authenticate()
}

type Command struct {
	URI    string
	Params map[string]interface{}
}

type Effector interface {
	Commands() []Command
	Exec(Command)
}

type Credentials struct {
	ClientID      string
	Secret        string
	TokenFilePath string
}

type System interface {
	Authenticate(Credentials)
	RegisterDevices(*registry.Registrar)
}

type Device struct {
	ID          string
	Name        string
	DisplayName string
	Attributes  map[string]interface{}
	Commands    []Command
}

type Switcher interface {
	On()
	Off()
}
