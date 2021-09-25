package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var options struct {
	// MaxTemp is the limit at which the fan will start running
	MaxTemp int
	// TempFile is the temperature as its being reported by the hardware
	TempFile string
	// Interval defines how often to check the temp
	Interval time.Duration
}

func init() {
	flag.IntVar(&options.MaxTemp, "temp", 50, "start fan after this temperature has been reached")
	flag.StringVar(&options.TempFile, "file", "/sys/class/thermal/thermal_zone0/temp", "path to the temp report")
	flag.DurationVar(&options.Interval, "time", 10*time.Second, "how often to check the temp")
}

func getTemp(tempFile string) int {

	file, err := os.Open(tempFile)

	if err != nil {
		panic(fmt.Sprintf("Couldn't open file %s for read!", tempFile))
	}

	defer file.Close()

	buffer := make([]byte, 7)

	count, err := file.Read(buffer)

	if err != nil {
		panic(err)
	}

	temp := string(buffer[:count-1])
	tempC, err := strconv.Atoi(temp)

	if err != nil {
		panic(err)
	}

	tempC /= 1000

	return tempC

}

func main() {
	flag.Parse()
	err := rpio.Open()

	if err != nil {
		panic(err)
	}

	defer rpio.Close()
	defer fmt.Println("defer run!")

	// https://gpiozero.readthedocs.io/en/stable/_images/pin_layout.svg
	pin := rpio.Pin(4) // GPIO4
	pin.Mode(rpio.Output)
	pin.Write(rpio.Low)

	spinning := false

	c := time.Tick(options.Interval)

	for range c {
		temp := getTemp(options.TempFile)

		if temp >= options.MaxTemp && !spinning {
			pin.Write(rpio.High)
			spinning = true
		} else if temp < options.MaxTemp && spinning {
			pin.Write(rpio.Low)
			spinning = false
		}

		fmt.Printf("Temp is at %d C\n", temp)
	}

}
