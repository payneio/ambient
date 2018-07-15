package discovery

import (
	"github.com/payneio/ambient"
	"github.com/payneio/ambient/registry"
	"github.com/payneio/ambient/systems/smartthings"
)

func Discover(config ambient.Config, registry *registry.Registry) {

	smartthings := smartthings.System{}
	smartthings.Authenticate(ambient.Credentials{
		ClientID:      config.SmartThings.ClientID,
		Secret:        config.SmartThings.Secret,
		TokenFilePath: config.SmartThings.TokenFilePath,
	})
	smartthings.RegisterDevices()

}
