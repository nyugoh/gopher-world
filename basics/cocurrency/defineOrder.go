package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func A(x, y chan struct{})  {
	<-x
	time.Sleep(time.Second*2)
	fmt.Println("A()!")
	waitGroup.Done()
	close(y)
}

func B(x, y chan struct{})  {
	<-x
	time.Sleep(time.Second * 3)
	fmt.Println("B()!")
	waitGroup.Done()
	close(y)
}

func C(y chan struct{})  {
	<-y
	time.Sleep(time.Second*2)
	fmt.Println("C()!")
	waitGroup.Done()
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	waitGroup.Add(5)
	go C(z)
	go A(x,y)
	go C(z)
	go B(y,z)
	go C(z)

	close(x)
	waitGroup.Wait()
}
