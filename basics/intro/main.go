package intro

import (
	"fmt"
	"github.com/nyugoh/tuts/utils"
	"math/rand"
	"time"
)

func chans() {
	for i:=0;i<10;i++ {
		// go f(i)
	}

	c := make(chan string)
	go pinger(c)
	time.Sleep(time.Second * 2)
	fmt.Println(<-c)
	go ponger(c)
	time.Sleep(time.Second * 2)
	fmt.Println(<-c)
	go receiver(c)
	var char string
	fmt.Scanf("%s", &char)
}


func pinger(c chan <- string) {
	var i int
	for i=0; ; i++ {
		msg := "Ping"
		c <- msg
		//fmt.Println("Res Ping:", <- c)
	}
}

func ponger(c chan<- string) {
	for {
		c <- "Pong"
		//fmt.Println("Res Pong:", <- c)
	}
}


func receiver(c <-chan string) {
	for {
		msg := <- c
		fmt.Println("Req:", msg)
		if msg != "" {
			//c <- "Okay"
		}
		time.Sleep(time.Millisecond * 500)
	}
}


func f(n int) {
	for i:=0;i<10;i++ {
		fmt.Printf("%d==>%d \n", n, i)
		time.Sleep(time.Millisecond* time.Duration(rand.Intn(250)))
	}
}

func testUtils() {
	xs := []float64{20, 40, 60, 80}
	fmt.Printf("Total avg of %v on date %s is :%f", xs, utils.CurrentDate(), utils.Average(xs))
}
