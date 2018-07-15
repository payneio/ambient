package registry

type Registrar interface {
	RegisterSystem()
	RegisterSensor()
	RegisterEffector()
}

type Registry struct {
}

func (r *Registry) RegisterSystem() {

}

func (r *Registry) RegisterSensor() {

}

func (r *Registry) RegisterEffector() {

}
