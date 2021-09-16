package main

import (
	"github.com/nyugoh/gopher-world/data-structures/queue/queue"
	"github.com/nyugoh/gopher-world/data-structures/stack/stack"
)

/*
Given a queue of integers, rearrange the elements by interleaving the first half of the list with the second half of the list.
For example,
	suppose a queue stores the following sequence of values: [11, 12, 13, 14, 15, 16, 17, 18, 19, 20].
	Consider the two halves of this list:
		first half: [11, 12, 13, 14, 15]
		second half: [16, 17, 18, 19, 20].

	These are combined in an alternating fashion to form a sequence of interleave pairs:
	the first values from each half (11 and 16), then the second values from each half (12 and 17),
	then the third values from each half (13 and 18), and so on.
	In each pair, the value from the first half appears before the value from the second half.
	Thus, after the call, the queue stores the following values:
		[11, 16, 12, 17, 13, 18, 14, 19, 15, 20].
*/

func main() {
	q := queue.CreateQueue()
	q.Enqueue(11)
	q.Enqueue(12)
	q.Enqueue(13)
	q.Enqueue(14)
	q.Enqueue(15)
	q.Enqueue(16)
	q.Enqueue(17)
	q.Enqueue(18)
	q.Enqueue(19)
	q.Enqueue(20)

	mid := q.Size() / 2
	s := &stack.Stack{}
	// Move first half to stack q= [16, 17, 18, 19, 20]
	for i := 0; i < mid; i++ {
		s.Push(q.Dequeue())
	}
	q.Print()

	// Move the 1st half back to the queue q=[16, 17, 18, 19, 20, 15, 14, 13, 12, 11]
	for !s.IsEmpty() {
		q.Enqueue(s.Top())
		s.Pop()
	}
	q.Print()

	// Move the 1st half to the front  q=[15, 14, 13, 12, 11, 16, 17, 18, 19, 20,]
	for i := 0; i < mid; i++ {
		q.Enqueue(q.Dequeue())
	}
	q.Print()

	// Move the first half to a stack q = [16, 17, 18, 19, 20,]
	for i := 0; i < mid; i++ {
		s.Push(q.Dequeue())
	}
	q.Print()

	// Combine from stack and from of queue
	for !s.IsEmpty() {
		q.Enqueue(s.Top())
		s.Pop()
		q.Enqueue(q.Dequeue())
	}
	q.Print()
}
