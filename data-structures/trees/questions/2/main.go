package main

import (
	"fmt"
	binaryTree "github.com/nyugoh/gopher-world/data-structures/trees/binary-tree"
)

/*
1. Give an algorithm for searching an element in binary tree.
2. Give an algorithm for finding the size of binary tree.
3. Give an algorithm for finding the height (or depth) of the binary tree.
4. Give an algorithm for finding the number of leaves in the binary tree without using recursion.
5. Given two binary trees, return true if they are structurally identical.
6. Given a binary tree, print out all its root-to-leaf paths.
*/

func findNode(root *binaryTree.BinaryTree, data int) bool {
	if root == nil {
		return false
	}

	if root.Data.(int) == data {
		return true
	}
	if findNode(root.LeftNode, data) {
		return true
	} else {
		return findNode(root.RightNode, data)
	}
}

// Non-recursive version can use Level order traversal
func findSize(root *binaryTree.BinaryTree) int {
	if root == nil {
		return 0
	} else {
		return findSize(root.LeftNode) + 1 + findSize(root.RightNode)
	}
}

func findHeight(root *binaryTree.BinaryTree) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := 0, 0
	leftHeight = findHeight(root.LeftNode)
	rightHeight = findHeight(root.RightNode)

	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}
	return maxHeight + 1
}

func findCountLeaves(root *binaryTree.BinaryTree) int {
	if root == nil {
		return 0
	}
	if root.LeftNode == nil && root.RightNode == nil {
		return 1
	}
	return findCountLeaves(root.LeftNode) + findCountLeaves(root.RightNode)
}

func areIdentical(r1, r2 *binaryTree.BinaryTree) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1.Data == r2.Data {
		return areIdentical(r1.LeftNode, r2.LeftNode) && areIdentical(r1.RightNode, r2.RightNode)
	} else {
		return false
	}
}

func printPaths(root *binaryTree.BinaryTree, path []int, length int) {
	if root == nil {
		return
	}
	path[length] = root.Data.(int)
	length += 1

	if root.LeftNode == nil && root.RightNode == nil {
		fmt.Println(path)
	}
	printPaths(root.LeftNode, path, length)
	printPaths(root.RightNode, path, length)
}

func main() {
	t := &binaryTree.BinaryTree{10, nil, nil}
	t.Insert(t, 5)
	t.Insert(t, 15)
	t.Insert(t, 3)
	t.Insert(t, 7)
	t.Insert(t, 12)
	t.Insert(t, 17)

	if findNode(t, 17) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}

	fmt.Println("Size of tree::", findSize(t))
	fmt.Println("Height of tree::", findHeight(t))
	fmt.Println("No. of leaf nodes::", findCountLeaves(t))
	fmt.Println("Are identical::", areIdentical(t, t))
	printPaths(t, make([]int, findHeight(t), findHeight(t)), 0)
}
