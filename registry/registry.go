package registry

type Authenticator interface {
	Authenticate()
}

type Command struct {
	Label  string
	Params map[string]interface{}
}

type Effector interface {
	Commands() map[string]Command
	Exec(Command)
}

type Sensor interface {
	Label() string
	Read() interface{}
	ReadString() string
	ReadBytes() []byte
	ReadInt() int
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

type Registrar interface {
	RegisterSystem()
	RegisterSensor()
	RegisterEffector()
}

func New() *Registry {
	return &Registry{
		Sensors:   make(map[string]*Sensor),
		Effectors: make(map[string]*Effector),
	}
}

type Registry struct {
	Sensors   map[string]*Sensor
	Effectors map[string]*Effector
}

func (r *Registry) RegisterSensor() {

}

func (r *Registry) RegisterEffector() {

}
