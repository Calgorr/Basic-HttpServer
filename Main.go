package main

import (
	"fmt"
)

func main() {
	xs := []float64{1, 2, 3, 4, 5}
	fmt.Println(Average(xs))
}

func Average(xs []float64) float64 {
	total := 0.0
	for _, x := range xs {
		total += x
	}
	xs[0] = 5
	return total / float64(len(xs))
}
