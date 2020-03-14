package main

import (
	"fmt"
)

var (
	arr = make([]int, 10)
)

func init() {
	arr = []int{10, 7, 0, 2, 9, 4, 3, 20}
}

func partition(low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quickSort(low int, high int) {
	if low < high {
		p := partition(low, high)
		quickSort(low, p-1)
		quickSort(p+1, high)
	}
}

func printArray() {
	for _, val := range arr {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func main() {
	printArray()
	quickSort(0, len(arr)-1)
	printArray()
}
