package main

import "fmt"

func isSorted(a []int, n int) bool {
	if n == 1{
		return true
	}
	if a[n-1] < a[n-2] {
		return false
	}
	return isSorted(a, n-1)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if isSorted(a, len(a)) {
		fmt.Println("Is sorted")
	} else {
		fmt.Println("Not sorted")
	}
}
