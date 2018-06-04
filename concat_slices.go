package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 6}
	b := []int{5, 9, 10, 12, 15}
	x := []int{}

	fmt.Println(a)
	fmt.Println(b)

	b = append(b, a...)
	x = append(x, b...)

	fmt.Println(x)

}
