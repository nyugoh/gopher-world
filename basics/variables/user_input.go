package main

import "fmt"

func main() {
	var age float64
	var name string
	fmt.Print("Enter your age:: ")
	fmt.Scanf("%f", &age)
	fmt.Print("Enter your name :")
	fmt.Scanf("%s", &name)
	fmt.Println("Your age is ", age)
	fmt.Println("Your name is ", name)
}
