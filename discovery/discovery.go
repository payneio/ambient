package discovery

import (
	"fmt"

	"github.com/payneio/ambient"
	"github.com/payneio/ambient/registry"
	"github.com/payneio/ambient/systems/smartthings"
)

func Discover(config ambient.Config, reg *registry.Registry) {

	stSystem := &smartthings.System{}
	stSystem.Authenticate(registry.Credentials{
		ClientID:      config.SmartThings.ClientID,
		Secret:        config.SmartThings.Secret,
		TokenFilePath: config.SmartThings.TokenFilePath,
	})
	devices, err := stSystem.ListDevices()
	if err != nil {
		// TODO: need to put retries and other fault tolerance here
		fmt.Println(err)
	}
	for _, device := range devices {
		if device.Commands != nil {
			effector := &smartthings.STEffector{
				ID:       device.ID,
				DeviceID: device.ID,
				Commands: device.Commands,
				System:   stSystem,
			}
			reg.RegisterEffector(effector)
		}
		for k, v := range device.Attributes {
			sensor := &smartthings.STSensor{
				ID:    fmt.Sprintf(`%v:%v`, device.ID, k),
				Value: v,
			}
			reg.RegisterSensor(sensor)
		}
	}

	// TODO: need to register WHAT and HOW to read sensor data / state
	// and how to map this data to the state tree

	// TODO: map effectors and what state variables they effect and how
	/*

		workbench-light
		  effectors
			toggle
		  state
			home.workshop.light.workbench: on/off
		main-light
		  effectors
			toggle
		  state
			home.workship.light.overhead: on/off
		workshop-motion
		  state
			home.workshop.motion.3m: yes/no
		workshop-door
		  state
			home.workshop.door: open/closed
		workshop-door-temp
		  device-attr: 40f1370c-c65a-4073-ba52-eec6544f4900:temperature
		  state
			home.workship.temperature.door
		workshop-wall-temp
		  state
		    home.workshop.temperature.wall
		workshop-win-console
		  state
		    home.workshop.console.win: sleeping/awake
	*/

	/*
		An example state tree looks like:
		home:
		  workshop:
			light:
			  workbench: on
			  overhead: on
			motion:
			  3m: yes
			door: closed
			temp:
			  door: 73
			  wall: 78
			console:
			  win: awake
			  ubuntu: awake
	*/

}
