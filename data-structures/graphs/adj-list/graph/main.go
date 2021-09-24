package adjlist

import (
	"fmt"
	list "github.com/nyugoh/gopher-world/data-structures/linked-list/singly-linked/list"
	"github.com/nyugoh/gopher-world/data-structures/queue/queue"
)

type Graph struct {
	V int
	E int
	Adj []*list.LinkedList
}

func (g *Graph) CreateAdjList() {
	// Read no. of vertices
	fmt.Print("Enter no. of vertices::")
	if _, err := fmt.Scanf("%d", &g.V); err != nil {
		fmt.Println("Error::", err.Error())
		return
	}
	// Read no. of edges
	fmt.Print("Enter no. of edges::")
	if _, err := fmt.Scanf("%d", &g.E); err != nil {
		fmt.Println("Error::", err.Error())
		return
	}

	g.Adj = make([]*list.LinkedList, g.V, g.V)
	for i := 0; i < g.V; i++ {
		// Cread a node and make it the head
		newNode := &list.LinkedList{Data: i}
		g.Adj[i] = newNode
	}

	DFSVisited = make([]bool, g.V, g.V)

	fmt.Println(g)
}

func (g *Graph) CreateAdjListAuto(vertices, edges int) {
	g.V = vertices
	g.E = edges
	g.Adj = make([]*list.LinkedList, g.V, g.V)
	for i := 0; i < g.V; i++ {
		// Cread a node and make it the head
		newNode := &list.LinkedList{Data: i}
		g.Adj[i] = newNode
	}

	DFSVisited = make([]bool, g.V, g.V)
}

func (g *Graph) AddEdge() {
	var src, dest int
	fmt.Println("Enter edges, format is 0,1")
	for i := 0; i < g.E; i++ {
		fmt.Print("Enter edge no.", i+1, "::")
		if _, err := fmt.Scanf("%d %d", &src, &dest); err != nil {
			fmt.Println("Error reading edge::", err.Error())
			return
		}
		// Create an edge from src -> dest
		//newNode := &list.LinkedList{Data: dest}
		//newNode.Next = g.Adj[src]
		//g.Adj[src] = newNode

		list.InsertNode(&g.Adj[src], dest, list.CountLength(g.Adj[src]))

		// Create an edge from dest -> src - for a non-directed graph
		//newNode = &list.LinkedList{Data: src}
		//newNode.Next = g.Adj[dest]
		//g.Adj[dest] = newNode
	}
	fmt.Println(g.Adj)
}

func (g *Graph) AddEdgeM(src, dest int)  {
	if src == dest {
		return
	}
	list.InsertNode(&g.Adj[src], dest, list.CountLength(g.Adj[src]))
}

func (g *Graph) Print() {
	for i, row := range g.Adj {
		fmt.Print("Adj list for vertex ", i, " :")
		for row != nil {
			fmt.Print(row.Data)
			if row.Next != nil {
				fmt.Print(" -> ")
			}
			if row != row.Next {
				row = row.Next
			}
		}
		fmt.Println()
	}
}

func (g *Graph) BFS(start int) {
	visited := make([]bool, g.V, g.V)
	q := queue.CreateQueue()
	q.Enqueue(g.Adj[start])

	for !q.IsEmpty() {
		top := q.Dequeue().(*list.LinkedList)
		fmt.Print(top.Data, " -> ")
		for top.Next!= nil {
			top = top.Next
			if !visited[top.Data] {
				visited[top.Data] = true
				q.Enqueue(g.Adj[top.Data])
			}
		}
	}
}

var DFSVisited []bool

func (g *Graph) DFS(start int) {
	DFSVisited[start] = true
	fmt.Print(start, " -> ")
	top := g.Adj[start]
	for top.Next != nil {
		top = top.Next
		if !DFSVisited[top.Data] {
			g.DFS(top.Data)
		}
	}
}
