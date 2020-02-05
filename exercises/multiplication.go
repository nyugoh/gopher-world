package main

import "fmt"

func main() {
	rows, columns := 10, 10
	for row := 1; row <= rows; row++ {
		fmt.Printf("%d | ", row)
		for col := 1; col < columns; col++ {
			fmt.Printf(" %d", row*col)
		}
		fmt.Println("")
	}
}
