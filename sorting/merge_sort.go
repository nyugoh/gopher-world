package main

import (
	"fmt"
)

var (
	arr     = make([]int, 10)
	arrTemp = make([]int, 10)
)

func init() {
	arr = []int{10, 7, 0, 2, 9, 4, 3, 20}
}

func printArray() {
	for _, val := range arr {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func merge(low, mid, high int) {
	tempIndex := low
	size := high - low - 1
	highStart := mid + 1
	for low <= mid && highStart <= high {
		if arr[low] <= arr[highStart] {
			arrTemp[tempIndex] = arr[low]
			tempIndex++
			low++
		}
		if arr[low] > arr[highStart] {
			arrTemp[tempIndex] = arr[highStart]
			tempIndex++
			highStart++
		}
	}
	for low <= mid {
		arrTemp[tempIndex] = arr[low]
		tempIndex++
		low++
	}
	for highStart <= high {
		arrTemp[tempIndex] = arr[highStart]
		tempIndex++
		highStart++
	}
	for j := 0; j < size; j++ {
		arr[high] = arrTemp[high]
		high--
	}
}

func mergeSort(low, high int) {
	if low < high {
		mid := low + (high-low)/2
		mergeSort(low, mid)
		mergeSort(mid+1, high)
		merge(low, mid, high)
	}
}

func main() {
	printArray()
	mergeSort(0, len(arr))
	printArray()
}
