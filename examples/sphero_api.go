//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/api"
	"github.com/szampardi/gobot/platforms/sphero"
)

func main() {
	master := gobot.NewMaster()
	api.NewAPI(master).Start()

	spheros := map[string]string{
		"Sphero-BPO": "/dev/rfcomm0",
	}

	for name, port := range spheros {
		spheroAdaptor := sphero.NewAdaptor(port)
		spheroDriver := sphero.NewSpheroDriver(spheroAdaptor)

		work := func() {
			spheroDriver.SetRGB(uint8(255), uint8(0), uint8(0))
		}

		robot := gobot.NewRobot(name,
			[]gobot.Connection{spheroAdaptor},
			[]gobot.Device{spheroDriver},
			work,
		)
		robot.AddCommand("turn_blue", func(params map[string]interface{}) interface{} {
			spheroDriver.SetRGB(uint8(0), uint8(0), uint8(255))
			return nil
		})

		master.AddRobot(robot)
	}

	master.Start()
}
