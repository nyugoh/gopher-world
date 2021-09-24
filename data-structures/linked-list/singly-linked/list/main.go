package linkedList

import "fmt"

type LinkedList struct {
	Data int
	Next *LinkedList
}

func CountLength(head *LinkedList) (length int) {
	current := head
	for current != nil {
		length += 1
		current = current.Next
	}
	return
}

func TraverseList(head *LinkedList) {
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

func InsertNode(head **LinkedList, data, position int) {
	fmt.Println("Inserting", data, "at", position)
	newNode := LinkedList{
		Data: data,
	}
	if position == 0 {
		newNode.Next = *head
		*head = &newNode
		return
	}
	prevNode, currNode, k := *head, *head, 0
	for currNode != nil && k < position {
		k++
		prevNode = currNode
		currNode = currNode.Next
	}
	prevNode.Next = &newNode
	newNode.Next = currNode
}

func DeleteNode(head **LinkedList, position int) {
	fmt.Println("Deleting position", position)
	if *head == nil {
		fmt.Println("List is empty")
		return
	}
	if position == 0 || position == 1 {
		tempNode := *head
		*head = tempNode.Next
		tempNode = nil
	}
	prevNode, currNode, k := *head, *head, 1
	for currNode != nil && k < position {
		k++
		prevNode = currNode
		currNode = currNode.Next
	}
	if currNode == nil {
		fmt.Println("Position doesn't exist")
		return
	}
	prevNode.Next = currNode.Next
	currNode = nil
	return
}

func DeleteList(head **LinkedList) {
	currNode := *head
	for currNode != nil {
		tempNode := currNode
		currNode = tempNode.Next
		tempNode = nil
	}
	*head = nil
	return
}
