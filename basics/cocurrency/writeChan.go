package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(time.Second)
	fmt.Printf("Reading :%d\n", <-c)
	time.Sleep(time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open.")
	} else {
		fmt.Println("Channel is closed.")
	}
	fmt.Println("Exiting...")

}

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}
