/*
Package sphero provides the Gobot adaptor and driver for the Sphero.

Installing:

	Please refer to the main [README.md](https://github.com/hybridgroup/gobot/blob/release/README.md)

Example:

	package main

	import (
		"fmt"
		"time"

		"github.com/szampardi/gobot"
		"github.com/szampardi/gobot/platforms/sphero"
	)

	func main() {
		adaptor := sphero.NewAdaptor("/dev/rfcomm0")
		driver := sphero.NewSpheroDriver(adaptor)

		work := func() {
			gobot.Every(3*time.Second, func() {
				driver.Roll(30, uint16(gobot.Rand(360)))
			})
		}

		robot := gobot.NewRobot("sphero",
			[]gobot.Connection{adaptor},
			[]gobot.Device{driver},
			work,
		)

		robot.Start()
	}

For further information refer to sphero readme:
https://github.com/hybridgroup/gobot/blob/master/platforms/sphero/README.md
*/
package sphero // import "github.com/szampardi/gobot/platforms/sphero"
