package main

import (
	"fmt"
	"math"
)

type geometry interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r *Rectangle) Perimeter() float64 {
	return (r.Length + r.Width) * 2
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * (c.Radius * 2)
}

func measure(g geometry) {
	fmt.Println("Area :: ", g.Area())
	fmt.Println("Permeter :: ", g.Perimeter())
}

func main() {
	rec := &Rectangle{10, 20}
	circle := &Circle{20}

	measure(rec)
	measure(circle)
}
