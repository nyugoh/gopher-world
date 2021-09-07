package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i; j > 0 && a[j-1] > a[j]; j-- {
			a[j-1], a[j] = a[j], a[j-1]
		}
	}
	return a
}

// don't modify below this line

func main() {
	start := time.Now()
	o1 := insertionSort(getValues(10))
	o2 := insertionSort(getValues(100))
	insertionSort(getValues(1000))
	insertionSort(getValues(10000))
	if time.Since(start) < time.Second {
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
