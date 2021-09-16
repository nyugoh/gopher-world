package main

import (
	"fmt"
	"github.com/nyugoh/gopher-world/data-structures/stack/stack"
)

func main() {
	s := &stack.Stack{}

	fmt.Println("Is Empty::", s.IsEmpty())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(5)
	s.Print()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Print()
	fmt.Println("Top::", s.Top())

}
