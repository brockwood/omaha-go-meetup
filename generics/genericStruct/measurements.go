package main

type Measurement interface {
	uint8 | float64
}

type Measurements[T Measurement] struct {
	Readings []T
}

func (h *Measurements[T]) AddReadings(reading ...T) {
	h.Readings = append(h.Readings, reading...)
}
func (h Measurements[T]) MaxReading() (max T) {
	for i, e := range h.Readings {
		if i == 0 || e > max {
			max = e
		}
	}
	return
}
func (h Measurements[T]) MinReading() (min T) {
	for i, e := range h.Readings {
		if i == 0 || e < min {
			min = e
		}
	}
	return
}
