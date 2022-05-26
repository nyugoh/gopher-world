package main

import (
	"fmt"
	"math/rand"
	"time"
)

const SIZE = 20

type Node struct {
	Data int
	Next *Node
}

type HashTable struct {
	Table map[int]*Node
	Size int
}

func hashFunction(value int) int {
	return value % SIZE
}

func insert(hashTable *HashTable, value int) {
	index := hashFunction(value)
	t := hashTable.Table[index]
	node := &Node{value, nil}
	if t == nil {
		hashTable.Table[index] = node
		return
	}
	for t.Next != nil {
		if t.Data == node.Data {
			return
		}
		t = t.Next
	}
	// Also check the last node
	if t.Data == node.Data {
		return
	}
	t.Next = node
}

func populate(table *HashTable, n int) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		data := rand.Intn(n)
		insert(table, data)
	}
}

func traverse(table *HashTable) {
	for k := range table.Table {
		t := table.Table[k]
		fmt.Println("Index ::", k)
		fmt.Print("\t")
		for t != nil {
			fmt.Printf("%d", t.Data)
			if t.Next != nil {
				fmt.Print(" -> ")
			}
			t = t.Next
		}
		fmt.Println()
	}
}

func lookUp(m *HashTable, value int) bool {
	index := hashFunction(value)
	t := m.Table[index]
	if t != nil {
		for t != nil {
			if t.Data == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	hashMap := &HashTable{Size: SIZE, Table: map[int]*Node{}}
	populate(hashMap, 500)
	traverse(hashMap)

	tt := []int{300, 0, 500, 115, 164}
	found := 0
	for _, i := range tt {
		if lookUp(hashMap, i) {
			fmt.Println(i, "not in hash map")
		} else {
			found += 1
			fmt.Println(i, "in hash map")
		}
	}
	fmt.Println(found, "/", len(tt))
}
