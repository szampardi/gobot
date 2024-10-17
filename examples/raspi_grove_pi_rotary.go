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
	"github.com/szampardi/gobot/drivers/i2c"
	"github.com/szampardi/gobot/platforms/raspi"
)

func main() {
	board := raspi.NewAdaptor()
	gp := i2c.NewGrovePiDriver(board)
	sensor := aio.NewGroveRotaryDriver(gp, "A1", aio.WithSensorCyclicRead(500*time.Millisecond))

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			fmt.Println("sensor", data)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{gp, sensor},
		work,
	)

	robot.Start()
}
