package main

import "fmt"

/*
Question:: Reverse a singly linked list.
*/

type Node struct {
	Data int
	Next *Node
}

func reverse(head *Node) (rest *Node) {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	nextNode := head.Next
	head.Next = nil
	fmt.Println(head.Data)
	reverseRest := reverse(nextNode)
	nextNode.Next = reverseRest
	return reverseRest
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
						Data: 5,
						Next: &Node{
							Data: 6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	traverse(head)
	traverse(reverse(head))

}
