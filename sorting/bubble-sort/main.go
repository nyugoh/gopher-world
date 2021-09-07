package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubbleSort(input []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i-1] > input[i] {
				input[i-1], input[i] = input[i], input[i-1]
				swapped = true
			}
		}
	}
	return input
}

// don't modify below this line

func main() {
	start := time.Now()
	o1 := bubbleSort(getValues(10))
	o2 := bubbleSort(getValues(100))
	bubbleSort(getValues(1000))
	bubbleSort(getValues(10000))
	if time.Since(start) < time.Millisecond*10000 {
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
