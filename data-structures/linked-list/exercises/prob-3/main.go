package main

import "fmt"

/*
Question :: Insert a node in a sorted linked list.
*/

/*
Solution:: Traverse list until you find position and insert item
*/

type Node struct {
	Data int
	Next *Node
}

func insertNode(head **Node, data int) {
	newNode := &Node{Data: data}
	prevNode, currNode := *head, *head
	if currNode == nil {
		fmt.Println("List is empty")
		*head = newNode
		return
	}
	if newNode.Data <= (*head).Data {
		newNode.Next = *head
		*head = newNode
		return
	}
	for currNode != nil && newNode.Data > currNode.Data {
		prevNode = currNode
		currNode = currNode.Next
	}
	prevNode.Next = newNode
	newNode.Next = currNode
}

func traverse(head *Node) {
	current := head
	if current == nil {
		fmt.Println("List is empty")
	}
	for current != nil {
		fmt.Print(current.Data)
		current = current.Next
		if current != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println("")
}

func main() {
	head := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: &Node{
					Data: 4,
					Next: &Node{
						Data: 6,
						Next: &Node{
							Data: 7,
							Next: nil,
						},
					},
				},
			},
		},
	}
	traverse(head)
	insertNode(&head, 5)
	traverse(head)
	insertNode(&head, 8)
	insertNode(&head, 0)
	traverse(head)
	insertNode(&head, 9)
	traverse(head)

}
