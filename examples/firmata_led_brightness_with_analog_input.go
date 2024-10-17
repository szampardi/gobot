//go:build example
// +build example

//
// Do not build by default.

/*
 How to run
 Pass serial port to use as the first param:

	go run examples/firmata_led_brightness_with_analog_input.go /dev/ttyACM0
*/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/drivers/aio"
	"github.com/szampardi/gobot/drivers/gpio"
	"github.com/szampardi/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor(os.Args[1])
	sensor := aio.NewAnalogSensorDriver(firmataAdaptor, "0", aio.WithSensorCyclicRead(500*time.Millisecond))
	led := gpio.NewLedDriver(firmataAdaptor, "3")

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			brightness := uint8(
				gobot.ToScale(gobot.FromScale(float64(data.(int)), 0, 1024), 0, 255),
			)
			fmt.Println("sensor", data)
			fmt.Println("brightness", brightness)
			led.Brightness(brightness)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{sensor, led},
		work,
	)

	robot.Start()
}
