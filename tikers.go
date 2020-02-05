package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for t := range ticker.C {
			fmt.Printf("\r Tick at %v", t)
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
	fmt.Println("\nTicker stopped")
}
