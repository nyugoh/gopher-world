package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	n := flag.Int("n", 10, "no. of goroutines to create")
	flag.Parse()

	count := *n

	fmt.Printf("Creating %d goroutines...\n", count)

	for i:=1; i <= count ;i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
