//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"fmt"
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/drivers/i2c"
	"github.com/szampardi/gobot/platforms/raspi"
)

const (
	ultrasonicPin = "4"
	delayMillisec = 10
)

func main() {
	r := raspi.NewAdaptor()
	gp := i2c.NewGrovePiDriver(r)

	work := func() {
		gobot.Every(1*time.Second, func() {
			if val, err := gp.UltrasonicRead(ultrasonicPin, delayMillisec); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Distance [cm]", val)
			}
		})
	}

	robot := gobot.NewRobot("ultrasonicBot",
		[]gobot.Connection{r},
		[]gobot.Device{gp},
		work,
	)

	robot.Start()
}
