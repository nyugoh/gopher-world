package main

import "fmt"

func main() {
  ch := make(chan int, 2)
  ch <- 5
  ch <- 10

  fmt.Println("Reading from channel ::", <- ch)
  fmt.Println("Reading from channel ::", <- ch)
  //fmt.Println("Reading from channel ::", <- ch) // error
}
