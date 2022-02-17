package main

import "fmt"

func main() {
	var temp Temperature
	var rh Humidity
	temp.AddReading(-32.0, 12.5, 17.6, -15.0)
	rh.AddReading(9, 52, 45, 12)
	fmt.Printf("max temp: %.1f, min temp %.1f\n", temp.MaxReading(), temp.MinReading())
	fmt.Printf("max humidity: %d, min humidity %d\n", rh.MaxReading(), rh.MinReading())
}
