package intro

import (
	"fmt"
	"math/rand"
)

func arrays() {
	var numbers [3]int
	var cities [3]string
	var matrix [3][3] int

	fmt.Println("Inserting data:")
	for i:=0;i<3;i++ {
		numbers[i] = i
		cities[i] = fmt.Sprintf("City %d", i)
	}
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			matrix[i][j] = rand.Intn(10)
		}
	}

	fmt.Println("Displaying content")
	for i, num := range cities {
		fmt.Printf("%d ==> %s\n", numbers[i], num)
	}
	for indexI, i := range matrix {
		for indexJ, j := range i {
			fmt.Printf("[%d][%d] = [%d]\t", indexI, indexJ, j)
		}
		fmt.Println()
	}
}
