//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/drivers/gpio"
	"github.com/szampardi/gobot/platforms/beaglebone"
)

func main() {
	beagleboneAdaptor := beaglebone.NewAdaptor()
	led := gpio.NewDirectPinDriver(beagleboneAdaptor, "P8_10")
	button := gpio.NewDirectPinDriver(beagleboneAdaptor, "P8_09")

	work := func() {
		gobot.Every(500*time.Millisecond, func() {
			val, _ := button.DigitalRead()
			if val == 1 {
				led.DigitalWrite(1)
			} else {
				led.DigitalWrite(0)
			}
		})
	}

	robot := gobot.NewRobot("pinBot",
		[]gobot.Connection{beagleboneAdaptor},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
