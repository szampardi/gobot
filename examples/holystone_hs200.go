//go:build example
// +build example

//
// Do not build by default.

/*
 How to run
	go run examples/holystone_hs200.go

*/

package main

import (
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/platforms/holystone/hs200"
)

func main() {
	drone := hs200.NewDriver("172.16.10.1:8888", "172.16.10.1:8080")

	work := func() {
		drone.TakeOff()

		gobot.After(5*time.Second, func() {
			drone.Land()
		})
	}

	robot := gobot.NewRobot("hs200",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	robot.Start()
}
