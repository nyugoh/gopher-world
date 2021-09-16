package main

import "fmt"

/* Question
Check whether the given linked list is either NULL-terminated or ends in a cycle (cyclic).
*/

/*
- If there's a loop, then there are more than one node pointing to the same element
1. Pick each node and check the remaining nodes if they are pointing to the same value
2. Use a hash map to track the occurence of next pointer, if a node appeares more than once, there's a loop
3. Floyd's method -  loop using two pointers, slowPtr moving 1 step at a time, fastPtr moving 2 steps at a time,
   if there's a loop, the two pointeres will meet at some point.
   Question b. Find the start of the loop
   Solution b. Once the two ptrs meet, reset slowPtr to start, i.e head and move them one step at a time till they meet

   Question c. Find length of the loop
   Solution c. Once you find the meeting point, move the fastPtr only one step at a time until it is the same as slowPtr
   while keeping a counter, counter will be the length
*/

func solution3(head *Node) (looped bool) {
	slowPtr, fastPtr := head, head
	for slowPtr != nil && fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if fastPtr == slowPtr {
			looped = true
			break
		}
	}
	if looped {
		slowPtr = head
		for slowPtr != nil && fastPtr != nil {
			slowPtr = slowPtr.Next
			fastPtr = fastPtr.Next
			if slowPtr == fastPtr {
				fmt.Println("Loop starts on node::", slowPtr.Data)
				break
			}
		}
		length := 0
		fastPtr = fastPtr.Next
		for slowPtr != fastPtr && fastPtr != nil {
			fastPtr = fastPtr.Next
			length += 1
		}
		fmt.Println("Length of the loop is::", length)
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
		Data: 7,
		Next: n4, // makes the loop, connects to an element already on list
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

	head = &Node{
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
	if solution3(head) {
		fmt.Println("Looped")
	} else {
		fmt.Println("Not looped")
	}
}
