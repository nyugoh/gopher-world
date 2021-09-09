package main

import "fmt"

/* Question
 Check whether the given linked list is either NULL-terminated or ends in a cycle (cyclic).
*/

/*
- If there's a loop, then there are more than one node pointing to the same element
1. Pick each node and check the remaining nodes if they are pointing to the same value
2. Use a hash map to track the occurence of next pointer, if a node appeares more than once, there's a loop
3. Floyd's method -  loop using two pointers,
*/

func solution3(head *Node) (looped bool) {
	slowPtr, fastPtr := head, head
	for slowPtr != nil && fastPtr != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if fastPtr == slowPtr {
			looped = true
			break
		}
	}
	return
}

type Node struct {
	Data int
	Next *Node
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
	n5 := &Node{
		Data: 5,
		Next: &Node{
			Data: 6,
			Next: nil,
		},
	}
	n4 := &Node{
		Data: 4,
		Next: n5,
	}
	n3 := &Node{
		Data: 4,
		Next: n5,
	}
	n5.Next.Next = n3

	head := &Node{
		Data: 1,
		Next: &Node{
			Data: 2,
			Next: &Node{
				Data: 3,
				Next: n4,
			},
		},
	}
	//traverse(head)

	if solution3(head) {
		fmt.Println("Looped")
	} else {
		fmt.Println("Not looped")
	}
}
