package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	mid := len(items)/2
	leftSide := mergeSort(items[:mid])
	rightSide := mergeSort(items[mid:])
	return merge(leftSide, rightSide)
}

func merge(a []int, b []int) []int {
	c := make([]int, len(a)+len(b))
	i , j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(a, b[j])
			j++
		}
	}

	// Append remaining elements in a & b
	for i < len(a) {
		c = append(c, a[i])
		i++
	}

	for j < len(b) {
		c = append(c, b[j])
		j++
	}
	return c
}

// don't modify below this line

func main() {
	start := time.Now()
	o1 := mergeSort(getValues(10))
	o2 := mergeSort(getValues(100))
	mergeSort(getValues(1000))
	mergeSort(getValues(10000))
	mergeSort(getValues(100000))
	if time.Since(start) < time.Second*3 {
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
