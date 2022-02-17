package main

import "fmt"

func main() {
	var temperature Measurements[Temperature]
	temperature.AddReadings(-32.0, 12.5, 17.6, -15.0)
	fmt.Printf("max temp: %.1f°C, min temp %.1f°C\n",
		temperature.MaxReading().Celsius(), temperature.MinReading().Celsius())
}
