package state

import (
	"github.com/hashicorp/go-immutable-radix"
	"github.com/payneio/ambient/registry"
)

func New(registry *registry.Registry) *Current {
	return &Current{
		registry: registry,
		root:     iradix.New(),
	}
}

type Current struct {
	registry *registry.Registry
	root     *iradix.Tree
}

func (c *Current) Update(k string, v interface{}) *iradix.Tree {
	t, _, _ := c.root.Insert([]byte(k), v)
	c.root = t
	return c.root
}

func (c *Current) Get(k string) (interface{}, bool) {
	return c.root.Get([]byte(k))
}

// Refresh builds root state from all registered sensors
func (s *Current) Refresh() {

}
