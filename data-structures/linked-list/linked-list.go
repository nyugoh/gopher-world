package main

import "fmt"

type Node struct {
	next *Node
	data int
}

func insertEnd(head *Node, data int) *Node {
	if head == nil {
		return &Node{
			next: nil,
			data: data,
		}
	}
	tempNode := head
	for tempNode.next != nil {
		tempNode = tempNode.next
	}
	tempNode.next = &Node{
		next: nil,
		data: data,
	}
	return head
}

func traverse(head *Node) {
	if head != nil {
		fmt.Print(head.data, " -->")
		traverse(head.next)
	}
}

func main() {
	head := &Node{
		data: 20,
		next: nil,
	}
	insertEnd(head, 21)
	insertEnd(head, 22)
	insertEnd(head, 23)
	insertEnd(head, 24)
	insertEnd(head, 25)
	traverse(head)
	fmt.Println()

}
