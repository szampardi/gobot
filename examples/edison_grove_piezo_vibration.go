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
	sensor := aio.NewGrovePiezoVibrationSensorDriver(board, "0", aio.WithSensorCyclicRead(500*time.Millisecond))

	work := func() {
		sensor.On(aio.Vibration, func(data interface{}) {
			fmt.Println("got one!")
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{board},
		[]gobot.Device{sensor},
		work,
	)

	robot.Start()
}
