package ambient

type Authenticator interface {
	Authenticate()
}

type Registrar interface {
	RegisterSystem()
	RegisterSensor()
	RegisterEffector()
}

type Command struct {
	Label  string
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
	RegisterDevices(*Registrar)
}
