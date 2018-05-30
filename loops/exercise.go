package main

import "fmt"

func main() {
    for i := 1; i <=100; i++ {
      if i%3 == 0 {
        fmt.Println(i," Fizz")
      }
      if i%5 == 0 {
        fmt.Println(i, " Buzz")
      }
      if i%3 == 0 && i%5 == 0 {
        fmt.Println(i, " FizzBuzz")
      }
    }
}
