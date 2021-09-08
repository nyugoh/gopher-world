package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(a []int, low, high int) []int {
	if low < high {
		p := partition(a, low, high)
		quickSort(a, low, p-1)
		quickSort(a, p+1, high)
	}
	return a
}

func partition(a []int, low, high int) int {
	pivot := a[high]
	i := low
	for j := low; j < high; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[high] = a[high], a[i]
	return i
}

// don't modify below this line

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	start := time.Now()
	o1 := quickSortStart(getValues(10))
	o2 := quickSortStart(getValues(100))
	quickSortStart(getValues(1000))
	quickSortStart(getValues(10000))
	quickSortStart(getValues(100000))
	if time.Since(start) < time.Millisecond*500 {
		fmt.Println("Fast :)")
	} else {
		fmt.Println("Slow :(")
	}
	fmt.Println(o1)
	fmt.Println(o2)
}

func getValues(num int) []int {
	rand.Seed(0)
	nums := []int{}
	for i := 0; i < num; i++ {
		nums = append(nums, rand.Intn(num))
	}
	return nums
}
