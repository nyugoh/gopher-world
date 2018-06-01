package main

import "fmt"

type Stock struct {
  sharePrice float64
  count float64
}

func (s *Stock) getValue() float64 {
  return s.sharePrice * s.count
}

type Car struct {
  Name string
  Model string
  Price float64
}

func (c *Car) getValue() float64 {
  return c.Price
}

type Valuation interface {
  getValue() float64
}

func showValue(obj Valuation)  {
  fmt.Println(obj.getValue())
}

func main() {
  var item Valuation = &Stock{4.5, 200}
  showValue(item)

  item = &Car{"Lexus", "2018", 9000}

  showValue(item)
}
