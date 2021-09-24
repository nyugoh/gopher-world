package main

import (
	"fmt"
	adjlist "github.com/nyugoh/gopher-world/data-structures/graphs/adj-list/graph"
	ll "github.com/nyugoh/gopher-world/data-structures/linked-list/singly-linked/list"
	"github.com/nyugoh/gopher-world/data-structures/queue/queue"
)

var (
	Distance []int
	Path [][]int
)

func unweightedShortestPath(start int, g *adjlist.Graph) {
	q := queue.CreateQueue()
	q.Enqueue(g.Adj[start])
	Distance = make([]int, g.V, g.V)
	Path = make([][]int, g.V, g.V)

	// Mark all distances as -1 except from the start vertex
	for i := 0; i < g.V; i++ {
		Distance[i] = -1
		Path[i] = make([]int, 0, g.V)
	}
	Distance[start] = 0
	Path[start] = []int{start}

	for !q.IsEmpty() {
		head := q.Dequeue().(*ll.LinkedList)
		currNode := head
		for currNode != nil {
			if Distance[currNode.Data] == -1 {
				Distance[currNode.Data] = Distance[head.Data] + 1
				Path[currNode.Data] = append(Path[currNode.Data], Path[head.Data]...)
				Path[currNode.Data] = append(Path[currNode.Data], currNode.Data)
				q.Enqueue(g.Adj[currNode.Data])
			}
			currNode = currNode.Next
		}
	}
}

func main() {
	list := adjlist.Graph{}
	list.CreateAdjListAuto(7, 10)

	list.AddEdgeM(0, 1)
	list.AddEdgeM(0, 3)
	list.AddEdgeM(1, 3)
	list.AddEdgeM(1, 4)
	list.AddEdgeM(2, 0)
	list.AddEdgeM(2, 5)
	list.AddEdgeM(3, 5)
	list.AddEdgeM(3, 6)
	list.AddEdgeM(4, 6)
	list.AddEdgeM(6, 5)

	list.Print()

	list.BFS(2)
	fmt.Println()

	unweightedShortestPath(2, &list)

	fmt.Println(Distance)
	fmt.Println(Path)

}
