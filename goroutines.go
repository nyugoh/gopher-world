package main

import "fmt"

func f(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, ": ", i)
	}
}

func main() {
	go f("goroutine")

	f("First")

	go func(msg string) {
		fmt.Println(msg)
	}("anno")

	fmt.Scanln()
	fmt.Println("Done")
}
