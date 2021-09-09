package main

import "fmt"

/*
Find nth node from the end of a Linked List.
*/

/* Solutions
1. Scan thru the list, take the first ele, and count the total elements after it say M , if M+1 == n then we found it,
   else if M+1 < n there are less elements than O(n2)
2. Create a map<position, Node> of list as you loop thru it, get element at M-n+1 Time complexity of creating the list O(M)
3. Find the length of list i.e M, then scan list upto M-n+1 O(n) - loops more than once i.e twice
4. Loop thru the list with two pointers, P & K, move K only until it makes n moves, from there increment both until K gets
   to the end. P will be at position n - Does this in one scan
*/

func solution1(head *Node, position int) (nthNode *Node) {
	currNode := head
	for currNode.Next != nil {
		tempNode, n := currNode, 0
		for tempNode.Next != nil {
			n += 1
			tempNode = tempNode.Next
		}
		if n+1 == position {
			nthNode = currNode
			break
		}
		currNode = currNode.Next
	}
	return
}

func solution2(head *Node, position int) (nthNode *Node) {
	items := map[int]*Node{}
	currNode, i := head, 1
	for currNode != nil {
		items[i] = currNode
		i += 1
		currNode = currNode.Next
	}
	if len(items) >= position {
		nthNode = items[(len(items) - position)+1]
	}
	return
}

func solution3(head *Node, position int) (nthNode *Node) {
	currNode, length := head, 0
	for currNode != nil {
		length += 1
		currNode = currNode.Next
	}
	k, i := (length-position)+1, 1
	currNode = head
	if k <= 0 || k > length {
		return
	}
	for currNode != nil && i < k {
		i += 1
		currNode = currNode.Next
	}
	fmt.Println(currNode.Data)
	nthNode = currNode
	return
}

func solution4(head *Node, position int) (nthNode *Node) {
	p, k, i := head, head, 0
	for k != nil {
		i += 1
		k = k.Next
		if i == position {
			break
		}
	}
	if position <= 0 || position > i {
		return
	}
	for k != nil && p != nil {
		k = k.Next
		p = p.Next
	}
	nthNode = p
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
	head := &Node{
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
	traverse(head)
	nthNode := solution4(head, 0)
	if nthNode != nil {
		fmt.Println(nthNode.Data)
	} else {
		fmt.Println("Node not found...")
	}
}