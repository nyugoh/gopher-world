package main

import "fmt"

func main() {
	length, width := 0.00, 0.00
	fmt.Println("Reactangle")
	fmt.Printf("\tEnter the length ::")
	fmt.Scanf("%f", &length)
	fmt.Printf("\tEnter the width ::")
	fmt.Scanf("%f", &width)
	fmt.Printf("\tArea = %f", length*width)

	fmt.Println("\nSquare")
	fmt.Printf("\tEnter the length ::")
	fmt.Scanf("%f", &length)
	fmt.Printf("\tArea = %f", length*length)
}
