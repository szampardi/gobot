//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/platforms/opencv"
	"gocv.io/x/gocv"
)

func main() {
	window := opencv.NewWindowDriver()
	camera := opencv.NewCameraDriver(0)

	work := func() {
		camera.On(opencv.Frame, func(data interface{}) {
			img := data.(gocv.Mat)
			window.ShowImage(img)
			window.WaitKey(1)
		})
	}

	robot := gobot.NewRobot("cameraBot",
		[]gobot.Device{window, camera},
		work,
	)

	robot.Start()
}
