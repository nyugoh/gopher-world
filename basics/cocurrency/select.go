package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func gen(min int, max int, createNo chan int, end chan bool)  {
	for {
		select {
		case createNo <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		case <-time.After(time.Second * 2):
			fmt.Println("\nAfter 2 seconds")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) != 2 {
		fmt.Println("Please provide at least 2 args")
		return
	}
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)

	createChan := make(chan int)
	end := make(chan bool)

	go gen(0, 2*n, createChan, end)

	for i:=0; i<n;i++ {
		fmt.Printf("%d ", <-createChan)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("Exiting...")
	end<-true
}
