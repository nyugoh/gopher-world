package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter two numbers to swap::")
	fmt.Printf("-->")
	fmt.Scanf("%d", &num1)
	fmt.Printf("-->")
	fmt.Scanf("%d", &num2)

	num1 = num1 + num2 // 10 + 5
	//fmt.Println(num1)  // 15
	//fmt.Println(num2)  // 5

	num2 = num1 - num2 // 15 - 5
	//fmt.Println(num1)  // 15
	//fmt.Println(num2)  // 10

	num1 = num1 - num2 // 15 - 10
	//fmt.Println(num1)  // 5
	//fmt.Println(num2)  // 10
}
