package main

type Humidity struct {
	Readings []uint8
}

func (h *Humidity) AddReading(reading ...uint8) {
	h.Readings = append(h.Readings, reading...)
}
func (h Humidity) MaxReading() (max uint8) {
	for i, e := range h.Readings {
		if i == 0 || e > max {
			max = e
		}
	}
	return
}
func (h Humidity) MinReading() (min uint8) {
	for i, e := range h.Readings {
		if i == 0 || e < min {
			min = e
		}
	}
	return
}
