//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/drivers/gpio"
	"github.com/szampardi/gobot/platforms/jetson"
)

func main() {
	r := jetson.NewAdaptor()
	led := gpio.NewLedDriver(r, "40")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
