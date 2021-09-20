package main

import (
	"fmt"
	"github.com/nyugoh/gopher-world/data-structures/queue/queue"
	binaryTree "github.com/nyugoh/gopher-world/data-structures/trees/binary-tree"
	"math"
)

/*
1. Give an algorithm for finding maximum element in binary tree.
2. Give an algorithm for finding the maximum element in binary tree without recursion.
*/

func findMax(root *binaryTree.BinaryTree) int {
	if root == nil {
		return 0
	}
	if root.LeftNode == nil && root.RightNode == nil {
		return root.Data.(int)
	}

	maxLeft := findMax(root.LeftNode)
	maxRight := findMax(root.RightNode)
	max := math.MinInt64

	if maxLeft > maxRight {
		max = maxLeft
	} else {
		max = maxRight
	}
	if root.Data.(int) > max {
		max = root.Data.(int)
	}
	return max
}

func findMaxNR(root *binaryTree.BinaryTree) (max int) {
	q := queue.CreateQueue()
	max = math.MinInt64
	if root == nil {
		return
	}
	q.Enqueue(root)
	for !q.IsEmpty() {
		temp := q.Dequeue().(*binaryTree.BinaryTree)
		if temp.Data.(int) > max {
			max = temp.Data.(int)
		}
		if temp.LeftNode != nil {
			q.Enqueue(temp.LeftNode)
		}
		if temp.RightNode != nil {
			q.Enqueue(temp.RightNode)
		}
	}
	return
}

func main() {
	t := &binaryTree.BinaryTree{10, nil, nil}
	t.Insert(t, 5)
	t.Insert(t, 15)
	t.Insert(t, 3)
	t.Insert(t, 7)
	t.Insert(t, 12)
	t.Insert(t, 17)

	fmt.Println("Max val::", findMax(t))
	fmt.Println("Max val::", findMaxNR(t))

}
