package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func top(head *Node) int {
	if head == nil {
		fmt.Println("Stack is empty")
		return 0
	}
	return head.Data
}

func pop(head **Node) {
	if *head == nil {
		fmt.Println("Stack is empty")
		return
	}
	tempNode := (*head).Next
	*head = nil
	*head = tempNode
}

func traverse(head *Node) {
	tempNode := head
	for tempNode != nil {
		fmt.Print(tempNode.Data)
		if tempNode.Next != nil {
			fmt.Print(" -> ")
		}
		tempNode = tempNode.Next
	}
	fmt.Println()
}

func push(head **Node, data int) {
	newNode := &Node{data, nil}
	if *head == nil {
		*head = newNode
		return
	}
	newNode.Next = *head
	*head = newNode
}

func main() {
	stockPrices := []int{6, 3, 4, 5, 2, 7}
	priceSpans := findPriceSpans(stockPrices)
	fmt.Println(priceSpans)

	priceSpans1 := findSpanStack(stockPrices)
	fmt.Println(priceSpans1)

}

func findPriceSpans(a []int) []int {
	s := make([]int, len(a), len(a))
	for j, _ := range a {
		i := 1
		for i <= j && a[j] > a[j-i] {
			i += 1
		}
		s[j] = i
	}
	return s
}

func findSpanStack(a []int) []int {
	s := make([]int, len(a), len(a))
	i := -1
	var stack *Node
	for j, _ := range a {
		for stack != nil && a[j] > a[top(stack)] {
			pop(&stack)
		}
		if stack == nil {
			i = -1
		} else {
			i = top(stack)
		}
		s[j] = j - i
		push(&stack, j)
	}
	return s
}
