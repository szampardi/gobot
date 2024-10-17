//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/drivers/i2c"
	"github.com/szampardi/gobot/platforms/intel-iot/edison"
)

func main() {
	board := edison.NewAdaptor()
	screen := i2c.NewGroveLcdDriver(board)

	work := func() {
		screen.Write("hello")

		screen.SetRGB(255, 0, 0)

		gobot.After(5*time.Second, func() {
			screen.Clear()
			screen.Home()
			screen.SetRGB(0, 255, 0)
			// set a custom character in the first position
			screen.SetCustomChar(0, i2c.CustomLCDChars["smiley"])
			// add the custom character at the end of the string
			screen.Write("goodbye\nhave a nice day " + string(byte(0)))
			gobot.Every(500*time.Millisecond, func() {
				screen.Scroll(false)
			})
		})

		screen.Home()
		time.Sleep(1 * time.Second)
		screen.SetRGB(0, 0, 255)
	}

	robot := gobot.NewRobot("screenBot",
		[]gobot.Connection{board},
		[]gobot.Device{screen},
		work,
	)

	robot.Start()
}
