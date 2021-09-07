package main

import (
	"fmt"
)

func print(arr []int) {
	for _, val := range arr {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func merge(arr, tempArr []int, low, mid, high int) {
	fmt.Println("Merge...")
	tempIndex := low
	size := (high - low) + 1
	highStart := mid + 1
	for low <= mid && highStart <= high {
		fmt.Println("Low:", low, "Mid:", mid, "High", high)
		if arr[low] <= arr[highStart] {
			tempArr[tempIndex] = arr[low]
			tempIndex++
			low++
		}
		if arr[low] > arr[highStart] {
			tempArr[tempIndex] = arr[highStart]
			tempIndex++
			highStart++
		}
	}
	for low <= mid {
		tempArr[tempIndex] = arr[low]
		tempIndex++
		low++
	}
	for highStart <= high {
		tempArr[tempIndex] = arr[highStart]
		tempIndex++
		highStart++
	}
	for j := 0; j < size; j++ {
		arr[high] = tempArr[high]
		high--
	}
}

func mergeSort(arr, tempArr []int, low, high int) {
	if low < high {
		mid := low + (high-low)/2
		fmt.Println("Mid:", mid)
		mergeSort(arr, tempArr, low, mid)
		mergeSort(arr, tempArr, mid+1, high)
		merge(arr, tempArr, low, mid, high)
	}
	return
}

func main1() {
	arr := []int{10, 20, 3, 4, 1, 23, 8, 0}
	tempArr := []int{0,0, 0, 0,0 ,0 ,0,0}
	print(arr)
	mergeSort(arr, tempArr, 0, len(arr)-1)
	print(arr)
}
