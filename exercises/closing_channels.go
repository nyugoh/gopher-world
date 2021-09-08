package main

import "fmt"

func main() {
	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			job, more := <-jobs
			if more {
				fmt.Println("Received job ::", job)
			} else {
				fmt.Println("No more jobs...")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("Sent a task ::", i)
		jobs <- i
	}
	close(jobs)
	fmt.Println("Sent all task your way.")
	<-done
}
