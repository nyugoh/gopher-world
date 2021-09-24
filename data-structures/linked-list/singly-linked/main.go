package main

import (
	"fmt"
	list "github.com/nyugoh/gopher-world/data-structures/linked-list/singly-linked/list"
)

func main() {
	head := &list.LinkedList{
		Data: 1,
		Next: &list.LinkedList{
			Data: 2,
			Next: &list.LinkedList{
				Data: 3,
				Next: &list.LinkedList{
					Data: 4,
					Next: nil,
				},
			},
		},
	}
	list.TraverseList(head)
	list.DeleteList(&head)
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

	fmt.Println("Length::", list.CountLength(head))
	list.TraverseList(head)
}
