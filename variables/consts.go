package main

import "fmt"

func main() {
  const x int = 22
  fmt.Println(x)
  x = 20 // results to error, cannot assing to x
}
