package main

import "fmt"

func main() {
	matrixA := [][]int{{1, 2}, {2, 3}}
	matrixB := [][]int{{10, 20}, {30, 40}}
	sum := [][]int{{0, 0}, {0, 0}}

	fmt.Println(matrixA)
	fmt.Println(matrixB)

	for indexi := 0; indexi < 2; indexi++ {
		for indexj := 0; indexj < 2; indexj++ {
			sum[indexi][indexj] = matrixA[indexi][indexj] + matrixB[indexi][indexj]
		}
	}
	fmt.Println(sum)
}
