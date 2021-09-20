package main

import (
	"fmt"
	binaryTree "github.com/nyugoh/gopher-world/data-structures/trees/binary-tree"
)

func main() {
	t := &binaryTree.BinaryTree{10, nil, nil}
	t.Insert(t, 5)
	t.Insert(t, 15)
	t.Insert(t, 3)
	t.Insert(t, 7)
	t.Insert(t, 12)
	t.Insert(t, 17)
	t.InsertInLevelOrder(34)
	t.InsertInLevelOrder(1)
	t.InsertInLevelOrder(2)

	t.PreorderTraversal(t)
	fmt.Println()

	t.PreorderTraversalNonRecursive()
	fmt.Println()

	t.InorderTraversal(t)
	fmt.Println()

	t.InorderTraversalNonRecursive()
	fmt.Println()

	t.PostorderTraversal(t)
	fmt.Println()

	t.PostorderTraversalNonRecursive()
	fmt.Println()

	t.LevelOrderTraversal()
	fmt.Println()
	//fmt.Println(t)
}
