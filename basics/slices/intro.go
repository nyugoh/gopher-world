package main

import "fmt"

func main() {
	// Declaration
	var scores []float64
	// or
	i := make([]float64, 5, 10) // Length of 5 but capacity of 10
	p := [5]float64{1, 2, 3, 4, 5}
	// q := p[0:] // index 0  to the end
	// q := p[:] // all elements
	// q := p[0:2] // index 0 - 2, 2 not included
	q := p[:2] // index 0 - 2
	// [: a] index a is not included

	fmt.Println(scores)
	fmt.Println(i)
	fmt.Println(q)
}
