package main

import "fmt"

func main() {
	hello := "ZABCDE"
	for i, char := range hello {
		var c byte
		c = 'a'
		c = hello[i] + 255
		fmt.Printf("Index %d = %d - %c C:%c\n", i, char, hello[i], c)
	}
}
