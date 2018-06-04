package main

import (
	"fmt"
)

func main() {
	size, total := 0, 0

	fmt.Printf("Enter the size of array ::")
	fmt.Scanf("%d", &size)
	numbers := make([]int, size)

	for idx := 0; idx < size; idx++ {
		temp := 0
		fmt.Print("Enter number :: ")
		fmt.Scanf("%d", &temp)
		numbers[idx] = temp
	}

	for _, n := range numbers {
		total += n
	}

	fmt.Printf("Total :: %d\n", total)
	fmt.Printf("Average :: %d", total/size)
}
