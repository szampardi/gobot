//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"fmt"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/platforms/joystick"
)

func main() {
	joystickAdaptor := joystick.NewAdaptor("0")
	stick := joystick.NewDriver(joystickAdaptor, joystick.Dualshock3)

	work := func() {
		// buttons
		stick.On(joystick.SquarePress, func(data interface{}) {
			fmt.Println("square_press")
		})
		stick.On(joystick.SquareRelease, func(data interface{}) {
			fmt.Println("square_release")
		})
		stick.On(joystick.TrianglePress, func(data interface{}) {
			fmt.Println("triangle_press")
		})
		stick.On(joystick.TriangleRelease, func(data interface{}) {
			fmt.Println("triangle_release")
		})
		stick.On(joystick.CirclePress, func(data interface{}) {
			fmt.Println("circle_press")
		})
		stick.On(joystick.CircleRelease, func(data interface{}) {
			fmt.Println("circle_release")
		})
		stick.On(joystick.XPress, func(data interface{}) {
			fmt.Println("x_press")
		})
		stick.On(joystick.XRelease, func(data interface{}) {
			fmt.Println("x_release")
		})
		stick.On(joystick.StartPress, func(data interface{}) {
			fmt.Println("start_press")
		})
		stick.On(joystick.StartRelease, func(data interface{}) {
			fmt.Println("start_release")
		})
		stick.On(joystick.SelectPress, func(data interface{}) {
			fmt.Println("select_press")
		})
		stick.On(joystick.SelectRelease, func(data interface{}) {
			fmt.Println("select_release")
		})
		stick.On(joystick.HomePress, func(data interface{}) {
			fmt.Println("home_press")
		})
		stick.On(joystick.HomeRelease, func(data interface{}) {
			fmt.Println("home_release")
		})
		stick.On(joystick.RightPress, func(data interface{}) {
			fmt.Println("right_press")
		})
		stick.On(joystick.RightRelease, func(data interface{}) {
			fmt.Println("right_release")
		})
		stick.On(joystick.LeftPress, func(data interface{}) {
			fmt.Println("left_press")
		})
		stick.On(joystick.LeftRelease, func(data interface{}) {
			fmt.Println("left_release")
		})
		stick.On(joystick.UpPress, func(data interface{}) {
			fmt.Println("up_press")
		})
		stick.On(joystick.UpRelease, func(data interface{}) {
			fmt.Println("up_release")
		})
		stick.On(joystick.DownPress, func(data interface{}) {
			fmt.Println("down_press")
		})
		stick.On(joystick.DownRelease, func(data interface{}) {
			fmt.Println("down_release")
		})

		// joysticks
		stick.On(joystick.LeftX, func(data interface{}) {
			fmt.Println("left_x", data)
		})
		stick.On(joystick.LeftY, func(data interface{}) {
			fmt.Println("left_y", data)
		})
		stick.On(joystick.RightX, func(data interface{}) {
			fmt.Println("right_x", data)
		})
		stick.On(joystick.RightY, func(data interface{}) {
			fmt.Println("right_y", data)
		})

		// triggers
		stick.On(joystick.R1Press, func(data interface{}) {
			fmt.Println("R1Press", data)
		})
		stick.On(joystick.R1Release, func(data interface{}) {
			fmt.Println("R1Release", data)
		})
		stick.On(joystick.R2Press, func(data interface{}) {
			fmt.Println("R2Press", data)
		})
		stick.On(joystick.R2Release, func(data interface{}) {
			fmt.Println("R2Release", data)
		})
		stick.On(joystick.L1Press, func(data interface{}) {
			fmt.Println("L1Press", data)
		})
		stick.On(joystick.L1Release, func(data interface{}) {
			fmt.Println("L1Release", data)
		})
		stick.On(joystick.L2Press, func(data interface{}) {
			fmt.Println("L2Press", data)
		})
		stick.On(joystick.L2Release, func(data interface{}) {
			fmt.Println("L2Release", data)
		})
	}

	robot := gobot.NewRobot("joystickBot",
		[]gobot.Connection{joystickAdaptor},
		[]gobot.Device{stick},
		work,
	)

	robot.Start()
}
