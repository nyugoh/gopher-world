package main

import "fmt"

type Tree struct {
	left  *Tree
	data  int
	right *Tree
}

//Check if root is not null,
func insert(root *Tree, data int) *Tree {
	// Check if root is null, make this node new root and return it
	if root == nil {
		return &Tree{
			left:  nil,
			data:  data,
			right: nil,
		}
	}

	// Root already exists
	if root.data == data {
		return root
	}

	// Check which side it belongs
	if data < root.data {
		root.left = insert(root.left, data)
	}

	if root.data < data {
		root.right = insert(root.right, data)
	}
	return root
}

func inorder(root *Tree) {
	if root != nil {
		inorder(root.left)
		fmt.Print(root.data, " ")
		inorder(root.right)
	}
}

func preorder(root *Tree) {
	if root != nil {
		fmt.Print(root.data, " ")
		preorder(root.left)
		preorder(root.right)
	}
}

func postorder(root *Tree) {
	if root != nil {
		postorder(root.left)
		postorder(root.right)
		fmt.Print(root.data, " ")
	}
}

func main() {
	root := &Tree{
		left:  nil,
		data:  50,
		right: nil,
	}
	insert(root, 20)
	insert(root, 30)
	insert(root, 10)
	insert(root, 40)
	insert(root, 60)
	insert(root, 70)
	insert(root, 80)
	insert(root, 90)
	inorder(root)
	fmt.Println()
	preorder(root)
	fmt.Println()
	postorder(root)
	fmt.Println()

}
