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
	"github.com/szampardi/gobot/platforms/dexter/gopigo3"
	"github.com/szampardi/gobot/platforms/raspi"
)

func main() {
	raspiAdaptor := raspi.NewAdaptor()
	gpg3 := gopigo3.NewDriver(raspiAdaptor)
	sensor := aio.NewGroveLightSensorDriver(gpg3, "AD_1_1", aio.WithSensorCyclicRead(500*time.Millisecond))

	work := func() {
		sensor.On(sensor.Event("data"), func(data interface{}) {
			fmt.Println("sensor", data)
		})
	}

	robot := gobot.NewRobot("gopigo3sensor",
		[]gobot.Connection{raspiAdaptor},
		[]gobot.Device{gpg3, sensor},
		work,
	)

	robot.Start()
}
