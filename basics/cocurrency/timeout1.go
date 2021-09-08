package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 is okay"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 2):
		fmt.Println("c1 timeout after 2 seconds")
	}

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "c2 is okay"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("c2 timeout after 4 seconds")
	}

}
