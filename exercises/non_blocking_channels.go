package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan string)

	// receive
	select {
	case msg := <-messages:
		fmt.Println("Received from messages ::", msg)
	default:
		fmt.Println("Received nothing")
	}

	// send
	msg := "Hello"
	select {
	case messages <- msg:
		fmt.Println("Sent message")
	default:
		fmt.Println("Nothing sent")
	}

	//send and receive combined
	select {
	case msg := <-messages:
		fmt.Println("Received message ::", msg)
	case sig := <-signals:
		fmt.Println("Received a signal", sig)
	default:
		fmt.Println("No activity")
	}
}
