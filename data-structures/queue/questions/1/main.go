package main

import (
	"github.com/nyugoh/gopher-world/data-structures/queue/queue"
	"github.com/nyugoh/gopher-world/data-structures/stack/stack"
)

/*
Given an integer k and a queue of integers, how do you reverse the order of the first k elements of the queue,
leaving the other elements in the same relative order?

For example, if k=4 and queue has the elements [10, 20, 30, 40, 50, 60, 70, 80, 90];
the output should be [40, 30, 20, 10, 50, 60, 70, 80, 90].
*/

func main() {
	k := 4
	q := queue.CreateQueue()
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)
	q.Enqueue(60)
	q.Enqueue(70)
	q.Enqueue(80)
	q.Enqueue(90)

	q.Print()

	s := &stack.Stack{}
	i := 0
	for i < k {
		s.Push(q.Dequeue())
		i += 1
	}
	for !s.IsEmpty() {
		q.Enqueue(s.Top())
		s.Pop()
	}

	s.Print()
	q.Print()

	for ; i < q.Size(); i++ {
		q.Enqueue(q.Dequeue())
	}
	q.Print()
}
