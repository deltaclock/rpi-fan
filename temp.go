package main

import (
	"os"
	"strconv"
)

// TempFile is the temperature as its being reported by the hardware
const TempFile string = "/sys/class/thermal/thermal_zone0/temp"

func getTemp() (uint64, error) {

	file, err := os.Open(TempFile)

	if err != nil {
		return 0, err
	}

	defer file.Close()

	buffer := make([]byte, 6)

	count, err := file.Read(buffer)

	if err != nil {
		panic(err)
	}

	temp := string(buffer[:count])
	tempC, err := strconv.ParseUint(temp, 10, 64)

	if err != nil {
		return 0, err
	}

	tempC /= 1000

	return tempC, nil

}
