//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/drivers/gpio"
	"github.com/szampardi/gobot/drivers/i2c"
	"github.com/szampardi/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	gp := i2c.NewGrovePiDriver(r)
	button := gpio.NewButtonDriver(gp, "D3", gpio.WithButtonPollInterval(50*time.Millisecond))
	led := gpio.NewLedDriver(gp, "D2")

	work := func() {
		button.On(gpio.ButtonPush, func(data interface{}) {
			fmt.Println("button pressed")
			led.On()
		})

		button.On(gpio.ButtonRelease, func(data interface{}) {
			fmt.Println("button released")
			led.Off()
		})
	}

	robot := gobot.NewRobot("buttonBot",
		[]gobot.Connection{r},
		[]gobot.Device{gp, button, led},
		work,
	)

	robot.Start()
}
