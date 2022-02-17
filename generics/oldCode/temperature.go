package main

type Temperature struct {
	Readings []float32
}

func (t *Temperature) AddReading(reading ...float32) {
	t.Readings = append(t.Readings, reading...)
}
func (t Temperature) MaxReading() (max float32) {
	for i, e := range t.Readings {
		if i == 0 || e > max {
			max = e
		}
	}
	return
}
func (t Temperature) MinReading() (min float32) {
	for i, e := range t.Readings {
		if i == 0 || e < min {
			min = e
		}
	}
	return
}
