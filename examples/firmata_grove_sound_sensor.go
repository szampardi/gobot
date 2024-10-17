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
	"github.com/szampardi/gobot/drivers/aio"
	"github.com/szampardi/gobot/platforms/firmata"
)

func main() {
	board := firmata.NewAdaptor(os.Args[1])
	sensor := aio.NewGroveSoundSensorDriver(board, "3", aio.WithSensorCyclicRead(500*time.Millisecond))

	work := func() {
		sensor.On(aio.Data, func(data interface{}) {
			fmt.Println("sensor", data)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{board},
		[]gobot.Device{sensor},
		work,
	)

	robot.Start()
}
