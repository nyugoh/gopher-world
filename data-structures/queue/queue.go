package main

import (
	"fmt"
	"github.com/nyugoh/gopher-world/data-structures/queue/queue"
)

func main() {
	q := queue.CreateQueue()

	fmt.Println("IS empty::", q.IsEmpty())

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Print()
	q.Dequeue()
	q.Dequeue()
	q.Print()
	q.DeleteQ()
	fmt.Println("IS empty::", q.IsEmpty())

}
