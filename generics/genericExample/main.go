package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](readings ...T) (max T) {
	for i, e := range readings {
		if i == 0 || e > max {
			max = e
		}
	}
	return
}

func main() {
	fmt.Printf("Max temp: %.1f\n", (Max[float64](-32.0, 12, 17.6))) // explicit set of type
	fmt.Printf("Max temp: %.1f\n", (Max(-32.0, 12.5, 17.6)))        // type is inferred
}
