package main

import "fmt"

const SIZE = 26
type TrieNode struct {
	Letter rune
	IsWord bool
	Children [SIZE]*TrieNode
}


func NewTrie() *TrieNode {
	root := &TrieNode{}
	for i := 0; i < SIZE; i++ {
		root.Children[i] = &TrieNode{}
	}
	return root
}

func (trie *TrieNode) InsertTrie(word string) {
	if len(word) < 1 {
		return
	}

	index := word[0] - 'a' // will give us where the new node will be
	root := trie.Children[index]
	i, wordLen := 0, len(word)
	for ;i < wordLen; i++ {
		index := word[i] - 'a'
		if root.Children[index] == nil {
			root.Children[index] = &TrieNode{}
		}
		if i == wordLen-1 {
			root.IsWord = true
		}
		root = root.Children[i]
	}
}

func main() {
	root := NewTrie()
	root.InsertTrie("alll")
	for i, _ := range root.Children {
		if root.Children[i] != nil {
			fmt.Printf("%+v\n", root.Children[i])
		}
		for j, _ := range root.Children[i].Children {
			if root.Children[i].Children[j] != nil {
				fmt.Printf("\t%+v\n", root.Children[i].Children[j])
			}
			for k, _ := range root.Children[i].Children[j].Children {
				if root.Children[i].Children[j].Children[k] != nil {
					fmt.Printf("\t%+v\n", root.Children[i].Children[j].Children[k])
				}
			}
		}
	}
}
