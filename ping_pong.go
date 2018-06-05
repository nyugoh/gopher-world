package main

import (
	"fmt"
	"strings"
)

func ping(pingChannel chan<- string, message string) {
	pingChannel <- message // Send the message thru the ping channel
}

func pong(pingChannel <-chan string, pongChannel chan<- string) {
	message := <-pingChannel                // Receive the message from ping channel
	pongChannel <- strings.ToUpper(message) // Send the message back on the channel
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	message := ""
	fmt.Print("Enter the message :: ")
	fmt.Scanf("%s", &message)

	ping(pings, message)
	pong(pings, pongs)

	fmt.Println(<-pongs) // Receive the message back
}
