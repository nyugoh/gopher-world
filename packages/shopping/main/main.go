package main

import (
  "fmt"
  "Tuts/packages/shopping"
  "Tuts/packages/shopping/checkout"
)
func main() {
  price, ok := shopping.PriceCheck(45)
  fmt.Println(price)
  fmt.Println(ok)

  fmt.Println(checkout.CalculateTotal([]*float64{&price, &price, &price}))
}
