package main

import (
	"fmt"
	adjlist "github.com/nyugoh/gopher-world/data-structures/graphs/adj-list/graph"
)

func main() {
	list := adjlist.Graph{}
	list.CreateAdjList()

	list.AddEdgeM(0, 1)
	list.AddEdgeM(0, 2)
	list.AddEdgeM(1, 2)
	list.AddEdgeM(2, 0)
	list.AddEdgeM(2, 3)
	list.AddEdgeM(3, 3)

	list.Print()

	list.BFS(2)
	fmt.Println()
	list.DFS(2)
}