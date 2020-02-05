// Closures are functions that maintain their value
/*
Below we initialize a sequence and increment it by calling the rturned function
*/
package main

import "fmt"

func initSeq() func() int {
	i := 0 // start of the sequence
	return func() int {
		i++      //increase the value
		return i //return it
	}
}

func main() {
	nextSeq := initSeq()

	fmt.Println(nextSeq())
	fmt.Println(nextSeq())
	fmt.Println(nextSeq())

	// Start sequence again
	newSeq := initSeq()
	fmt.Println(newSeq())
	fmt.Println(newSeq())
	fmt.Println(newSeq())

	// has it's closure
	fmt.Println(nextSeq())
	fmt.Println(newSeq())
}
