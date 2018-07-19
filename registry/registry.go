package registry

type Authenticator interface {
	Authenticate()
}

type Command struct {
	ID     string
	Params map[string]interface{}
}

type Effector interface {
	GetID() string
	ListCommands() map[string]*Command
	Exec(Command)
}

type Sensor interface {
	GetID() string
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

type Registrar interface {
	RegisterSystem()
	RegisterSensor()
	RegisterEffector()
}

func New() *Registry {
	return &Registry{
		Sensors:   make(map[string]Sensor),
		Effectors: make(map[string]Effector),
	}
}

type Registry struct {
	Sensors   map[string]Sensor
	Effectors map[string]Effector
}

func (r *Registry) RegisterSensor(sensor Sensor) {
	r.Sensors[sensor.GetID()] = sensor
}

func (r *Registry) RegisterEffector(effector Effector) {
	r.Effectors[effector.GetID()] = effector
}
