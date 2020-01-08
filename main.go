package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()

	if err != nil {
		panic(err)
	}

	defer rpio.Close()
	// https://gpiozero.readthedocs.io/en/stable/_images/pin_layout.svg
	pin := rpio.Pin(4)
	pin.Mode(rpio.Output)
	pin.Write(rpio.Low)

	for {
		time.Sleep(time.Second * 2)
		pin.Toggle()
		fmt.Println(getTemp())
	}

}
