package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "No. of goroutines to create")
	flag.Parse()

	count := *n

	fmt.Printf("Going to create %d goroutines...\n", count)

	var waitGroup sync.WaitGroup
	fmt.Printf("\n%#v\n", waitGroup)

	for i := 1; i <= count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("\n%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
