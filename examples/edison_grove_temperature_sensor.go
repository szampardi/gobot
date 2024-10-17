//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/drivers/aio"
	"github.com/szampardi/gobot/platforms/intel-iot/edison"
)

func main() {
	board := edison.NewAdaptor()
	sensor := aio.NewGroveTemperatureSensorDriver(board, "0", aio.WithSensorCyclicRead(500*time.Millisecond))

	work := func() {
		gobot.Every(500*time.Millisecond, func() {
			fmt.Println("current temp (c): ", sensor.Temperature())
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{sensor},
		work,
	)

	robot.Start()
}
