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
	GetSensorMap() map[string]*Sensor
	GetEffectorMap() map[string]*Effector
	GetEffector(string) Effector
	GetSensor(string) Sensor
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

func (r *Registry) GetSensorMap() map[string]Sensor {
	return r.Sensors
}

func (r *Registry) GetSensor(id string) (Sensor, bool) {
	v, ok := r.Sensors[id]
	return v, ok
}

func (r *Registry) GetEffectorMap() map[string]Effector {
	return r.Effectors
}

func (r *Registry) GetEffector(id string) (Effector, bool) {
	v, ok := r.Effectors[id]
	return v, ok
}
