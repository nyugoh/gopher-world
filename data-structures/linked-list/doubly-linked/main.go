package main

import "fmt"

type DllNode struct {
	Data int
	Next *DllNode
	Prev *DllNode
}

func CountLength(head *DllNode) (length int) {
	currNode := head
	for currNode != nil {
		length += 1
		currNode = currNode.Next
	}
	return
}

func TraverseForward(head *DllNode) {
	if head == nil {
		fmt.Println("List is empty")
		return
	}
	currNode := head
	for currNode != nil {
		if currNode.Prev != nil {
			fmt.Print(" <- ")
		}
		fmt.Print(currNode.Data)
		if currNode.Next != nil {
			fmt.Print(" -> ")
		}
		currNode = currNode.Next
	}
	fmt.Println()
}

func InsertNode(head **DllNode, data, position int) {
	fmt.Println("Inserting", data, "at position", position)
	newNode := DllNode{
		Data: data,
	}
	if position == 1 || *head == nil {
		newNode.Next = *head
		newNode.Prev = nil
		if *head != nil {
			(*head).Prev = &newNode
		}
		*head = &newNode
		return
	}
	currNode, prevNode, k := *head, *head, 1
	for currNode != nil && k < position {
		k++
		prevNode = currNode
		currNode = currNode.Next
	}
	newNode.Next = prevNode.Next
	prevNode.Next = &newNode
	if currNode != nil {
		newNode.Prev = currNode.Prev
		currNode.Prev = &newNode
	} else {
		newNode.Prev = prevNode
	}
}

func DeleteNode(head **DllNode, position int) {
	if head == nil {
		fmt.Println("List is empty")
		return
	}
	if position == 0 || position == 1 {
		nextNode := (*head).Next
		nextNode.Prev = nil
		*head = nextNode
		return
	}
	currNode, k := *head, 1
	for currNode.Next != nil && k < position-1 {
		k++
		currNode = currNode.Next
	}
	currNode.Prev.Next = currNode.Next
	if currNode.Next != nil {
		currNode.Next.Prev = currNode.Prev
	}
	currNode = nil

}

func main() {
	head := &DllNode{
		Data: 1,
	}
	head.Next = &DllNode{
		Data: 2,
		Prev: head,
	}
	head.Next.Next = &DllNode{
		Data: 3,
		Prev: head.Next,
	}

	fmt.Println("Length::", CountLength(head))
	TraverseForward(head)
	DeleteNode(&head, 1)
	TraverseForward(head)

	InsertNode(&head, 0, 1)
	InsertNode(&head, 4, CountLength(head)+1)
	InsertNode(&head, 5, CountLength(head)+3)

	fmt.Println("Length::", CountLength(head))
	TraverseForward(head)
}
