//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/drivers/gpio"
	"github.com/szampardi/gobot/platforms/dragonboard"
)

func main() {
	dragonAdaptor := dragonboard.NewAdaptor()
	button := gpio.NewButtonDriver(dragonAdaptor, "GPIO_A")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pressed")
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button released")
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{dragonAdaptor},
		[]gobot.Device{button},
		work,
	)

	robot.Start()
}
