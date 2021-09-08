package utils

import "fmt"

// Returns the average of a float64 slice
func Average(xs []float64) (avg float64) {
	if len(xs) == 0 || xs == nil {
		return 0
	}
	for _, x := range xs {
		avg += x
	}
	avg = avg / float64(len(xs))
	return
}

// Tests local package usage
func test(arg string) {
	fmt.Println("Printing arg sent: ", arg)
}
