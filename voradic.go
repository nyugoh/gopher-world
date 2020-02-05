package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	sum := 0
	for _, n := range nums {
		sum += n
	}
	fmt.Print(sum)
}
func main() {
	nums := []int{1, 3, 5, 6, 6}
	sum(nums...)
}
