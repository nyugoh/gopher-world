package main

import (
	"fmt"
	list "github.com/nyugoh/gopher-world/data-structures/linked-list/singly-linked/list"
)

func main() {
	head := &list.LinkedList{Data: 1}
	list.InsertNode(&head, 2, 2)
	list.InsertNode(&head, 3, 3)
	list.InsertNode(&head, 4, 4)
	list.InsertNode(&head, 5, list.CountLength(head))
	list.InsertNode(&head, 6, list.CountLength(head)+1)
	//list.DeleteListRecursive(&head)
	list.TraverseList(head)
	fmt.Println(head)
	/*
		list.DeleteNode(&head, 1)
		list.TraverseList(head)
		list.InsertNode(&head, 0, 1)
		list.TraverseList(head)
		list.DeleteNode(&head, 2)
		list.TraverseList(head)
		list.InsertNode(&head, 8, list.CountLength(head)-1)
		list.TraverseList(head)
		list.DeleteNode(&head, list.CountLength(head))
		list.DeleteNode(&head, list.CountLength(head)+1)
		list.TraverseList(head)
		list.InsertNode(&head, 12, list.CountLength(head))
		list.TraverseList(head)
		list.InsertNode(&head, 14, list.CountLength(head)+1)
		list.TraverseList(head)
		list.InsertNode(&head, 15, list.CountLength(head)+1)
	*/
	fmt.Println("Length::", list.CountLength(head))
	list.TraverseList(head)
}
