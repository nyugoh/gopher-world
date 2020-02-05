package main

import "fmt"

func main() {
	rows := 0

	// Ask user input
	fmt.Print("Enter the number of rows ::")
	fmt.Scanf("%d", &rows)

	for row := 0; row < rows; row++ {
		for spaces := rows; spaces > row; spaces-- {
			fmt.Printf(" ")
		}
		for stars := 0; stars <= row*2; stars++ {
			fmt.Printf("*")
		}

		fmt.Printf("\n")
	}
}
