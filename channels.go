package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 5)
	ch1 := make(chan bool)
	go func() { channel <- 10 }()

	/*for val, ok := <- channel; ok != false; {
		fmt.Println(val)
	}*/

	go func(channel chan bool) {
		fmt.Println("Working ...")
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("Done !")
		channel <- true
	}(ch1)

	fmt.Print(<-ch1) // block to wait for goroutine to exit
}
