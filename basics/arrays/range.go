package main

import "fmt"

func main() {
	// Define variables
	scores := []float64{60, 70, 50, 90, 60}
	var total float64 = 0

	// Loop thru them
	/* for i, value := range scores { // results in error coz var i is not used
	   total += value
	 } */
	for _, value := range scores { // use _ to represent the unsed var, tells compiler we don't need this
		total += value
	}

	fmt.Println("Total ::", total)
	fmt.Println("Average ::", total/float64(len(scores)))
}
