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
	"github.com/szampardi/gobot/platforms/chip"
)

func main() {
	board := chip.NewAdaptor()
	mpu6050 := i2c.NewMPU6050Driver(board)

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			mpu6050.GetData()

			fmt.Println("Accelerometer", mpu6050.Accelerometer)
			fmt.Println("Gyroscope", mpu6050.Gyroscope)
			fmt.Println("Temperature", mpu6050.Temperature)
		})
	}

	robot := gobot.NewRobot("mpu6050Bot",
		[]gobot.Connection{board},
		[]gobot.Device{mpu6050},
		work,
	)

	robot.Start()
}
