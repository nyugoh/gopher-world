package main

import (
	"fmt"
	"github.com/nyugoh/gopher-world/data-structures/stack/stack"
	binaryTree "github.com/nyugoh/gopher-world/data-structures/trees/binary-tree"
	"math"
)

func find(root *binaryTree.BinaryTree, data int) *binaryTree.BinaryTree {
	if root == nil {
		return nil
	}
	if data < root.Data.(int) {
		return find(root.LeftNode, data)
	} else if data > root.Data.(int) {
		return find(root.RightNode, data)
	}
	return root
}

func findNR(root *binaryTree.BinaryTree, data int)  *binaryTree.BinaryTree {
	if root == nil {
		return nil
	}
	for root != nil {
		if root.Data.(int) == data {
			return root
		}
		if data < root.Data.(int) {
			root = root.LeftNode
		} else if data > root.Data.(int) {
			root = root.RightNode
		}
	}
	return root
}

func findMin(root *binaryTree.BinaryTree) *binaryTree.BinaryTree {
	if root == nil {
		return nil
	}
	if root.LeftNode == nil {
		return root
	}
	return findMin(root.LeftNode)
}

func findMinNR(root *binaryTree.BinaryTree) *binaryTree.BinaryTree {
	if root == nil {
		return nil
	}
	for root.LeftNode != nil {
		root = root.LeftNode
	}
	return root
}

func findMax(root *binaryTree.BinaryTree) *binaryTree.BinaryTree {
	if root == nil {
		return nil
	}
	for root.RightNode != nil {
		root = root.RightNode
	}
	return root
}

func findMaxRecursive(root *binaryTree.BinaryTree) *binaryTree.BinaryTree {
	if root == nil {
		return nil
	}
	if root.RightNode == nil {
		return root
	}
	return findMaxRecursive(root.RightNode)
}

/*
1. Given pointers to two nodes in a binary search tree, find the lowest common ancestor (LCA).
Assume that both values already exist in the tree.
2. Give an algorithm for finding the shortest path between two nodes in a BST. - same as finding LCA
*/

// LCA If a < root.Data > b is the LCA
func LCA(root, a, b *binaryTree.BinaryTree) *binaryTree.BinaryTree {
	if root  == nil {
		return root
	}
	for {
		if (a.Data.(int) < root.Data.(int) && root.Data.(int) < b.Data.(int)) || (a.Data.(int) > root.Data.(int) && root.Data.(int) > b.Data.(int)) {
			return root
		}
		if a.Data.(int) < root.Data.(int) {
			root = root.LeftNode
		} else {
			root = root.RightNode
		}
	}
}

/*
1. Give an algorithm to check whether the given binary tree is a BST or not.
*/

func isBST(root *binaryTree.BinaryTree) bool {

/*
	// Below implementation is wrong since it doesn't check deeper than the current node
	if root == nil {
		return true
	}
	if root.LeftNode != nil && root.LeftNode.Data.(int) > root.Data.(int) {
		return false
	}
	if root.RightNode != nil && root.RightNode.Data.(int) < root.Data.(int) {
		return false
	}
	return isBST(root.LeftNode) && isBST(root.RightNode)

*/

/*
	// Correct but O(n2)
	if root == nil {
			return true
	}
	if root.LeftNode != nil && FindMax(root.LeftNode) > root.Data.(int) {
		return false
	}
	if root.RightNode != nil && FindMin(root.RigtNode) < root.Data.(int) {
		return false
	}
	return isBST(root.LeftNode) && isBST(root.RightNode)
*/
	// Efficient O(n)
	return isBSTUtil(root, math.MinInt64, math.MaxInt64)
}

func isBSTUtil(root *binaryTree.BinaryTree, min, max int) bool {
	// empty tree is a BST
	if root == nil {
		return true
	}

	if root.Data.(int) < min || root.Data.(int) > max {
		return false
	}

	return isBSTUtil(root.LeftNode, min, root.Data.(int)-1) && isBSTUtil(root.RightNode, root.Data.(int)+1, max)
}

func isBSTInorder(root *binaryTree.BinaryTree) bool {
	s := stack.Stack{}
	temp := root
	isBst := true
	prev := math.MinInt64
	for {
		for temp != nil {
			s.Push(temp)
			temp = temp.LeftNode
		}
		if s.IsEmpty() {
			break
		}
		temp, _ = s.Top().(*binaryTree.BinaryTree)
		s.Pop()
		if temp.Data.(int) < prev {
			isBst = false
			break
		}
		prev = temp.Data.(int)
		temp = temp.RightNode
	}
	return isBst
}

/*
Given a sorted array, give an algorithm for converting the array to BST.
*/
func buildBST(a []int, j, k int) *binaryTree.BinaryTree {
	if j > k {
		return nil
	}
	newNode := &binaryTree.BinaryTree{}
	if j == k {
		newNode.Data = a[j]
	} else {
		mid := j + (k-j)/2
		newNode.Data = a[mid]
		newNode.LeftNode = buildBST(a, j, mid-1)
		newNode.RightNode = buildBST(a, mid+1, k)
	}
	return newNode
}

// The insert function used below creates a binary search tree
func main() {
	t := &binaryTree.BinaryTree{10, nil, nil}
	t.Insert(t, 5)
	t.Insert(t, 15)
	t.Insert(t, 3)
	t.Insert(t, 7)
	t.Insert(t, 6)
	t.Insert(t, 12)
	t.Insert(t, 17)

	node:= find(t, 15)
	fmt.Println(node)

	fmt.Println(findNR(t, 17))

	fmt.Println("Min::",  findMin(t))
	fmt.Println("Min::",  findMinNR(t))

	fmt.Println("Max::",  findMax(t))
	fmt.Println("Max::",  findMaxRecursive(t))


	fmt.Println("Is BST::",  isBSTInorder(t))

	a := []int{1, 2, 3, 4, 5, 6, 7}
	t1 := buildBST(a, 0, len(a)-1)
	fmt.Println("Is BST::", isBST(t1))
}

