package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func randon(min int, max int) int {
	return rand.Intn(max-min) + min
}

// Should generate random no and send them thru a channel
func first(min int, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- randon(min, max)
	}

}

// Reads data from channel A and checks if it has been generated before, if not, sends it to last function
func second(in <-chan int, out chan<- int) {
	for x := range in {
		fmt.Printf("%d ", x)
		_, ok := DATA[x]
		if ok {
			//fmt.Println("Found a duplicate...", x)
			CLOSEA = true
		} else {
			if CLOSEA {
				return
			}
			//fmt.Println("Adding", x)
			DATA[x] = true
			out <- x
		}
	}
	close(out)
}

// Reads from channel B and sums up all the values
func third(in <-chan int) {
	var sum int
	sum = 0
	for val := range in {
		sum += val
	}
	fmt.Println("Sum:", sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Max and min args required")
		return
	}
	min, _ := strconv.Atoi(os.Args[1])
	max, _ := strconv.Atoi(os.Args[2])

	if min > max {
		fmt.Println("Min should be less than max")
		return
	}
	chanA := make(chan int)
	chanB := make(chan int)

	rand.Seed(time.Now().UnixNano())

	go first(min, max, chanA)
	go second(chanA, chanB)
	third(chanB)
	fmt.Println(DATA)
}
