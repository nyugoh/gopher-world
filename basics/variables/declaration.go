package main

import "fmt"

func main() {
	var y int = 30 // Type is explictly assinged
	x := 20        // Type is infered, determined on runtime
	fmt.Println(x)
	fmt.Println(y)
}
