package main

import "fmt"

func main() {
	fmt.Println(true && true)   // TRUE
	fmt.Println(false && false) // false
	fmt.Println(true && false)  // FALSE
	fmt.Println(true || true)   // TRUE
	fmt.Println(false || false) // false
	fmt.Println(true || false)  // TRUE
	fmt.Println(!true)          // FALSE
	fmt.Println(!false)         // TRUE
}
