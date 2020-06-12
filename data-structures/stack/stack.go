package main

import "fmt"

type Node struct {
	next *Node
	data int
}

// Remove element at the top
func pop(stack **Node) {
	// remove element at the front of list
	if stack == nil {
		fmt.Println("Stack is empty")
		return
	}
	tempNode := *stack
	*stack = tempNode.next
}

// Returns the item at the top
func top(stack *Node) (int, bool)  {
	// return element at the head of the list
	if stack == nil {
		fmt.Println("Stack is empty")
		return 0, true
	}
	return stack.data, false
}

// Add an element at the top of the stack
func push(head **Node, data int)  {
	// add element at the front
	newNode := &Node{data: data, next: nil}

	// if head is nil
	if head == nil {
		*head = newNode
	}
	newNode.next = *head
	*head = newNode
}

// Print stack content
func traverse(stack *Node)  {
	tempNode := stack
	for tempNode != nil {
		fmt.Print(tempNode.data, " ")
		tempNode = tempNode.next
	}
	fmt.Println()
}

func main() {
	stack := &Node{}
	push(&stack, 10)
	push(&stack, 20)
	push(&stack, 30)
	push(&stack, 40)
	push(&stack, 50)

	traverse(stack)
	pop(&stack)
	pop(&stack)
	traverse(stack)

}
