package main

import (
	"fmt"
)

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func printArray(arr []int) {
	for _, val := range arr {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func main() {
	arr := []int{10, 7, 0, 2, 9, 4, 3, 20}
	printArray(arr)
	quickSort(arr, 0, len(arr)-1)
	printArray(arr)
}
