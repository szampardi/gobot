/*
Package digispark provides the Gobot adaptor for the Digispark ATTiny-based USB development board.

Installing:

This package requires installing `libusb`.
Then you can install the package with:

	Please refer to the main [README.md](https://github.com/hybridgroup/gobot/blob/release/README.md)

Example:

	package main

	import (
		"time"

		"github.com/szampardi/gobot"
		"github.com/szampardi/gobot/drivers/gpio"
		"github.com/szampardi/gobot/platforms/digispark"
	)

	func main() {
		digisparkAdaptor := digispark.NewAdaptor()
		led := gpio.NewLedDriver(digisparkAdaptor, "0")

		work := func() {
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		}

		robot := gobot.NewRobot("blinkBot",
			[]gobot.Connection{digisparkAdaptor},
			[]gobot.Device{led},
			work,
		)

		robot.Start()
	}

For further information refer to digispark README:
https://github.com/hybridgroup/gobot/blob/master/platforms/digispark/README.md
*/
package digispark // import "github.com/szampardi/gobot/platforms/digispark"
