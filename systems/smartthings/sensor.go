package smartthings

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
