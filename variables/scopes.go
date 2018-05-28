package main

import "fmt"

// Global scope Variables, accesible outside
var x string = "Hello world" // global scope

func main() {
  y := 40 //localscope
  fmt.Println(x)
  fmt.Println(y)
}
func foo()  {
  fmt.Println(x)
  fmt.Println(y) // Results to undefined error
}
