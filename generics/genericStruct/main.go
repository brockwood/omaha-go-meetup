package main

import "fmt"

func main() {
	var temperature Measurements[float64]
	var rh Measurements[uint8]
	temperature.AddReadings(-32.0, 12.5, 17.6, -15.0)
	rh.AddReadings(9, 52, 45, 12)
	fmt.Printf("max temp: %.1f, min temp %.1f\n", temperature.MaxReading(), temperature.MinReading())
	fmt.Printf("max humidity: %d, min humidity %d\n", rh.MaxReading(), rh.MinReading())
}
