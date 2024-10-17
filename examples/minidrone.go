//go:build example
// +build example

//
// Do not build by default.

/*
 How to run
 Pass the Bluetooth name or address as first param:

	go run examples/minidrone.go "Mambo_1234"

 NOTE: sudo is required to use BLE in Linux
*/

package main

import (
	"os"
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/platforms/ble"
	"github.com/szampardi/gobot/platforms/parrot/minidrone"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	drone := minidrone.NewDriver(bleAdaptor)

	work := func() {
		drone.TakeOff()

		gobot.After(5*time.Second, func() {
			drone.Land()
		})
	}

	robot := gobot.NewRobot("minidrone",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}
