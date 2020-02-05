package main

import (
  "fmt"
  "math"
)

type Shape struct {
  Name string
}

type Shaper interface {
  Area() float64
}

type Square struct {
  Length float64
  Shape *Shape
}

type Rectangle struct {
  Length float64
  Width float64
}

type Circle struct {
  Radius float64
}

// Implement the interface for shaper
func (s *Square) Area() float64 {
  return s.Length*s.Length
}

func (s *Rectangle) Area() float64 {
  return s.Length*s.Width
}

func (s *Circle) Area() float64 {
  return math.Pi*s.Radius*s.Radius
}

func main() {
  s := &Square{5, &Shape{"Square pants"}}
  r := &Rectangle{5, 10}
  c := &Circle{5}
  shapes := []Shaper{s, r, c}

  for index, _ := range shapes {
    fmt.Println(shapes[index])
    fmt.Println("Area ::", shapes[index].Area())
  }
}
