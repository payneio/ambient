package smartthings

import "strconv"

type Sensor interface {
	GetID() string
	Read() interface{}
	ReadString() string
	ReadBytes() []byte
	ReadInt() int
}

type STSensor struct {
	ID    string
	Value interface{}
}

func (s *STSensor) GetID() string {
	return s.ID
}

func (s *STSensor) Read() interface{} {
	return s.Value
}

func (s *STSensor) ReadString() string {
	switch s.Value.(type) {
	case string:
		return s.Value.(string)
	case int:
		return strconv.Itoa(s.Value.(int))
	default:
		return ""
	}
}

func (s *STSensor) ReadBytes() []byte {
	return []byte(s.ReadString())
}

func (s *STSensor) ReadInt() int {
	switch s.Value.(type) {
	case int:
		return s.Value.(int)
	case string:
		if i, err := strconv.Atoi(s.Value.(string)); err == nil {
			return i
		}
		return 0
	default:
		return 0
	}
}
