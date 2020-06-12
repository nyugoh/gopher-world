package main

import "fmt"

type Node struct {
	next *Node
	data int
}

type Queue struct {
	head *Node
}

// Remove the first element infront of the queue
func dequeue(queue *Queue) (int, bool) {
	// element is remove from the front of list
	// basically removing the head
	if queue.head == nil {
		fmt.Println("Queue is empty")
		return 0, false
	}
	data := queue.head.data
	queue.head = queue.head.next
	return data, false
}

// Add an element to the queue
func enqueue(queue Queue, data int)  {
	// element is added at the end of list
	newNode := &Node{
		data: data,
		next: nil,
	}
	if queue.head == nil {
		queue.head = newNode
		return
	}

	tempNode := queue.head
	for tempNode.next != nil {
		tempNode = tempNode.next
	}
	tempNode.next = newNode
}

// Return element at the front
func top(queue Queue) (int, bool) {
	if queue.head == nil {
		return 0, true
	}
	return queue.head.data, false
}

func printQueue(queue Queue)  {
	tempNode := queue.head
	for tempNode != nil {
		fmt.Print(tempNode.data, ", ")
		tempNode = tempNode.next
	}
}

func main() {
	queue := Queue{&Node{data: 20}}
	enqueue(queue, 50)
	enqueue(queue, 51)
	enqueue(queue, 52)
	printQueue(queue)
	fmt.Println()
	top, _ := top(queue)
	fmt.Println(top)
	printQueue(queue)
	fmt.Println()
	dequeue(&queue)
	dequeue(&queue)
	printQueue(queue)
	fmt.Println()
}
