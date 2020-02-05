package main

import (
	"fmt"
	"time"
)

func main() {
	// Run a goroutine that execs for 2 seconds
	chan1 := make(chan string)

	go func(c1 chan string) {
		fmt.Print("Working ...")
		time.Sleep(2 * time.Second)
		c1 <- "done!"
	}(chan1)

	select {
	case response1 := <-chan1:
		fmt.Println("Response :: ", response1)
	case <-time.After(1 * time.Second): // Timeout after 1 second
		fmt.Println("Timeout")
	}

	chan2 := make(chan string)

	go func(c2 chan string) {
		fmt.Print("Working ...")
		time.Sleep(2 * time.Second)
		c2 <- "done!"
	}(chan2)

	select {
	case response2 := <-chan2:
		fmt.Println("Response :: ", response2)
	case <-time.After(3 * time.Second): // Timeout after 3 seconds
		fmt.Println("Timeout")
	}

}
