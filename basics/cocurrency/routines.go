package main

import (
  "fmt"
  "time"
)

func main() {
  go function()
  go func() {
    for i:=10; i <= 20; i++ {
      fmt.Print(i, " ")
    }
  }()

  time.Sleep(time.Second * 2)
  fmt.Println("Done...!")

}

func function()  {
  for i := 0; i < 10; i++ {
    fmt.Print(i)
  }
}
