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
		return
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

// InsertNode
/*
	If positon is 1, i.e. start of the list, create new node and make it head
	If position p is in the middle:
		- Move to p-1, i.e. position before the position of the new node
		- Do this by maintaining two nodes, current and previous node
	If the inserting at the end, i.e. last position will be lenght of list + 1 then follow as in step above
	Update prevNode's next to be the new node and newNode's next to be current node
*/
func InsertNode(head **LinkedList, data, position int) {
	fmt.Println(fmt.Sprintf("Inserting %d at %d", data, position))
	newNode := LinkedList{
		Data: data,
	}
	if position < 1 {
		return
	}
	if position == 1 {
		newNode.Next = *head
		*head = &newNode
		return
	}
	prevNode, currNode, k := *head, *head, 1
	for currNode != nil && k < position {
		k++
		prevNode = currNode
		currNode = currNode.Next
	}
	prevNode.Next = &newNode
	newNode.Next = currNode
}

// DeleteNode deletes node at position P.
// Deleting at position 1:
//		- store current head in a temp var, update the head pointer to current head's next
//		- delete the temp var
// Deleting at the middle:
//		- traverse with prev node and current node till position P-1,
// Deleting at the end
//		- traverse to position l-1, where l is length of list
// once at the position, update prevNode's next to currentNode's next and delete current node
func DeleteNode(head **LinkedList, position int) {
	fmt.Println(fmt.Sprintf("Deleting position %d", position))
	if *head == nil {
		fmt.Println("List is empty")
		return
	}
	if position < 1 {
		return
	}
	if position == 1 {
		tempNode := *head
		*head = tempNode.Next
		tempNode = nil
		return
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

func DeleteListRecursive(head **LinkedList) {
	// base
	if *head == nil {
		return
	}

	tempNode := *head
	*head = tempNode.Next
	tempNode = nil
	DeleteListRecursive(head)
}
