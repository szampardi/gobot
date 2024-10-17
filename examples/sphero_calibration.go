//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/api"
	"github.com/szampardi/gobot/platforms/keyboard"
	"github.com/szampardi/gobot/platforms/sphero"
)

func main() {
	master := gobot.NewMaster()
	a := api.NewAPI(master)
	a.Start()

	ballConn := sphero.NewAdaptor("/dev/rfcomm0")
	ball := sphero.NewSpheroDriver(ballConn)

	keys := keyboard.NewDriver()

	calibrating := false

	work := func() {
		keys.On(keyboard.Key, func(data interface{}) {
			key := data.(keyboard.KeyEvent)

			switch key.Key {
			case keyboard.ArrowUp:
				if calibrating {
					break
				}
				ball.Roll(100, 0)
			case keyboard.ArrowDown:
				if calibrating {
					break
				}
				ball.Roll(100, 100)
			case keyboard.ArrowLeft:
				ball.Roll(100, 270)
			case keyboard.ArrowRight:
				ball.Roll(100, 90)
			case keyboard.Spacebar:
				if calibrating {
					ball.FinishCalibration()
				} else {
					ball.StartCalibration()
				}
				calibrating = !calibrating
			}
		})
	}

	robot := gobot.NewRobot("sphero-calibration",
		[]gobot.Connection{ballConn},
		[]gobot.Device{ball, keys},
		work,
	)

	master.AddRobot(robot)

	master.Start()
}
