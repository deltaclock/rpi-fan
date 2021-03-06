package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	// MaxTemp is the limit at which the fan will start running
	MaxTemp = 50
	// INTERVAL defines how often to check the temp
	INTERVAL = time.Second * 5
)

func main() {
	err := rpio.Open()

	if err != nil {
		panic(err)
	}

	defer rpio.Close()
	// https://gpiozero.readthedocs.io/en/stable/_images/pin_layout.svg
	pin := rpio.Pin(4) // GPIO4
	pin.Mode(rpio.Output)
	pin.Write(rpio.Low)
	spinning := false

	for {
		temp, err := getTemp()

		if err != nil {
			panic(err)
		}

		if temp >= MaxTemp && !spinning {
			pin.Write(rpio.High)
			spinning = true
		} else if temp < MaxTemp && spinning {
			pin.Write(rpio.Low)
			spinning = false
		}

		fmt.Printf("Temp is at %d C\n", temp)
		time.Sleep(INTERVAL)
	}

}
