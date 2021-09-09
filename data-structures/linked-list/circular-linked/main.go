package main

import "fmt"

type CllNode struct {
	Data int
	Next *CllNode
}

func Traverse(head *CllNode) {
	currNode := head
	if currNode == nil {
		fmt.Println("List is empty")
		return
	}
	for b := true; b ; b = currNode != head {
		fmt.Print(currNode.Data)
		if currNode.Next != nil {
			fmt.Print(" -> ")
		}
		currNode = currNode.Next
	}
	fmt.Println()
}

func CountLength(head *CllNode) (length int) {
	currNode := head
	if head == nil {
		return
	}
	if currNode.Next == currNode {
		return 1
	}
	for b := true; b ; b = currNode != head {
		length += 1
		currNode = currNode.Next
	}
	return
}

// AddNodeAtFront Add node at the front
func AddNodeAtFront(head **CllNode, data int) {
	fmt.Println("Adding", data, "at the front")
	newNode := &CllNode{
		Data: data,
	}
	newNode.Next = newNode
	if *head == nil {
		*head = newNode
		return
	}
	currNode := *head
	for b := true; b ; b = currNode.Next != *head {
		currNode = currNode.Next
	}
	newNode.Next = *head
	currNode.Next = newNode
	*head = newNode
}

func AddNodeAtEnd(head **CllNode, data int) {
	fmt.Println("Adding", data, "at the end")
	newNode := &CllNode{Data: data}
	newNode.Next = newNode
	if *head == nil {
		fmt.Println("List is empty")
		return
	}
	currNode := *head
	for b := true; b ; b = currNode.Next != *head {
		currNode = currNode.Next
	}
	newNode.Next = *head
	currNode.Next = newNode
}

func DeleteAtLastNode(head **CllNode) {
	fmt.Println("Deleting last node...")
	if *head == nil {
		fmt.Println("List is empty")
		return
	}
	if *head == (*head).Next {
		*head = nil
		return
	}
	prevNode, currNode := *head, *head
	for b:=true; b ; b = currNode.Next != *head {
		prevNode = currNode
		currNode = currNode.Next
	}
	//fmt.Println(prevNode.Data)
	//fmt.Println(currNode.Data)
	//currNode.Next = currNode
	prevNode.Next = *head

}

func DeleteFirstNode(head **CllNode) {
	fmt.Println("Deleting first node...")
	if *head == nil {
		fmt.Println("List is empty")
		return
	}
	if *head == (*head).Next {
		*head = nil
		return
	}
	currNode := *head
	for b := true; b ; b = currNode.Next != *head {
		currNode = currNode.Next
	}
	currNode.Next = (*head).Next
	*head = nil
	*head = currNode.Next
}

func main() {
	head := &CllNode{
		Data: 1,
	}
	head.Next = head
	n := &CllNode{
		Data: 2,
		Next: head,
	}
	n1 := &CllNode{Data: 3}
	n.Next = n1
	n1.Next = head
	head.Next = n

	AddNodeAtFront(&head, 0)
	Traverse(head)
	fmt.Println("Length::", CountLength(head))

	AddNodeAtEnd(&head, 4)
	Traverse(head)
	fmt.Println("Length::", CountLength(head))

	DeleteFirstNode(&head)
	Traverse(head)
	fmt.Println("Length::", CountLength(head))

	DeleteAtLastNode(&head)
	Traverse(head)
	fmt.Println("Length::", CountLength(head))
}
