package stack

import "fmt"

type ListNode struct {
	Data interface{}
	Next *ListNode
}

type Stack struct {
	Head *ListNode
}

func (s *Stack) IsEmpty() bool {
	return s.Head == nil
}

func (s *Stack) Push(data interface{}) {
	//fmt.Println("Pushing ", data, "into stack")
	newNode := &ListNode{data, nil}

	if s.IsEmpty() {
		s.Head = newNode
		return
	}
	// Make new node the head
	newNode.Next = s.Head
	s.Head = newNode
}

func (s *Stack) Pop() {
	//fmt.Println("Poping stack")
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	//tempNode := s.Head
	s.Head = s.Head.Next
	//tempNode.Data
	//tempNode = nil
}

func (s *Stack) Top() (data interface{}) {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return data
	}
	data = s.Head.Data
	return data
}

func (s *Stack) Print() {
	tempNode := s.Head
	for tempNode != nil {
		fmt.Print(tempNode.Data)
		if tempNode.Next != nil {
			fmt.Print(" -> ")
		}
		tempNode = tempNode.Next
	}
	fmt.Println()
}
