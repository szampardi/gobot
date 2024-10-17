//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/drivers/i2c"
	"github.com/szampardi/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	bme280 := i2c.NewBME280Driver(firmataAdaptor)

	work := func() {
		gobot.Every(1*time.Second, func() {
			t, _ := bme280.Temperature()
			fmt.Println("Temperature", t)

			p, _ := bme280.Pressure()
			fmt.Println("Pressure", p)

			a, _ := bme280.Altitude()
			fmt.Println("Altitude", a)

			h, _ := bme280.Humidity()
			fmt.Println("Humidity", h)
		})
	}

	robot := gobot.NewRobot("bme280bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{bme280},
		work,
	)

	robot.Start()
}
