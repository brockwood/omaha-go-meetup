package main

type Temperature float64

func (t Temperature) Celsius() Temperature {
	return (t - 32.0) * (5.0 / 9.0)
}
