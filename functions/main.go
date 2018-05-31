package main

import "fmt"

func main() {
  a, b := 5, 7
  message := "Hello world"

  logger(message)

  fmt.Println(add(a, b))

  fmt.Println(multiply(&a, &b))
}

func logger(message string)  {
  fmt.Println(message)
}

func add(a, b int) int  {
  return a+b
}

func multiply(a, b *int) int  {
  return *a**b
}
