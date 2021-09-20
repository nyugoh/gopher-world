package binaryTree

import (
	"fmt"
	"github.com/nyugoh/gopher-world/data-structures/queue/queue"
	"github.com/nyugoh/gopher-world/data-structures/stack/stack"
)

/*
Properties :-
 - Has 0, 1 or 2 children per node

Operations :-
 - Insert new node
 - Delete node
 - Traversal, pre-order, in-order, post-order, level, breadth
 - Search for element
 - Find Size,
 - Find height

Applications :-
 - Expression tree, used in compilers
 - Huffman coding trees are used in data compression
 - BST, used in search, insertion and deletion in O(log(n)) in average case
 - Priority queue, which support search and delete in O(log(n)) in worst case
*/

type BinaryTree struct {
	Data interface{}
	LeftNode *BinaryTree
	RightNode *BinaryTree
}

// PreorderTraversal D->L->R
func (tree *BinaryTree) PreorderTraversal(root *BinaryTree) {
	if root != nil {
		fmt.Print(root.Data, " ")
		tree.PreorderTraversal(root.LeftNode)
		tree.PreorderTraversal(root.RightNode)
	}
}

func (tree *BinaryTree) PreorderTraversalNonRecursive() {
	s := stack.Stack{}
	root := tree
	for {
		for root != nil {
			fmt.Print(root.Data, " ")
			s.Push(root)
			root = root.LeftNode
		}
		if s.IsEmpty() {
			break
		}
		root, _ = s.Top().(*BinaryTree)
		s.Pop()
		root = root.RightNode
	}
}

// InorderTraversal L->D->R
func (tree *BinaryTree) InorderTraversal(root *BinaryTree) {
	if root != nil {
		tree.InorderTraversal(root.LeftNode)
		fmt.Print(root.Data, " ")
		tree.InorderTraversal(root.RightNode)
	}
}

func (tree *BinaryTree) InorderTraversalNonRecursive() {
	s := stack.Stack{}
	root := tree
	for {
		for root != nil {
			s.Push(root)
			root = root.LeftNode
		}
		if s.IsEmpty() {
			break
		}
		root, _ = s.Top().(*BinaryTree)
		s.Pop()
		fmt.Print(root.Data, " ")
		root = root.RightNode
	}
}

// PostorderTraversal L->R->D
func (tree *BinaryTree) PostorderTraversal(root *BinaryTree) {
	if root != nil {
		tree.PostorderTraversal(root.LeftNode)
		tree.PostorderTraversal(root.RightNode)
		fmt.Print(root.Data, " ")
	}
}

func (tree *BinaryTree) PostorderTraversalNonRecursive() {
	s := stack.Stack{}
	root := tree
	var prev *BinaryTree
	for k := true; k ; k = !s.IsEmpty() {
		for root != nil {
			s.Push(root)
			root = root.LeftNode
		}
		for root == nil && !s.IsEmpty() {
			root, _ = s.Top().(*BinaryTree)
			if root.RightNode == nil || root.RightNode == prev {
				fmt.Print(root.Data, " ")
				s.Pop()
				prev = root
				root = nil
			} else {
				root = root.RightNode
			}
		}
	}
}

func (tree *BinaryTree) LevelOrderTraversal() {
	q := queue.CreateQueue()
	root := tree
	if root == nil {
		return
	}
	q.Enqueue(root)
	for !q.IsEmpty() {
		root = q.Dequeue().(*BinaryTree)
		fmt.Print(root.Data, " ")

		if root.LeftNode != nil {
			q.Enqueue(root.LeftNode)
		}
		if root.RightNode != nil {
			q.Enqueue(root.RightNode)
		}
	}
}

func (tree *BinaryTree) BreadthFirstTraversal(root *BinaryTree) {

}

func (tree *BinaryTree) Insert(root *BinaryTree, data interface{}) *BinaryTree {
	// Check if root is null, make this node new root and return it
	if root == nil {
		return &BinaryTree{data, nil,nil}

	}

	// Root already exists
	if root.Data == data {
		return root
	}

	// Check which side it belongs
	if data.(int) < root.Data.(int) {
		root.LeftNode = tree.Insert(root.LeftNode, data)
	}

	if data.(int) > root.Data.(int) {
		root.RightNode = tree.Insert(root.RightNode, data)
	}
	return root
}

func (tree *BinaryTree) InsertInLevelOrder(data int) {
	q := queue.CreateQueue()
	temp := tree
	newNode := &BinaryTree{data, nil, nil}
	if tree == nil {
		tree = newNode
		return
	}
	q.Enqueue(temp)
	for !q.IsEmpty() {
		temp = q.Dequeue().(*BinaryTree)
		if temp.LeftNode != nil {
			q.Enqueue(temp.LeftNode)
		} else {
			temp.LeftNode = newNode
			return
		}

		if temp.RightNode != nil {
			q.Enqueue(temp.RightNode)
		} else {
			temp.RightNode = newNode
			return
		}
	}
}