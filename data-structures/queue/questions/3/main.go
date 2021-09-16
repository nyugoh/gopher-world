package main

import (
	"fmt"
	"github.com/nyugoh/gopher-world/data-structures/queue/queue"
	"github.com/nyugoh/gopher-world/data-structures/stack/stack"
	"math"
)

/*
Given a stack of integers, how do you check whether each successive pair of numbers in the stack is consecutive or not.
The pairs can be increasing or decreasing, and if the stack has an odd number of elements, the element at the top is left out of a pair.

For example, if the stack of elements are
	[4, 5, -2, -3, 11, 10, 5, 6, 20],
then the output should be true because each of the pairs
	(4, 5), (-2, -3), (11, 10), and (5, 6) consists of consecutive numbers.
*/

func main() {
	s := &stack.Stack{}
	s.Push(4)
	s.Push(5)
	s.Push(-2)
	s.Push(-3)
	s.Push(11)
	s.Push(10)
	s.Push(5)
	s.Push(6)
	s.Push(20)

	s.Print()
	q := queue.CreateQueue()
	// 1 Reverse the stack - so that the top element is at the bottom
	// 1.1 Push to queue
	for !s.IsEmpty() {
		q.Enqueue(s.Top())
		s.Pop()
	}
	// 1.2 Push back to stack
	for !q.IsEmpty() {
		s.Push(q.Dequeue())
	}

	// 2. Remove the elements from the stack and compare them in order
	isOrdered := true
	s.Print()
	for !s.IsEmpty() {
		m := s.Top()
		s.Pop()
		q.Enqueue(m)
		if !s.IsEmpty() {
			n := s.Top()
			s.Pop()
			q.Enqueue(n)
			if math.Abs(float64(m-n)) != 1.0 {
				isOrdered = false
			}
		}
	}
	// Move the elements back to stack
	for !q.IsEmpty() {
		s.Push(q.Dequeue())
	}
	s.Print()
	if isOrdered {
		fmt.Println("Is ordered")
	} else {
		fmt.Println("Not ordered")
	}
}
