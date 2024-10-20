//go:build example
// +build example

//
// Do not build by default.

package main

import (
	"log"
	"time"

	"github.com/szampardi/gobot"
	"github.com/szampardi/gobot/drivers/i2c"
	"github.com/szampardi/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	// we use the default address 0x60 for DC/Stepper Motor HAT on top of the Pi
	adaFruit := i2c.NewAdafruit2348Driver(r)

	work := func() {
		gobot.Every(5*time.Second, func() {
			motor := 0 // 0-based
			adafruitStepperMotorRunner(adaFruit, motor)
		})
	}

	robot := gobot.NewRobot("adaFruitBot",
		[]gobot.Connection{r},
		[]gobot.Device{adaFruit},
		work,
	)

	robot.Start()
}

func adafruitStepperMotorRunner(a *i2c.Adafruit2348Driver, motor int) error {
	log.Printf("Stepper Motor Run Loop...\n")
	// set the speed state:
	speed := 30 // rpm
	style := i2c.Adafruit2348Double
	steps := 20

	a.SetStepperMotorSpeed(motor, speed)

	if err := a.Step(motor, steps, i2c.Adafruit2348Forward, style); err != nil {
		log.Printf(err.Error())
		return err
	}
	if err := a.Step(motor, steps, i2c.Adafruit2348Backward, style); err != nil {
		log.Printf(err.Error())
		return err
	}
	return nil
}
