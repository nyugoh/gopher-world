package main

import "fmt"

func main() {
  go sayHello("Hello")
  sayHello("World")
}

func sayHello(s string)  {
  for i := 0; i < 5; i++ {
    fmt.Println(i, s)
  }
}
