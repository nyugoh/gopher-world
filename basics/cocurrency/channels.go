package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, n := range a {
		sum += n
	}

	c <- sum
}

func main() {
	scores := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	channel := make(chan int)

	go sum(scores[:len(scores)/2], channel)
	go sum(scores[len(scores)/2:], channel)
	x, y := <-channel, <-channel

	fmt.Printf("%d+%d=%d\n", x, y, x+y)
}
