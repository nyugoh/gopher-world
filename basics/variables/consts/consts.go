package main

import "fmt"

func main() {
	const x int = 22
	const y = 22
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	//x = 20 // results to error, cannot assing to x
}
