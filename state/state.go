package state

import "github.com/payneio/ambient/registry"

func New(registry *registry.Registry) *Current {
	return &Current{
		registry: registry,
	}
}

type Current struct {
	registry *registry.Registry
}

func (s *Current) Refresh() {

}
